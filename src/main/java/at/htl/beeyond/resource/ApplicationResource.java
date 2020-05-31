package at.htl.beeyond.resource;

import at.htl.beeyond.entity.Application;
import at.htl.beeyond.entity.ApplicationStatus;
import at.htl.beeyond.service.DeploymentService;

import javax.annotation.security.RolesAllowed;
import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/application")
@Consumes("application/json")
@Produces("application/json")
public class ApplicationResource {

    @Inject
    DeploymentService deploymentService;

    @GET
    public Response getAll() {
        return Response.ok(Application.findAll().list()).build();
    }

    @PATCH
    @Path("/approve/{id}")
    @RolesAllowed("teacher")
    @Transactional
    public Response approve(@PathParam("id") Long id) {
        Application application = Application.findById(id);
        if (application == null) {
            return Response.status(404).build();
        }

        // Start application
        application.setStatus(ApplicationStatus.RUNNING);
        return Response.noContent().build();
    }

    @PATCH
    @Path("/deny/{id}")
    @RolesAllowed("teacher")
    @Transactional
    public Response deny(@PathParam("id") Long id) {
        Application application = Application.findById(id);
        if (application == null) {
            return Response.status(404).build();
        }

        application.setStatus(ApplicationStatus.DENIED);
        return Response.noContent().build();
    }
}
