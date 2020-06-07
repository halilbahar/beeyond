package at.htl.beeyond.service;

import at.htl.beeyond.model.FailedField;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import javax.validation.ConstraintViolation;
import javax.validation.Validator;
import javax.validation.groups.Default;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

@ApplicationScoped
public class ValidationService {

    @Inject
    Validator validator;

    public List<FailedField> validate(Object object) {
        return this.validate(object, Default.class);
    }

    public List<FailedField> validate(Object object, Class<?> clazz) {
        return this.validator.validate(object, clazz).stream()
                .map(o -> new FailedField(
                        o.getPropertyPath().toString(),
                        o.getInvalidValue() != null ? o.getInvalidValue().toString() : "",
                        o.getMessage()
                )).collect(Collectors.toList());
    }
}
