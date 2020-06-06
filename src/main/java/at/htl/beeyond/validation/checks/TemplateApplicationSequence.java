package at.htl.beeyond.validation.checks;

import javax.validation.GroupSequence;
import javax.validation.groups.Default;

@GroupSequence({Default.class, TemplateFieldChecks.class})
public interface TemplateApplicationSequence {
}
