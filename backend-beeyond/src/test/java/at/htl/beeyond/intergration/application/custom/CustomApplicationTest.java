package at.htl.beeyond.intergration.application.custom;

import at.htl.beeyond.intergration.DatabaseResource;
import at.htl.beeyond.intergration.IdentityProviderResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@QuarkusTestResource(IdentityProviderResource.class)
public class CustomApplicationTest {

    @Karate.Test
    Karate testCreateCustomApplicationAsStudent() {
        return Karate.run("custom-application-creation").tags("student").relativeTo(getClass());
    }
}
