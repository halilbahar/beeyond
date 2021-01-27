package at.htl.beeyond.resource;

import at.htl.beeyond.dto.TemplateDto;
import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.User;

import javax.annotation.security.RolesAllowed;
import javax.transaction.Transactional;
import javax.validation.Valid;
import javax.ws.rs.*;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;
import javax.ws.rs.core.SecurityContext;
import javax.ws.rs.core.UriInfo;
import java.net.URI;

@Path("/template")
@Consumes("application/json")
@Produces("application/json")
public class TemplateResource {

    @GET
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response getAll() {
        return Response.ok(Template.getDtos()).build();
    }

    @POST
    @RolesAllowed("teacher")
    @Transactional
    public Response create(@Context SecurityContext sc, @Context UriInfo uriInfo, @Valid TemplateDto templateDto) {
        User owner = User.find("name", sc.getUserPrincipal().getName()).firstResult();
        Template template = templateDto.map(owner);
        template.persist();
        template.setDeleted(false);

        URI uri = uriInfo.getAbsolutePathBuilder().path(Long.toString(template.getId())).build();

        return Response.created(uri).build();
    }

    @GET
    @Path("/{id}")
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response getTemplateById(@PathParam("id") Long id) {
        Template template = Template.findById(id);
        if (template == null) {
            return Response.status(Status.NOT_FOUND).build();
        }

        TemplateDto templateDto = new TemplateDto(template);
        return Response.ok(templateDto).build();
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

        template.setDeleted(true);
        return Response.noContent().build();
    }
}
