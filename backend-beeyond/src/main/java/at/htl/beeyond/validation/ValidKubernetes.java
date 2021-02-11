package at.htl.beeyond.validation;

import javax.validation.Constraint;
import javax.validation.Payload;
import java.lang.annotation.*;

@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Constraint(validatedBy = ValidKubernetesValidator.class)
@Documented
public @interface ValidKubernetes {

    String message() default "{at.htl.beeyond.validation.ValidKubernetes.message}";

    Class<?>[] groups() default {};

    Class<? extends Payload>[] payload() default {};
}
