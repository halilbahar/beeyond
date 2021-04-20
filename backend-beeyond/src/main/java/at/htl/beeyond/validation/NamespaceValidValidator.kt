package at.htl.beeyond.validation

import at.htl.beeyond.entity.User
import javax.validation.ConstraintValidator
import javax.validation.ConstraintValidatorContext

class NamespaceValidValidator : ConstraintValidator<NamespaceValid, String> {
    override fun isValid(namespace: String?, context: ConstraintValidatorContext): Boolean {
        if (namespace == null) {
            return true
        }

        return User.find<User>("name", namespace).firstResult<User>() == null
    }
}
