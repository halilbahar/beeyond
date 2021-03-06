package at.htl.beeyond.validation;

import at.htl.beeyond.dto.ApplicationDto;
import at.htl.beeyond.dto.CustomApplicationDto;
import at.htl.beeyond.dto.TemplateApplicationDto;
import at.htl.beeyond.service.ValidationRestClient;
import org.eclipse.microprofile.rest.client.inject.RestClient;

import javax.inject.Inject;
import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;

public class ValidKubernetesValidator implements ConstraintValidator<ValidKubernetes, ApplicationDto> {

    @Inject
    @RestClient
    ValidationRestClient validationRestClient;

    @Override
    public boolean isValid(ApplicationDto applicationDto, ConstraintValidatorContext context) {
        if (applicationDto == null) {
            return true;
        }

        String content;

        if (applicationDto instanceof CustomApplicationDto) {
            content = ((CustomApplicationDto) applicationDto).getContent();
        } else {
            content = ((TemplateApplicationDto) applicationDto).getContent();
        }
        var result = this.validationRestClient.validateKubernetesYaml(content);
        if (result != null) {
            context.disableDefaultConstraintViolation();
            for (var error:result) {
                context.buildConstraintViolationWithTemplate(error.getMessage())
                        .addPropertyNode(error.getKey())
                        .addConstraintViolation();
            }
        }

        return result == null;
    }
}
