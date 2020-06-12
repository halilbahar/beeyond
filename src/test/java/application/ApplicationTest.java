package application;

import com.intuit.karate.junit5.Karate;

public class ApplicationTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("application-get").relativeTo(getClass());
    }

    @Karate.Test
    Karate testApprove() {
        return Karate.run("application-patch").tags("approve").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDeny() {
        return Karate.run("application-patch").tags("deny").relativeTo(getClass());
    }
}
