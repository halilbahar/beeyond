package at.htl.beeyond.resource;

import at.htl.beeyond.model.Application;
import at.htl.beeyond.repository.ApplicationRepository;
import at.htl.beeyond.service.DeploymentService;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/approval")
public class ApprovalResource {

    @Inject
    ApplicationRepository applicationRepository;

    @Inject
    DeploymentService deploymentService;

    @PUT
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/{id}")
    public Response approveOrDenyApplication(@PathParam("id") Long id, @QueryParam("approved") boolean isApproved) {
        Application application = applicationRepository.getApplicationById(id);

        if (application == null) {
            return Response.status(404).build();
        }
        Application processedApplication = applicationRepository.approveOrDenyApplication(application, isApproved);
        if (processedApplication.getApproved()) {
            deploymentService.deployNginx(processedApplication.getReplica());
        }
        return Response.ok(processedApplication).build();
    }
}
