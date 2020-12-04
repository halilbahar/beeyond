package at.htl.beeyond.intergration.application.custom;

import at.htl.beeyond.intergration.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import io.quarkus.test.security.TestSecurity;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class CustomApplicationTest {

    @Karate.Test
    @TestSecurity(user = "sonja-teacher", roles = "teacher")
    Karate testCreateCustomApplicationAsStudent() {
        return Karate.run("custom-application-creation").tags("student").relativeTo(getClass());
    }
}
