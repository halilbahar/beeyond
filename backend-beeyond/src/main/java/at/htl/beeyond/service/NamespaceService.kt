package at.htl.beeyond.service

import io.fabric8.kubernetes.api.model.Namespace
import io.fabric8.kubernetes.api.model.ObjectMeta
import io.fabric8.kubernetes.client.KubernetesClientException
import java.util.*
import javax.enterprise.context.ApplicationScoped
import javax.inject.Inject

@ApplicationScoped
class NamespaceService {

    @Inject
    lateinit var deploymentService: DeploymentService

    fun createNamespace(name: String?) {
        val namespace = Namespace()
        val metadata = ObjectMeta()
        metadata.name = name
        metadata.labels = Collections.singletonMap("managment", "beeyond")
        namespace.metadata = metadata

        this.deploymentService.client.namespaces().create(namespace)
    }

    fun deleteNamespace(namespace: String?) {
        try {
            this.deploymentService.client.namespaces().withName(namespace).delete()
        } catch (ignored: KubernetesClientException) {
        }
    }
}
