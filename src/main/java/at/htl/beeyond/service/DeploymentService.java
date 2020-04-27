package at.htl.beeyond.service;

import io.vertx.core.json.JsonObject;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.List;

@ApplicationScoped
public class DeploymentService {

    @Inject
    DeploymentYamlService deploymentYamlService;

    public void deployNginx(int replicaCount) {
        List<JsonObject> jsonObjects = this.deploymentYamlService.readYaml("templates", "nginx-deployment.yml");
        jsonObjects.get(1).getJsonObject("spec").put("replicas", replicaCount);
        this.deploymentYamlService.executeYaml(jsonObjects);
    }
}
