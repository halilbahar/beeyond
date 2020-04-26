package at.htl.beeyond.resource;

import at.htl.beeyond.model.Application;
import at.htl.beeyond.repository.ApplicationRepository;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/approval")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
public class ApprovalResource {

    @Inject
    ApplicationRepository applicationRepository;

    @PUT
    @Path("/{id}")
    public Response approveOrDenyApplication(@PathParam("id") Long id, @QueryParam("approved") boolean isApproved) {
        Application application = applicationRepository.getApplicationById(id);

        if (application == null) {
            return Response.status(404).build();
        }
        return Response.ok(
                applicationRepository.approveOrDenyApplication(application, isApproved)
        ).build();
    }
}
