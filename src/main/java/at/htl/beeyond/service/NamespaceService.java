package at.htl.beeyond.service;

import io.fabric8.kubernetes.api.model.Namespace;
import io.fabric8.kubernetes.api.model.NamespaceList;
import io.vertx.core.json.JsonObject;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.List;

@ApplicationScoped
public class NamespaceService {

    @Inject
    DeploymentYamlService deploymentYamlService;

    public void createNamespace(String namespace) {
        List<JsonObject> jsonObjects = this.deploymentYamlService.readYaml("config-templates", "namespace-template.yml");
        jsonObjects.get(0).getJsonObject("metadata").put("name", namespace);
        this.deploymentYamlService.executeYaml(jsonObjects);
    }

    public void deleteNamespace(String namespace) {
       this.deploymentYamlService.getClient().namespaces().withName(namespace).delete();
    }

    public List<Namespace> getAllNamespaces() {
        NamespaceList namespaces = this.deploymentYamlService.getClient().namespaces().list();
        return namespaces.getItems();
    }
}
