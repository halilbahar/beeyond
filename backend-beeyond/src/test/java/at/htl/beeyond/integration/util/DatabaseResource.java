package at.htl.beeyond.integration.util;

import io.quarkus.test.common.QuarkusTestResourceLifecycleManager;
import org.testcontainers.containers.PostgreSQLContainer;

import java.util.Collections;
import java.util.Map;

public class DatabaseResource implements QuarkusTestResourceLifecycleManager {

    private static final PostgreSQLContainer<?> DATABASE = new PostgreSQLContainer<>("postgres:13.0")
            .withDatabaseName("beeyond_db")
            .withUsername("beeyond")
            .withPassword("beeyond");

    @Override
    public Map<String, String> start() {
        try {
            DATABASE.setPortBindings(Collections.singletonList("5432:5432"));
            DATABASE.start();
        } catch (Exception ignored) {
        }

        return Collections.singletonMap("beeyond.database.jdbc", "jdbc:postgresql://localhost:5432");
    }

    @Override
    public void stop() {
        DATABASE.stop();
    }
}
