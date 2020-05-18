package at.htl.beeyond.resource;

import at.htl.beeyond.entity.ApplicationStatus;
import at.htl.beeyond.entity.CustomApplication;
import at.htl.beeyond.service.DeploymentService;

import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.validation.ConstraintViolation;
import javax.validation.Validator;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;
import java.util.Set;

@Path("/custom-application")
@Consumes("application/json")
@Produces("application/json")
public class CustomApplicationResource {

    @Inject
    Validator validator;
    @Inject
    DeploymentService deploymentService;

    @GET
    @Transactional
    public Response getAll() {
        return Response.ok(CustomApplication.findAll().list()).build();
    }

    @POST
    @Transactional
    public Response create(CustomApplication customApplication) {
        Set<ConstraintViolation<CustomApplication>> violations = this.validator.validate(customApplication);
        if (!violations.isEmpty()) {
            return Response.status(422).build();
        }

        customApplication.persist();
        return Response.noContent().build();
    }

    @DELETE
    @Path("/{id}")
    @Transactional
    public Response delete(@PathParam("id") Long id) {
        CustomApplication customApplication = CustomApplication.findById(id);
        if (customApplication == null) {
            return Response.status(404).build();
        }

        customApplication.delete();
        return Response.ok(customApplication).build();
    }

    @PUT
    @Path("/approve/{id}")
    @Transactional
    public Response approve(@PathParam("id") Long id) {
        CustomApplication customApplication = CustomApplication.findById(id);
        if (customApplication == null) {
            Response.status(404).build();
        }

        this.deploymentService.deploy(customApplication);
        customApplication.setStatus(ApplicationStatus.RUNNING);
        return Response.noContent().build();
    }
}
