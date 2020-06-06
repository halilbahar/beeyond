package at.htl.beeyond.validation;

import at.htl.beeyond.dto.TemplateApplicationDto;
import at.htl.beeyond.dto.TemplateFieldValueDto;
import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateField;
import org.hibernate.validator.constraintvalidation.HibernateConstraintValidatorContext;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import java.util.LinkedList;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;

public class TemplateFieldsCompleteValidator implements ConstraintValidator<TemplateFieldsComplete, TemplateApplicationDto> {

    private TemplateFieldsComplete constraint;

    public void initialize(TemplateFieldsComplete constraint) {
        this.constraint = constraint;
    }

    public boolean isValid(TemplateApplicationDto obj, ConstraintValidatorContext context) {
        if (obj == null) {
            return true;
        }

        Template template = Template.findById(obj.getTemplateId());

        if (template == null) {
            return true;
        }

        List<Long> fieldIds = template.getFields()
                .stream()
                .map(TemplateField::getId).collect(Collectors.toList());

        List<Long> dtoFieldValueIds = obj.getFieldValues()
                .stream()
                .map(TemplateFieldValueDto::getFieldId).collect(Collectors.toList());

        List<Long> missingIds = new LinkedList<>();
        for (Long id : fieldIds) {
            if (!dtoFieldValueIds.remove(id)) {
                missingIds.add(id);
            }
        }

        boolean isValid = missingIds.isEmpty() && dtoFieldValueIds.isEmpty();

        if (!isValid) {
            HibernateConstraintValidatorContext ctx = context.unwrap(HibernateConstraintValidatorContext.class);

            String missingIdString = missingIds.stream().map(Objects::toString).collect(Collectors.joining(", "));
            ctx.addMessageParameter("missing-ids", missingIdString);

            String obsoleteIdString = dtoFieldValueIds.stream().map(Objects::toString).collect(Collectors.joining(", "));
            ctx.addMessageParameter("obsolete-ids", obsoleteIdString);
        }

        return isValid;
    }
}
