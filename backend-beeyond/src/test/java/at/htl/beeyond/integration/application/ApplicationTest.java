package at.htl.beeyond.integration.application;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class ApplicationTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("application-get").relativeTo(getClass());
    }

    @Karate.Test
    Karate testApprove() {
        return Karate.run("application-approve").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDeny() {
        return Karate.run("application-deny").relativeTo(getClass());
    }

    @Karate.Test
    Karate testFinish() {
        return Karate.run("application-finish").relativeTo(getClass());
    }
}
