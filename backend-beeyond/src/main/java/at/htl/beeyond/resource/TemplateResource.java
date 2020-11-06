package at.htl.beeyond.resource;

import at.htl.beeyond.dto.TemplateDto;
import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.User;
import at.htl.beeyond.model.FailedField;
import at.htl.beeyond.service.ValidationService;
import at.htl.beeyond.validation.Sequence;

import javax.annotation.security.RolesAllowed;
import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.SecurityContext;
import java.util.List;

@Path("/template")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class TemplateResource {

    @Inject
    ValidationService validationService;

    @GET
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response getAll() {
        return Response.ok(Template.getDtos()).build();
    }

    @GET
    @Path("/{id}")
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response getById(@PathParam("id") Long id) {
        Template template = Template.findById(id);
        if (template == null) {
            return Response.status(404).build();
        }
        return Response.ok(Template.getDto(id)).build();
    }

    @POST
    @RolesAllowed("teacher")
    @Transactional
    public Response create(@Context SecurityContext sc, TemplateDto templateDto) {
        List<FailedField> failedFields = this.validationService.validate(templateDto, Sequence.Template.class);
        if (!failedFields.isEmpty()) {
            return Response.status(422).entity(failedFields).build();
        }

        User owner = User.find("name", sc.getUserPrincipal().getName()).firstResult();
        Template template = templateDto.map(owner);
        template.persist();

        return Response.noContent().build();
    }

    @DELETE
    @Path("/{id}")
    @RolesAllowed("teacher")
    @Transactional
    public Response delete(@PathParam("id") Long id) {
        Template template = Template.findById(id);
        if (template == null) {
            return Response.status(404).build();
        }

        template.delete();
        return Response.noContent().build();
    }
}
