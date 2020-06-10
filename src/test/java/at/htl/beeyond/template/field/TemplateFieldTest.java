package at.htl.beeyond.template.field;

import com.intuit.karate.junit5.Karate;

public class TemplateFieldTest {

    @Karate.Test
    Karate testLabelLength() {
        return Karate.run("template-field-post").tags("labelLength").relativeTo(getClass());
    }
}
