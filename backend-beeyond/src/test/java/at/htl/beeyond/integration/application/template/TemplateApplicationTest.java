package at.htl.beeyond.integration.application.template;

import at.htl.beeyond.integration.util.DatabaseResource;
import at.htl.beeyond.integration.util.WiremockValidation;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@QuarkusTestResource(WiremockValidation.class)
public class TemplateApplicationTest {

    @Karate.Test
    Karate testCreateTemplateApplicationAsStudent() {
        return Karate.run("template-application-creation").relativeTo(getClass());
    }
}
