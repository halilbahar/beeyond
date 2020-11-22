package at.htl.beeyond.intergration;

import io.quarkus.test.common.QuarkusTestResourceLifecycleManager;
import org.testcontainers.containers.BindMode;
import org.testcontainers.containers.GenericContainer;
import org.testcontainers.containers.wait.strategy.Wait;

import java.util.Collections;
import java.util.Map;

public class IdentityProviderResource implements QuarkusTestResourceLifecycleManager {

    private final GenericContainer<?> IDENTITY_PROVIDER = new GenericContainer<>("jboss/keycloak:11.0.2")
            .waitingFor(Wait.forHttp("/auth").forPort(8080))
            .withEnv("KEYCLOAK_USER", "beeyond")
            .withEnv("KEYCLOAK_PASSWORD", "beeyond")
            .withClasspathResourceMapping("school-realm.json", "/tmp/school-realm.json", BindMode.READ_ONLY)
            .withCommand(
                    "-b", "0.0.0.0",
                    "-Dkeycloak.migration.action=import",
                    "-Dkeycloak.profile.feature.upload_scripts=enabled",
                    "-Dkeycloak.migration.provider=singleFile",
                    "-Dkeycloak.migration.file=/tmp/school-realm.json"
            );

    @Override
    public Map<String, String> start() {
        IDENTITY_PROVIDER.start();

        String host = IDENTITY_PROVIDER.getHost();
        Integer port = IDENTITY_PROVIDER.getMappedPort(8080);

        String url = "http://" + host + ":" + port;
        return Collections.singletonMap("beeyond.keycloak.host", url);
    }

    @Override
    public void stop() {
        IDENTITY_PROVIDER.stop();
    }
}
