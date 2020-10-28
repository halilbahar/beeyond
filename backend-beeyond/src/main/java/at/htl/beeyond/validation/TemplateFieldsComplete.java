package at.htl.beeyond.validation;

import javax.validation.Constraint;
import javax.validation.Payload;
import java.lang.annotation.*;

@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = TemplateFieldsCompleteValidator.class)
@Documented
public @interface TemplateFieldsComplete {

    String message() default "{at.htl.beeyond.validation.TemplateFieldsComplete.message}";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};
}
