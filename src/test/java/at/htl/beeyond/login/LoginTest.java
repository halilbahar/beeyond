package at.htl.beeyond.login;

import com.intuit.karate.junit5.Karate;

public class LoginTest {

    @Karate.Test
    Karate test001Login() {
        return Karate.run("login").relativeTo(getClass());
    }
}
