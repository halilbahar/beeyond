package at.htl.beeyond.application.custom;

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
    Karate test003InvalidContent() { return Karate.run("custom-application-post").tags("invalidContent").relativeTo(getClass()); }

    @Karate.Test
    Karate test004Approve() {
        return Karate.run("custom-application-put").tags("validApprove").relativeTo(getClass());
    }

    @Karate.Test
    Karate test005InvalidApprove() { return Karate.run("custom-application-put").tags("invalidApprove").relativeTo(getClass()); }

    @Karate.Test
    Karate test006Deny() {
        return Karate.run("custom-application-put").tags("validDeny").relativeTo(getClass());
    }

    @Karate.Test
    Karate test007InvalidDeny() { return Karate.run("custom-application-put").tags("invalidDeny").relativeTo(getClass()); }

    @Karate.Test
    Karate test008Delete() {
        return Karate.run("custom-application-delete").tags("valid").relativeTo(getClass());
    }

    @Karate.Test
    Karate test009InvalidDelete() { return Karate.run("custom-application-delete").tags("invalid").relativeTo(getClass()); }
}
