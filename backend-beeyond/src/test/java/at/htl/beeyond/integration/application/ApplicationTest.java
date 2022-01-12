package at.htl.beeyond.integration.application;

import at.htl.beeyond.integration.util.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Order;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
@Order(30)
public class ApplicationTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("application-get")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testApprove() {
        return Karate.run("application-approve")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testDeny() {
        return Karate.run("application-deny")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testFinish() {
        return Karate.run("application-finish")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testStart() {
        return Karate.run("application-start")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testStop() {
        return Karate.run("application-stop")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }

    @Karate.Test
    Karate testRequest() {
        return Karate.run("application-request")
                .outputCucumberJson(true)
                .outputHtmlReport(false)
                .relativeTo(getClass());
    }
}
