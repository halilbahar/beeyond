package at.htl.beeyond.integration.application.custom;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class CustomApplicationTest {

    @Karate.Test
    Karate testCreateCustomApplication() {
        return Karate.run("custom-application-creation").relativeTo(getClass());
    }
}
