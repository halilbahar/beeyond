package at.htl.beeyond.resource;

import at.htl.beeyond.model.TemplateApplication;
import at.htl.beeyond.repository.TemplateApplicationRepository;

import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/application")
@Transactional
public class TemplateApplicationResource {

    @Inject
    TemplateApplicationRepository templateApplicationRepository;

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
        this.templateApplicationRepository.persistApplication(templateApplication);
        return Response.noContent().build();
    }
}
