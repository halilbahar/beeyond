package at.htl.beeyond;

import at.htl.beeyond.service.DeploymentService;

import javax.inject.Inject;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.Response;

@Path("/template")
public class TemplateResource {

    @Inject
    DeploymentService deploymentService;

    @POST
    @Path("/nginx")
    public Response hello(@QueryParam("replica") Integer replica) {
        if (replica == null || replica <= 0) {
            return Response.status(400).build();
        }
        this.deploymentService.deployNginx(replica);
        return Response.noContent().build();
    }
}