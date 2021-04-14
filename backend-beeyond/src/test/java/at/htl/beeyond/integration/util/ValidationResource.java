package at.htl.beeyond.integration.util;

import io.quarkus.test.common.QuarkusTestResourceLifecycleManager;
import org.testcontainers.containers.GenericContainer;
import org.testcontainers.containers.MongoDBContainer;
import org.testcontainers.images.builder.ImageFromDockerfile;

import java.nio.file.Path;
import java.util.Collections;
import java.util.Map;

public class ValidationResource implements QuarkusTestResourceLifecycleManager {

    private static final MongoDBContainer DATABASE = new MongoDBContainer("mongo:4.4.3")
            .withEnv("MONGO_INITDB_DATABASE", "beeyond_validation_db")
            .withEnv("MONGO_INITDB_ROOT_USERNAME", "beeyond")
            .withEnv("MONGO_INITDB_ROOT_PASSWORD", "beeyond");

    private static final GenericContainer<?> KUBERNETES_VALIDATION = new GenericContainer<>(
            new ImageFromDockerfile()
                    .withDockerfile(Path.of("..", "yaml-validation", "Dockerfile"))
    );

    @Override
    public Map<String, String> start() {
        try {
            KUBERNETES_VALIDATION.setPortBindings(Collections.singletonList("8180:8180"));
            KUBERNETES_VALIDATION.start();
        } catch (Exception ignored) {
        }

        try {
            DATABASE.setPortBindings(Collections.singletonList("27017:27017"));
            DATABASE.start();
        } catch (Exception ignored) {
        }

        return null;
    }

    @Override
    public void stop() {
        if (DATABASE.isRunning()) {
            DATABASE.stop();
        }
        if (KUBERNETES_VALIDATION.isRunning()) {
            KUBERNETES_VALIDATION.stop();
        }
    }
}
