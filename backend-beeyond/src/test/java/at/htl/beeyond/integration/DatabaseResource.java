package at.htl.beeyond.integration;

import io.quarkus.test.common.QuarkusTestResourceLifecycleManager;
import org.testcontainers.containers.PostgreSQLContainer;

import java.util.Collections;
import java.util.Map;

public class DatabaseResource implements QuarkusTestResourceLifecycleManager {

    private static final PostgreSQLContainer<?> DATABASE = new PostgreSQLContainer<>("postgres:10.5")
            .withDatabaseName("beeyond_db")
            .withUsername("beeyond")
            .withPassword("beeyond");

    @Override
    public Map<String, String> start() {
        DATABASE.start();

        String url = DATABASE.getHost() + ":" + DATABASE.getMappedPort(5432);
        return Collections.singletonMap("beeyond.postgres.host", url);
    }

    @Override
    public void stop() {
        DATABASE.stop();
    }
}
