package at.htl.beeyond.resource;

import at.htl.beeyond.model.TemplateApplication;
import at.htl.beeyond.repository.TemplateApplicationRepository;
import at.htl.beeyond.service.DeploymentService;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/approval")
public class ApprovalResource {

    @Inject
    TemplateApplicationRepository templateApplicationRepository;
    @Inject
    DeploymentService deploymentService;

    @PUT
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/{id}")
    public Response approveOrDenyApplication(@PathParam("id") Long id, @QueryParam("approved") boolean isApproved) {
        TemplateApplication templateApplication = this.templateApplicationRepository.findById(id);
        if (templateApplication == null) {
            return Response.status(404).build();
        }

        this.templateApplicationRepository.setApproval(id, isApproved);
        if (isApproved) {
            this.deploymentService.deployNginx(templateApplication.getReplica());
        }

        return Response.noContent().build();
    }
}
