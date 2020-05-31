package at.htl.beeyond.template;

import com.intuit.karate.junit5.Karate;

public class TemplateTest {

    @Karate.Test
    Karate test001GetAll() {
        return Karate.run("template-get").relativeTo(getClass());
    }

    @Karate.Test
    Karate test002Create() {
        return Karate.run("template-post").relativeTo(getClass());
    }

    @Karate.Test
    Karate test003Delete() {
        return Karate.run("template-delete").tags("valid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test004InvalidDelete() {
        return Karate.run("template-delete").tags("invalid").relativeTo(getClass());
    }
}
