package at.htl.beeyond.resource;

import at.htl.beeyond.dto.TemplateApplicationDto;
import at.htl.beeyond.entity.TemplateApplication;
import at.htl.beeyond.entity.User;
import at.htl.beeyond.model.FailedField;
import at.htl.beeyond.service.ValidationService;

import javax.annotation.security.RolesAllowed;
import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.SecurityContext;
import java.util.List;

@Path("/application/template")
@Consumes("application/json")
@Produces("application/json")
public class TemplateApplicationResource {

    @Inject
    ValidationService validationService;

    @POST
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response create(@Context SecurityContext context, TemplateApplicationDto templateApplicationDto) {
        List<FailedField> failedFields = this.validationService.validate(templateApplicationDto);
        if (!failedFields.isEmpty()) {
            return Response.status(422).entity(failedFields).build();
        }

        User user = User.find("name", context.getUserPrincipal().getName()).firstResult();
        TemplateApplication templateApplication = templateApplicationDto.map(user);

        templateApplication.persist();

        return Response.noContent().build();
    }
}
