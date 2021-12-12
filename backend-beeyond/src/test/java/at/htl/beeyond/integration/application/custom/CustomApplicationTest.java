package at.htl.beeyond.integration.application.custom;

import at.htl.beeyond.integration.util.DatabaseResource;
import at.htl.beeyond.integration.util.WiremockValidation;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Order;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@QuarkusTestResource(WiremockValidation.class)
@Order(10)
public class CustomApplicationTest {

    @Karate.Test
    Karate testCreateCustomApplication() {
        return Karate.run("custom-application-creation")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }
}
