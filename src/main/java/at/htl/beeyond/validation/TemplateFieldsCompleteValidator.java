package at.htl.beeyond.validation;

import at.htl.beeyond.dto.TemplateApplicationDto;
import at.htl.beeyond.dto.TemplateFieldValueDto;
import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateField;

import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
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

        List<Long> fieldValueIds = obj.getFieldValues()
                .stream()
                .map(TemplateFieldValueDto::getFieldId).collect(Collectors.toList());

        for (Long id : fieldIds) {
            if (!fieldValueIds.remove(id)) {
                return false;
            }
        }

        boolean isValid = fieldValueIds.isEmpty();

        if (!isValid) {
            context.disableDefaultConstraintViolation();
            context.buildConstraintViolationWithTemplate("atwasdf")
                    .addPropertyNode("fieldValues").addConstraintViolation();
        }

        return isValid;
    }
}
