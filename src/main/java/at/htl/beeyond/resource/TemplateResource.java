package at.htl.beeyond.resource;

import at.htl.beeyond.service.DeploymentService;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Path("/template")
public class TemplateResource {

    @Inject
    DeploymentService deploymentService;

    @POST
    @Consumes(MediaType.APPLICATION_JSON)
    @Produces(MediaType.APPLICATION_JSON)
    @Path("/nginx")
    public Response hello(@QueryParam("replica") Integer replica) {
        if (replica == null || replica <= 0) {
            return Response.status(400).build();
        }
        this.deploymentService.deployNginx(replica);
        return Response.noContent().build();
    }
}
