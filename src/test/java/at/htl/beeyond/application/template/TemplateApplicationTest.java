package at.htl.beeyond.application.template;

import com.intuit.karate.junit5.Karate;

public class TemplateApplicationTest {

    @Karate.Test
    Karate test001Upload() {
        return Karate.run("template-application-post").relativeTo(getClass());
    }

    @Karate.Test
    Karate test003GetAll() {
        return Karate.run("template-application-get").tags("all").relativeTo(getClass());
    }

    @Karate.Test
    Karate test004GetById() {
        return Karate.run("template-application-get").tags("validId").relativeTo(getClass());
    }

    @Karate.Test
    Karate test005GetByNotExistingId() {
        return Karate.run("template-application-get").tags("invalidId").relativeTo(getClass());
    }
}
