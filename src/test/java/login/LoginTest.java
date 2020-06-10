package login;

import com.intuit.karate.junit5.Karate;

public class LoginTest {

    @Karate.Test
    Karate testLogin() {
        return Karate.run("login").relativeTo(getClass());
    }
}
