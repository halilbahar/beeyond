package at.htl.beeyond.integration.application;

import at.htl.beeyond.integration.DatabaseResource;
import com.intuit.karate.junit5.Karate;
import io.quarkus.test.common.QuarkusTestResource;
import io.quarkus.test.junit.QuarkusTest;
import io.quarkus.test.security.TestSecurity;

@QuarkusTest
@QuarkusTestResource(DatabaseResource.class)
public class ApplicationTest {
    @Karate.Test
    @TestSecurity(user = "sonja-teacher", roles = "teacher")
    Karate createTeacherApplication() {
        return Karate.run("create-applications").tags("teacher").relativeTo(getClass());
    }

    @Karate.Test
    @TestSecurity(user = "moritz", roles = "student")
    Karate createStudentApplication() {
        return Karate.run("create-applications").tags("student").relativeTo(getClass());
    }

    @Karate.Test
    @TestSecurity(user = "sonja-teacher", roles = "teacher")
    Karate testGetAllAsTeacher() {
        return Karate.run("application-get").tags("teacher").relativeTo(getClass());
    }

    @Karate.Test
    @TestSecurity(user = "moritz", roles = "student")
    Karate testGetAllAsStudent() {
        return Karate.run("application-get").tags("student").relativeTo(getClass());
    }

    @Karate.Test
    Karate testApproveAsTeacher() {
        return Karate.run("application-approve").tags("teacher").relativeTo(getClass());
    }

    @Karate.Test
    Karate testApproveAsStudent() {
        return Karate.run("application-approve").tags("student").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDenyAsTeacher() {
        return Karate.run("application-deny").tags("teacher").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDenyAsStudent() {
        return Karate.run("application-deny").tags("teacher").relativeTo(getClass());
    }
}
