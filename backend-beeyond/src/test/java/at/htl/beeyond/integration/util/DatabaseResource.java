package at.htl.beeyond.integration.util;

import com.github.dockerjava.api.exception.InternalServerErrorException;
import io.quarkus.test.common.QuarkusTestResourceLifecycleManager;
import org.testcontainers.containers.ContainerLaunchException;
import org.testcontainers.containers.PostgreSQLContainer;

import java.util.Arrays;
import java.util.Collections;
import java.util.Map;

public class DatabaseResource implements QuarkusTestResourceLifecycleManager {

    private static final PostgreSQLContainer<?> DATABASE = new PostgreSQLContainer<>("postgres:10.5")
            .withDatabaseName("beeyond_db")
            .withUsername("beeyond")
            .withPassword("beeyond");

    @Override
    public Map<String, String> start() {
        DATABASE.setPortBindings(Collections.singletonList("5432:5432"));

        try {
            DATABASE.start();
        } catch (ContainerLaunchException ex) {
            if(!ex.getMessage().contains("Container startup failed")) {
                ex.printStackTrace();
            }
        }

        String url = DATABASE.getHost() + ":5432";
        return Collections.singletonMap("beeyond.postgres.host", url);
    }

    @Override
    public void stop() {
        DATABASE.stop();
    }
}
