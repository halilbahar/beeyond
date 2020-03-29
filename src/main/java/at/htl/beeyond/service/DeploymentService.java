package at.htl.beeyond.service;

import org.yaml.snakeyaml.DumperOptions;
import org.yaml.snakeyaml.Yaml;

import javax.enterprise.context.ApplicationScoped;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;

@ApplicationScoped
public class DeploymentService {

    private Yaml yaml;
    private Jsonb jsonb;

    public DeploymentService() {
        DumperOptions dumperOptions = new DumperOptions();
        dumperOptions.setDefaultFlowStyle(DumperOptions.FlowStyle.BLOCK);
        dumperOptions.setPrettyFlow(true);

        this.yaml = new Yaml(dumperOptions);
        this.jsonb = JsonbBuilder.create();
    }
}
