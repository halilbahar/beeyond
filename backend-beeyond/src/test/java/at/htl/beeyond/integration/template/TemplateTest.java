package at.htl.beeyond.integration.template;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import io.quarkus.test.security.TestSecurity;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class TemplateTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("template-get").relativeTo(getClass());
    }

    @Karate.Test
    Karate testCreate() {
        return Karate.run("template-creation").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDeletion() {
        return Karate.run("template-deletion").relativeTo(getClass());
    }

    @Karate.Test
    Karate testUpdate() {
        return Karate.run("template-update").relativeTo(getClass());
    }
}
