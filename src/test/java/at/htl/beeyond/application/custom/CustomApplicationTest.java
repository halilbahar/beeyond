package at.htl.beeyond.application.custom;

import com.intuit.karate.junit5.Karate;

public class CustomApplicationTest {

    @Karate.Test
    Karate testCreate() {
        return Karate.run("custom-application-post").tags("validCustomApplication").relativeTo(getClass());
    }
}
