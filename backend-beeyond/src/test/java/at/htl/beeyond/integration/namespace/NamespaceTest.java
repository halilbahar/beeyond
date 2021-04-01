package at.htl.beeyond.integration.namespace;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class NamespaceTest {
    @Karate.Test
    Karate testAssign() {
        return Karate.run("namespace-assign").relativeTo(getClass());
    }

    @Karate.Test
    Karate testRevoke() {
        return Karate.run("namespace-revoke").relativeTo(getClass());
    }
}
