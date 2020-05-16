package at.htl.beeyond.service;

import at.htl.beeyond.model.CustomApplication;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;

@ApplicationScoped
public class DeploymentService {

    @Inject
    DeploymentYamlService deploymentYamlService;
    Jsonb jsonb;

    public DeploymentService() {
        this.jsonb = JsonbBuilder.create();
    }

    public void deploy(CustomApplication customApplication) {
        this.deploymentYamlService.executeYaml(customApplication.getContent());
    }
}
