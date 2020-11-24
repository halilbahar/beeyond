package at.htl.beeyond.intergration.application.template;

import at.htl.beeyond.intergration.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class TemplateApplicationTest {

    @Karate.Test
    Karate testCreateTemplateApplicationAsStudent() {
        return Karate.run("template-application-creation").tags("student").relativeTo(getClass());
    }
}
