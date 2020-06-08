package at.htl.beeyond.application.template;

import com.intuit.karate.junit5.Karate;

public class TemplateApplicationTest {

    @Karate.Test
    Karate testCreate() {
        return Karate.run("template-application-post").relativeTo(getClass());
    }
}
