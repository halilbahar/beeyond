package at.htl.beeyond.service;

import io.fabric8.kubernetes.api.model.HasMetadata;
import io.fabric8.kubernetes.client.DefaultKubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.vertx.core.json.JsonObject;
import org.yaml.snakeyaml.DumperOptions;
import org.yaml.snakeyaml.Yaml;

import javax.enterprise.context.ApplicationScoped;
import javax.json.Json;
import javax.json.JsonObjectBuilder;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;
import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.nio.charset.StandardCharsets;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@ApplicationScoped
public class DeploymentService {

    private Yaml yaml;
    private Jsonb jsonb;
    private KubernetesClient client;

    public DeploymentService() {
        DumperOptions dumperOptions = new DumperOptions();
        dumperOptions.setDefaultFlowStyle(DumperOptions.FlowStyle.BLOCK);
        dumperOptions.setPrettyFlow(true);

        this.yaml = new Yaml(dumperOptions);
        this.jsonb = JsonbBuilder.create();
        this.client = new DefaultKubernetesClient();
    }

    public void deployNginx(int replicaCount) {
        List<JsonObject> jsonObjects = readYaml("nginx-deployment.yml");

        jsonObjects.get(1).getJsonObject("spec").put("replicas", replicaCount);

        String yaml = jsonObjects.stream()
                .map(JsonObject::getMap)
                .map(this.yaml::dump)
                .collect(Collectors.joining("---\n"));

        InputStream stream = new ByteArrayInputStream((yaml.getBytes()));
        List<HasMetadata> list = client.load(stream).inNamespace("default").createOrReplace();
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
