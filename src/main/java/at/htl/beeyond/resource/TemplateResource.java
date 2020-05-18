package at.htl.beeyond.resource;

import at.htl.beeyond.entity.Template;

import javax.annotation.security.RolesAllowed;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/template")
@Consumes("application/json")
@Produces("application/json")
public class TemplateResource {

    @GET
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response getAll() {
        return Response.ok(Template.findAll().list()).build();
    }

    @POST
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response create(Template template) {
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
