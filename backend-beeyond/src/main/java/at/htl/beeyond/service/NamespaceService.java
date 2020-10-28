package at.htl.beeyond.service;

import io.fabric8.kubernetes.api.model.Namespace;
import io.fabric8.kubernetes.api.model.NamespaceList;

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

    public void createNamespace(String namespace) {
        this.deploymentYamlService.getClient().namespaces().createNew()
                .withNewMetadata()
                .withName(namespace)
                .withLabels(Collections.singletonMap("managment", "beeyond"))
                .endMetadata()
                .done();
    }

    public void deleteNamespace(String namespace) {
        this.deploymentYamlService.getClient().namespaces().withName(namespace).delete();
    }
}
