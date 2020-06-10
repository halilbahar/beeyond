package at.htl.beeyond.template;

import com.intuit.karate.junit5.Karate;

public class TemplateTest {

    @Karate.Test
    Karate testGetAll() {
        return Karate.run("template-get").relativeTo(getClass());
    }

    @Karate.Test
    Karate testCreate() {
        return Karate.run("template-post").tags("create").relativeTo(getClass());
    }

    @Karate.Test
    Karate testBlankContent() {
        return Karate.run("template-post").tags("blankContent").relativeTo(getClass());
    }

    @Karate.Test
    Karate testNameLength() {
        return Karate.run("template-post").tags("nameLength").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDelete() {
        return Karate.run("template-delete").tags("deleteExistingTemplate").relativeTo(getClass());
    }

    @Karate.Test
    Karate testDeleteNotExisting() {
        return Karate.run("template-delete").tags("deleteNotExistingTemplate").relativeTo(getClass());
    }
}
