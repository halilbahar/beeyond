package at.htl.beeyond.resource;

import at.htl.beeyond.entity.Application;
import at.htl.beeyond.entity.ApplicationStatus;
import at.htl.beeyond.entity.CustomApplication;
import at.htl.beeyond.entity.TemplateApplication;
import at.htl.beeyond.service.DeploymentService;

import javax.annotation.security.RolesAllowed;
import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.List;
import java.util.stream.Collectors;

@Path("/application")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class ApplicationResource {

    @Inject
    DeploymentService deploymentService;

    @GET
    @Transactional
    public Response getAll() {
        List<Object> applications = Application.findAll().stream().map(o -> {
            if (o instanceof CustomApplication) {
                return CustomApplication.getDto((CustomApplication) o);
            } else if (o instanceof TemplateApplication) {
                return TemplateApplication.getDto((TemplateApplication) o);
            }
            return null;
        }).collect(Collectors.toList());
        return Response.ok(applications).build();
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

        this.deploymentService.deploy(application);
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
