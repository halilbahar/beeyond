package at.htl.beeyond.resource;

import at.htl.beeyond.model.ApplicationStatus;
import at.htl.beeyond.model.CustomApplication;
import at.htl.beeyond.service.DeploymentService;

import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/custom-application")
@Consumes("application/json")
@Produces("application/json")
public class CustomApplicationResource {

    @Inject
    DeploymentService deploymentService;

    @GET
    public Response getAll() {
        return Response.ok(CustomApplication.findAll().list()).build();
    }

    @POST
    @Transactional
    public Response create(CustomApplication customApplication) {
        customApplication.persist();
        return Response.noContent().build();
    }

    @PUT
    @Path("/approve/{id}")
    @Transactional
    public Response approve(@PathParam("id") Long id) {
        CustomApplication customApplication = CustomApplication.findById(id);
        this.deploymentService.deploy(customApplication);
        customApplication.setStatus(ApplicationStatus.RUNNING);
        return Response.ok().build();
    }
}
