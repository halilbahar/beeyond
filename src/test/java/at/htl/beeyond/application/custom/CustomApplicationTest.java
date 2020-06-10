package at.htl.beeyond.application.custom;

import com.intuit.karate.junit5.Karate;

public class CustomApplicationTest {

    @Karate.Test
    Karate testCreate() {
        return Karate.run("custom-application-post").tags("validCustomApplication").relativeTo(getClass());
    }

    @Karate.Test
    Karate testBlankContent() {
        return Karate.run("custom-application-post").tags("blankContent").relativeTo(getClass());
    }

    @Karate.Test
    Karate testBlankNote() {
        return Karate.run("custom-application-post").tags("blankNote").relativeTo(getClass());
    }

    @Karate.Test
    Karate testNoteLength() {
        return Karate.run("custom-application-post").tags("noteLength").relativeTo(getClass());
    }
}
