package at.htl.beeyond.resource;

import at.htl.beeyond.model.Template;

import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/template")
@Consumes("application/json")
@Produces("application/json")
public class TemplateResource {

    @GET
    public Response getAll() {
        return Response.ok(Template.findAll().list()).build();
    }

    @POST
    @Transactional
    public Response create(Template template) {
        template.persist();
        return Response.noContent().build();
    }

    @DELETE
    @Path("/{id}")
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
