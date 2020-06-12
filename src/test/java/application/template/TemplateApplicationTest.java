package application.template;

import com.intuit.karate.junit5.Karate;

public class TemplateApplicationTest {

    @Karate.Test
    Karate testCreate() {
        return Karate.run("template-application-post").tags("create").relativeTo(getClass());
    }

    @Karate.Test
    Karate testNoteLength() {
        return Karate.run("template-application-post").tags("noteLength").relativeTo(getClass());
    }
}
