package at.htl.beeyond.validation;

import javax.validation.Constraint;
import javax.validation.Payload;
import java.lang.annotation.*;

@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = TemplateFieldsMatchingValidator.class)
@Documented
public @interface TemplateFieldsMatching {

    String message() default "{at.htl.beeyond.validation.TemplateFieldsMatching.message}";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};
}
