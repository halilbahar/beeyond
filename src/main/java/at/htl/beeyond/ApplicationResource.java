package at.htl.beeyond;

import at.htl.beeyond.repository.ApplicationRepository;

import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/application")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
@Transactional
public class ApplicationResource {

    @Inject
    ApplicationRepository applicationRepository;

    @GET
    public Response getAllApplications() {
        return Response.ok(
                applicationRepository.getAllApplications()
        ).build();
    }
}
