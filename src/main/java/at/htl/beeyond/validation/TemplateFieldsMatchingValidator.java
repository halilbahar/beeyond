package at.htl.beeyond.validation;

import at.htl.beeyond.dto.TemplateDto;
import at.htl.beeyond.dto.TemplateFieldDto;
import org.hibernate.validator.constraintvalidation.HibernateConstraintValidatorContext;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import java.util.LinkedList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class TemplateFieldsMatchingValidator implements ConstraintValidator<TemplateFieldsMatching, TemplateDto> {

    public boolean isValid(TemplateDto obj, ConstraintValidatorContext context) {
        if (obj == null) {
            return true;
        }

        List<TemplateFieldDto> fields = new LinkedList<>(obj.getFields());
        String content = obj.getContent();

        Pattern pattern = Pattern.compile("%([\\w]+)%");
        Matcher matcher = pattern.matcher(content);

        List<String> wildcardsInContent = new LinkedList<>();
        while (matcher.find()) {
            wildcardsInContent.add(matcher.group(1));
        }

        List<String> obsoleteWildCards = new LinkedList<>();
        for (TemplateFieldDto field : fields) {
            if (!wildcardsInContent.remove(field.getWildcard())) {
                obsoleteWildCards.add(field.getWildcard());
            }
        }

        boolean isValid = obsoleteWildCards.isEmpty() && wildcardsInContent.isEmpty();

        if (!isValid) {
            HibernateConstraintValidatorContext ctx = context.unwrap(HibernateConstraintValidatorContext.class);

            String missingWildCardsString = String.join(", ", wildcardsInContent);
            String obseleteWildCardsString = String.join(", ", obsoleteWildCards);
            ctx.addMessageParameter("missing-wildcards", missingWildCardsString);
            ctx.addMessageParameter("obsolete-wildcards", obseleteWildCardsString);
        }

        return isValid;
    }
}
