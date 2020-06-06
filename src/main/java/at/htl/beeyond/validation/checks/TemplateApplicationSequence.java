package at.htl.beeyond.validation.checks;

import javax.validation.GroupSequence;
import javax.validation.groups.Default;

@GroupSequence({Default.class, TemplateFieldChecks.class, TemplateFieldsCompleteChecks.class})
public interface TemplateApplicationSequence {
}
