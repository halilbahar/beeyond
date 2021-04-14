package at.htl.beeyond.service;

import io.fabric8.kubernetes.api.model.Namespace;
import io.fabric8.kubernetes.api.model.NamespaceList;
import io.fabric8.kubernetes.api.model.ObjectMeta;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.Collections;
import java.util.List;

@ApplicationScoped
public class NamespaceService {

    @Inject
    DeploymentYamlService deploymentYamlService;

    public List<Namespace> getAllNamespaces() {
        NamespaceList namespaces = this.deploymentYamlService.getClient().namespaces().list();
        return namespaces.getItems();
    }

    public void createNamespace(String name) {
        Namespace namespace = new Namespace();
        ObjectMeta metadata = new ObjectMeta();

        metadata.setName(name);
        metadata.setLabels(Collections.singletonMap("managment", "beeyond"));
        namespace.setMetadata(metadata);

        this.deploymentYamlService.getClient().namespaces().create(namespace);
    }

    public void deleteNamespace(String namespace) {
        this.deploymentYamlService.getClient().namespaces().withName(namespace).delete();
    }
}
