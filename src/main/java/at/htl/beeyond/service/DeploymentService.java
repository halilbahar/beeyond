package at.htl.beeyond.service;

import io.vertx.core.json.JsonObject;
import org.yaml.snakeyaml.DumperOptions;
import org.yaml.snakeyaml.Yaml;

import javax.enterprise.context.ApplicationScoped;
import javax.json.Json;
import javax.json.JsonObjectBuilder;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;
import java.io.InputStream;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

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

    private List<JsonObject> readYaml(String file) {
        List<JsonObject> result = new LinkedList<>();

        InputStream inputStream = this.getClass()
                .getClassLoader()
                .getResourceAsStream("/templates/" + file);
        Iterable<Object> objects = yaml.loadAll(inputStream);

        for (Object object : objects) {
            result.add(JsonObject.mapFrom(object));
        }

        return result;
    }
}
