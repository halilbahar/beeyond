package at.htl.beeyond.resource;

import at.htl.beeyond.dto.CustomApplicationDto;
import at.htl.beeyond.dto.TemplateApplicationDto;
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
import javax.ws.rs.core.Response.Status;
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
        List<Object> applications = Application.streamAll().map(o -> {
            if (o instanceof CustomApplication) {
                return CustomApplication.getDto((CustomApplication) o);
            } else if (o instanceof TemplateApplication) {
                return TemplateApplication.getDto((TemplateApplication) o);
            }
            return null;
        }).collect(Collectors.toList());
        return Response.ok(applications).build();
    }

    @GET
    @Path("/{id}")
    @RolesAllowed("teacher")
    @Transactional
    public Response getApplicationById(@PathParam("id") Long id) {
        Application application = Application.findById(id);
        if (application == null) {
            return Response.status(Status.NOT_FOUND).build();
        }

        if (application instanceof CustomApplication) {
            return Response.ok(CustomApplicationDto.map((CustomApplication) application)).build();
        } else {
            return Response.ok(TemplateApplicationDto.map((TemplateApplication) application)).build();
        }
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
