package at.htl.beeyond.service;

import at.htl.beeyond.entity.*;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;
import java.util.List;

@ApplicationScoped
public class DeploymentService {

    @Inject
    DeploymentYamlService deploymentYamlService;
    Jsonb jsonb;

    public DeploymentService() {
        this.jsonb = JsonbBuilder.create();
    }

    public void deploy(Application application) {
        if (application instanceof CustomApplication) {
            CustomApplication customApplication = (CustomApplication) application;
            this.deploymentYamlService.executeYaml(customApplication.getContent());
        } else if (application instanceof TemplateApplication) {
            TemplateApplication templateApplication = (TemplateApplication) application;
            Template template = templateApplication.getTemplate();
            List<TemplateFieldValue> fieldValues = templateApplication.getFieldValues();

            String content = template.getContent();

            for (TemplateFieldValue fieldValue : fieldValues) {
                String wildcard = fieldValue.getField().getWildcard();
                content = content.replace("%" + wildcard + "%", fieldValue.getValue());
            }

            this.deploymentYamlService.executeYaml(content);
        }
    }
}
