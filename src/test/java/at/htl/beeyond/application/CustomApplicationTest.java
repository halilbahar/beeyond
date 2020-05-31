package at.htl.beeyond.application;

import com.intuit.karate.junit5.Karate;

public class CustomApplicationTest {

    @Karate.Test
    Karate test001GetAll() {
        return Karate.run("custom-application-get").relativeTo(getClass());
    }

    @Karate.Test
    Karate test002Create() {
        return Karate.run("custom-application-post").tags("valid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test003InvalidCreate() {
        return Karate.run("custom-application-post").tags("invalid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test004Update() {
        return Karate.run("custom-application-put").tags("valid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test005InvalidUpdate() {
        return Karate.run("custom-application-put").tags("invalid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test006Delete() {
        return Karate.run("custom-application-delete").tags("valid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test007InvalidDelete() {
        return Karate.run("custom-application-delete").tags("invalid").relativeTo(getClass());
    }
}
