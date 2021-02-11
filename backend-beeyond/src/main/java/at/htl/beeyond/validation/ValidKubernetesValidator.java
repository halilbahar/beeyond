package at.htl.beeyond.validation;

import at.htl.beeyond.dto.ApplicationDto;
import at.htl.beeyond.dto.CustomApplicationDto;
import at.htl.beeyond.dto.TemplateApplicationDto;
import at.htl.beeyond.service.ValidationRestClient;
import org.eclipse.microprofile.rest.client.inject.RestClient;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import javax.validation.ConstraintValidator;
import javax.validation.ConstraintValidatorContext;

@ApplicationScoped
public class ValidKubernetesValidator implements ConstraintValidator<ValidKubernetes, ApplicationDto> {

    @Inject
    @RestClient
    ValidationRestClient validationRestClient;

    public boolean isValid(ApplicationDto applicationDto, ConstraintValidatorContext constraintValidatorContext) {
        if (applicationDto == null) {
            return true;
        }

        String content;

        if (applicationDto instanceof CustomApplicationDto) {
            content = ((CustomApplicationDto) applicationDto).getContent();
        } else {
            content = ((TemplateApplicationDto) applicationDto).getContent();
        }

        return this.validationRestClient.validateKubernetesYaml(content).getStatus() == 200;
    }
}
