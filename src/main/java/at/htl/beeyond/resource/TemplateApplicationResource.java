package at.htl.beeyond.resource;

import at.htl.beeyond.model.Template;
import at.htl.beeyond.model.TemplateApplication;
import at.htl.beeyond.repository.TemplateApplicationRepository;
import at.htl.beeyond.repository.TemplateRepository;

import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/template-application")
@Transactional
public class TemplateApplicationResource {

    @Inject
    TemplateApplicationRepository templateApplicationRepository;
    @Inject
    TemplateRepository templateRepository;

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public Response getAllApplications() {
        return Response.ok(this.templateApplicationRepository.findAll().list()).build();
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/{id}")
    public Response getApplicationById(@PathParam("id") Long id) {
        TemplateApplication templateApplication = this.templateApplicationRepository.findById(id);
        if (templateApplication == null) {
            return Response.status(404).build();
        }

        return Response.ok(templateApplication).build();
    }

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    public Response uploadApplication(TemplateApplication templateApplication) {
        Template template = this.templateRepository
                .find("name", templateApplication.getTemplateName())
                .firstResult();
        templateApplication.setTemplate(template);
        this.templateApplicationRepository.persistApplication(templateApplication);
        return Response.noContent().build();
    }
}
