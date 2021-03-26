package at.htl.beeyond.service

import at.htl.beeyond.entity.Application
import javax.enterprise.context.ApplicationScoped
import javax.json.bind.Jsonb
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.TemplateApplication
import io.fabric8.kubernetes.client.DefaultKubernetesClient
import io.fabric8.kubernetes.client.KubernetesClient
import org.yaml.snakeyaml.DumperOptions
import org.yaml.snakeyaml.Yaml
import org.yaml.snakeyaml.representer.Representer
import java.io.ByteArrayInputStream
import javax.json.Json
import javax.json.JsonObject
import javax.json.bind.JsonbBuilder

@ApplicationScoped
class DeploymentService {
    var jsonb: Jsonb = JsonbBuilder.create()
    lateinit var yaml: Yaml
    lateinit var client: KubernetesClient

    init {
        val dumperOptions = DumperOptions()
        dumperOptions.defaultFlowStyle = DumperOptions.FlowStyle.BLOCK
        dumperOptions.isPrettyFlow = true
        dumperOptions.tags = null
        this.yaml = Yaml(dumperOptions)
        this.client = DefaultKubernetesClient()
    }

    fun deploy(application: Application?) {
        if (application is CustomApplication) {
            this.executeYaml(application.content, application.id)
        } else if (application is TemplateApplication) {
            this.executeYaml(application.content, application.id)
        }
    }

    fun stop(application: Application?) {
        if (application is CustomApplication) {
            this.executeYaml(application.content, application.id, delete = true)
        } else if (application is TemplateApplication) {
            this.executeYaml(application.content, application.id, delete = true)
        }
    }

    fun executeYaml(content: String, applicationId: Long, delete: Boolean = false) {
        val yamlArray: MutableList<JsonObject> = mutableListOf()
        val yamlIterator: MutableIterator<Any> = this.yaml.loadAll(content).iterator()

        while (yamlIterator.hasNext()) {
            val next = yamlIterator.next()
            if (next is Map<*, *>) {
                yamlArray.add(Json.createObjectBuilder(next as Map<String, Any>).build())
            }
        }

        val yamlString = yamlArray
                .map {
                    val kubernetesBuilder = Json.createObjectBuilder(it)
                    val metadataBuilder = Json.createObjectBuilder(it.getJsonObject("metadata"))
                    metadataBuilder.add("labels", Json.createObjectBuilder().add("beeyond-application-id", applicationId))

                    kubernetesBuilder.add("metadata", metadataBuilder)
                    kubernetesBuilder.build().toString()
                }
                .map { this.jsonb.fromJson(it, Object::class.java) }
                .map { this.yaml.dump(it) }
                .joinToString("---\n")

        // If problems occur: https://stackoverflow.com/a/25750748/11125147
        if(!delete) {
            client.load(ByteArrayInputStream(yamlString.toByteArray())).inNamespace("default").createOrReplace()
        } else {
            client.load(ByteArrayInputStream(yamlString.toByteArray())).inNamespace("default").delete()
        }
    }
}