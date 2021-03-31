package at.htl.beeyond.validation;

import javax.validation.Constraint;
import javax.validation.Payload;
import java.lang.annotation.*;

@Target(ElementType.PARAMETER)
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = NamespaceValidValidator.class)
@Documented
public @interface NamespaceValid {

    String message() default "{at.htl.beeyond.validation.NamespaceValid.message}";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};
}
