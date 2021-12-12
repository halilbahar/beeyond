package at.htl.beeyond.integration.namespace;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import io.quarkus.test.kubernetes.client.KubernetesServerTestResource;
import org.junit.jupiter.api.Order;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@QuarkusTestResource(KubernetesServerTestResource.class)
@Order(40)
public class NamespaceTest {
    @Karate.Test
    Karate testAssign() {
        return Karate.run("namespace-assign")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }
}
