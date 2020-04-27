package at.htl.beeyond.resource;

import at.htl.beeyond.model.Template;
import at.htl.beeyond.repository.TemplateRepository;

import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/template")
public class TemplateResource {

    @Inject
    TemplateRepository templateRepository;

    @GET
    @Produces("application/json")
    public Response getAllTemplates() {
        return Response.ok(this.templateRepository.findAll().list()).build();
    }

    @POST
    @Consumes("application/json")
    @Transactional
    public Response createTemplate(Template template) {
        this.templateRepository.persist(template);
        return Response.noContent().build();
    }

    @DELETE
    @Path("/{id}")
    @Transactional
    public Response deleteTemplate(@PathParam("id") Long id) {
        Template template = this.templateRepository.findById(id);
        if (template == null) {
            return Response.status(404).build();
        }

        this.templateRepository.delete(template);
        return Response.noContent().build();
    }
}
