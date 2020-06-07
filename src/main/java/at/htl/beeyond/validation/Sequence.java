package at.htl.beeyond.validation;

import javax.validation.GroupSequence;
import javax.validation.groups.Default;

public interface Sequence {
    @GroupSequence({Default.class, Checks.TemplateField.class})
    interface TemplateApplication {
    }

    @GroupSequence({Default.class, Checks.TemplateContent.class})
    interface Template {
    }
}
