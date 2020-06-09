package at.htl.beeyond.application;

import com.intuit.karate.junit5.Karate;

public class ApplicationTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("application-get").relativeTo(getClass());
    }
}
