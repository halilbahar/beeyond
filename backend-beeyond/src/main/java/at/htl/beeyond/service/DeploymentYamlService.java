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

    public void executeYaml(String content) {
        this.client.load(new ByteArrayInputStream(content.getBytes())).inNamespace("default").createOrReplace();
    }

    public KubernetesClient getClient() {
        return client;
    }
}
