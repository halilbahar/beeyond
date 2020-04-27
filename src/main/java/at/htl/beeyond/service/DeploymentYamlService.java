package at.htl.beeyond.service;

import io.fabric8.kubernetes.api.model.HasMetadata;
import io.fabric8.kubernetes.client.DefaultKubernetesClient;
import io.fabric8.kubernetes.client.KubernetesClient;
import io.vertx.core.json.JsonObject;
import org.yaml.snakeyaml.DumperOptions;
import org.yaml.snakeyaml.Yaml;

import javax.enterprise.context.ApplicationScoped;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;
import java.io.ByteArrayInputStream;
import java.io.InputStream;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

@ApplicationScoped
public class DeploymentYamlService {

    private Yaml yaml;
    private Jsonb jsonb;
    private KubernetesClient client;

    public DeploymentYamlService() {
        DumperOptions dumperOptions = new DumperOptions();
        dumperOptions.setDefaultFlowStyle(DumperOptions.FlowStyle.BLOCK);
        dumperOptions.setPrettyFlow(true);

        this.yaml = new Yaml(dumperOptions);
        this.jsonb = JsonbBuilder.create();
        this.client = new DefaultKubernetesClient();
    }

    public void executeYaml(List<JsonObject> jsonObjects) {
        executeYaml(jsonObjects, "default");
    }

    public void executeYaml(List<JsonObject> jsonObjects, String namespace) {
        String yaml = jsonObjects.stream()
                .map(JsonObject::getMap)
                .map(this.yaml::dump)
                .collect(Collectors.joining("---\n"));

        InputStream stream = new ByteArrayInputStream((yaml.getBytes()));
        List<HasMetadata> list = this.client.load(stream).inNamespace(namespace).createOrReplace();
    }

    public List<JsonObject> readYaml(String directory, String file) {
        List<JsonObject> result = new LinkedList<>();

        InputStream inputStream = this.getClass()
                .getClassLoader()
                .getResourceAsStream("/" + directory + "/" + file);
        Iterable<Object> objects = this.yaml.loadAll(inputStream);

        for (Object object : objects) {
            result.add(JsonObject.mapFrom(object));
        }
        return result;
    }

    public KubernetesClient getClient() {
        return client;
    }
}
