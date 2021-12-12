package at.htl.beeyond.integration.template;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Order;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@Order(50)
public class TemplateTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("template-get")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testCreate() {
        return Karate.run("template-creation")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testDeletion() {
        return Karate.run("template-deletion")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }
}
