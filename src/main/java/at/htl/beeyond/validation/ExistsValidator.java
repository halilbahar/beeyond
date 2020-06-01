package at.htl.beeyond.validation;

import io.quarkus.hibernate.orm.panache.runtime.JpaOperations;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;

public class ExistsValidator implements ConstraintValidator<Exists, Object> {

    private Exists exists;

    @Override
    public void initialize(Exists constraintAnnotation) {
        this.exists = constraintAnnotation;
    }

    @Override
    public boolean isValid(Object value, ConstraintValidatorContext context) {
        if (value == null) {
            return true;
        }

        return JpaOperations.count(this.exists.entity(), this.exists.fieldName(), value) != 0;
    }
}
