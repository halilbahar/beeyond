package at.htl.beeyond.validation

import at.htl.beeyond.entity.User
import java.nio.charset.Charset
import javax.validation.ConstraintValidator
import javax.validation.ConstraintValidatorContext

class NamespaceValidValidator : ConstraintValidator<NamespaceValid, String> {
    override fun isValid(namespace: String?, context: ConstraintValidatorContext): Boolean {
        if (namespace == null) {
            return false
        }

        if (User.find<User>("name", namespace).firstResult<User>() != null ||
            !Charset.forName("US-ASCII").newEncoder().canEncode(namespace)
        ) {
            return false
        }

        return true
    }
}
