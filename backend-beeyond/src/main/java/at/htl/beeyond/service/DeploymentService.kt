package at.htl.beeyond.service

import at.htl.beeyond.entity.Application
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.entity.TemplateApplication
import io.fabric8.kubernetes.api.model.IntOrString
import io.fabric8.kubernetes.api.model.ObjectMetaBuilder
import io.fabric8.kubernetes.api.model.extensions.*
import io.fabric8.kubernetes.client.DefaultKubernetesClient
import io.fabric8.kubernetes.client.KubernetesClient
import io.quarkus.security.UnauthorizedException
import org.eclipse.microprofile.config.inject.ConfigProperty
import org.yaml.snakeyaml.DumperOptions
import org.yaml.snakeyaml.Yaml
import java.io.ByteArrayInputStream
import javax.enterprise.context.ApplicationScoped
import javax.inject.Inject
import javax.json.Json
import javax.json.JsonObject
import javax.json.bind.Jsonb
import javax.json.bind.JsonbBuilder

@ApplicationScoped
class DeploymentService {
    var jsonb: Jsonb = JsonbBuilder.create()
    lateinit var yaml: Yaml
    lateinit var client: KubernetesClient

    @Inject
    lateinit var namespaceService: NamespaceService

    @ConfigProperty(name = "beeyond.kubernetes.host")
    lateinit var kubernetesHost: String

    init {
        val dumperOptions = DumperOptions()
        dumperOptions.defaultFlowStyle = DumperOptions.FlowStyle.BLOCK
        this.yaml = Yaml(dumperOptions)
        this.client = DefaultKubernetesClient()
    }

    fun deploy(application: Application?) {
        if (application is CustomApplication) {
            this.executeYaml(application.content, application.id, application.namespace)
        } else if (application is TemplateApplication) {
            this.executeYaml(application.content, application.id, application.namespace)
        }
    }

    fun stop(application: Application?) {
        val ingresses =
            this.client.extensions().ingresses().inNamespace(application!!.namespace.namespace)
                .withLabel("beeyond-application-id", application.id.toString())
                .list().items
        ingresses.forEach {
            this.client.extensions().ingresses().inNamespace(application!!.namespace.namespace).delete(it)
        }

        if (application is CustomApplication) {
            this.executeYaml(application.content, application.id, application.namespace, delete = true)
        } else if (application is TemplateApplication) {
            this.executeYaml(application.content, application.id, application.namespace, delete = true)
        }
    }

    fun executeYaml(content: String, applicationId: Long, namespace: Namespace, delete: Boolean = false) {
        val yamlArray: MutableList<JsonObject> = mutableListOf()
        val yamlIterator: MutableIterator<Any> = this.yaml.loadAll(content).iterator()

        while (yamlIterator.hasNext()) {
            val next = yamlIterator.next()
            if (next is Map<*, *>) {
                yamlArray.add(Json.createObjectBuilder(next as Map<String, Any>).build())
            }
        }

        val services = mutableMapOf<String, Int>()

        val yamlString = yamlArray
            .map {
                val kubernetesBuilder = Json.createObjectBuilder(it)
                val metadataBuilder = Json.createObjectBuilder(it.getJsonObject("metadata"))
                metadataBuilder.add("labels", Json.createObjectBuilder().add("beeyond-application-id", applicationId))
                kubernetesBuilder.add("metadata", metadataBuilder)

                if (it.getString("kind") == "Service" &&
                    it.getJsonObject("metadata").getJsonObject("labels")?.getString("beeyond-create-ingress") == "true"
                ) {
                    val serviceName = it.getJsonObject("metadata").getString("name");
                    it.getJsonObject("spec").getJsonArray("ports").map {
                        services.put(serviceName, it.asJsonObject().getInt("port"))
                    }
                }

                if (it.getString("kind") == "Ingress" &&
                    it.getJsonObject("spec").getJsonArray("rules").any {
                        it.asJsonObject().getJsonObject("http").getJsonArray("paths")
                            .any {
                                !it.asJsonObject().getString("path").startsWith("\"/" + namespace.namespace)
                            }
                    }
                ) {
                    throw UnauthorizedException("The ingress does not direct to a path with the namespace.")
                }

                kubernetesBuilder.build().toString()
            }
            .map { this.jsonb.fromJson(it, Object::class.java) }
            .map {
                this.yaml.dump(it)
            }
            .joinToString("---\n")
            .replace(Regex("!!float '([0-9]+)'"), "$1")

        // If problems occur: https://stackoverflow.com/a/25750748/11125147
        if (!delete) {
            if (!client.namespaces().list().items.map { it.metadata.name }.contains(namespace.namespace)) {
                namespaceService.createNamespace(namespace.namespace)
            }

            client.load(ByteArrayInputStream(yamlString.toByteArray())).inNamespace(namespace.namespace)
                .createOrReplace()

            if (!services.isEmpty()) {
                val rules = HTTPIngressRuleValueBuilder()
                services.forEach {
                    rules.addNewPath().withPath("/" + namespace.namespace + "(/|$)(.*)$").withBackend(
                        IngressBackendBuilder()
                            .withServicePort(IntOrString(it.value))
                            .withServiceName(it.key).build()
                    )
                        .withPathType("Prefix")
                        .endPath()
                }

                client.extensions().ingresses().inNamespace(namespace.namespace).create(
                    IngressBuilder()
                        .withMetadata(
                            ObjectMetaBuilder()
                                .withName(namespace.namespace + "-" + System.currentTimeMillis())
                                .addToLabels("beeyond-application-id", applicationId.toString())
                                .addToAnnotations("nginx.ingress.kubernetes.io/rewrite-target", "/$2")
                                .build()
                        )
                        .withSpec(
                            IngressSpecBuilder()
                                .addNewRule()
                                .withHost(kubernetesHost)
                                .withHttp(rules.build())
                                .endRule()
                                .build()
                        ).build()
                )
            }
        } else {
            client.load(ByteArrayInputStream(yamlString.toByteArray())).inNamespace(namespace.namespace).delete()
        }
    }
}
