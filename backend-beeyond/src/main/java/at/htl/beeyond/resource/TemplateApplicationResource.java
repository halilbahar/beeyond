package at.htl.beeyond.resource;

import at.htl.beeyond.dto.TemplateApplicationDto;
import at.htl.beeyond.entity.TemplateApplication;
import at.htl.beeyond.entity.User;

import javax.annotation.security.RolesAllowed;
import javax.transaction.Transactional;
import javax.validation.Valid;
import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.SecurityContext;

@Path("/application/template")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class TemplateApplicationResource {

    @POST
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response create(@Context SecurityContext context, @Valid TemplateApplicationDto templateApplicationDto) {
        User user = User.find("name", context.getUserPrincipal().getName()).firstResult();
        TemplateApplication templateApplication = templateApplicationDto.map(user);

        templateApplication.persist();

        return Response.noContent().build();
    }
}
