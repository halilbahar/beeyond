package at.htl.beeyond.intergration.template;

import at.htl.beeyond.intergration.DatabaseResource;
import at.htl.beeyond.intergration.IdentityProviderResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import io.quarkus.test.security.TestSecurity;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@QuarkusTestResource(IdentityProviderResource.class)
public class TemplateTest {

    @Karate.Test
    Karate testGetAllAsStudent() {
        return Karate.run("template-get").tags("student").relativeTo(getClass());
    }

    @Karate.Test
    Karate testGetAllAsTeacher() {
        return Karate.run("template-get").tags("teacher").relativeTo(getClass());
    }

    @Karate.Test
    @TestSecurity(user = "sonja-teacher", roles = "teacher")
    Karate testCreateAsTeacher() {
        return Karate.run("template-creation").tags("teacher").relativeTo(getClass());
    }
}
