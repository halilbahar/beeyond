package at.htl.beeyond.integration.namespace;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import io.quarkus.test.kubernetes.client.KubernetesServerTestResource;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@QuarkusTestResource(KubernetesServerTestResource.class)
public class NamespaceTest {
    @Karate.Test
    Karate testAssign() {
        return Karate.run("namespace-assign").relativeTo(getClass());
    }
}
