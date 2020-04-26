package at.htl.beeyond.resource;

import at.htl.beeyond.service.TemplateService;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Response;
import java.util.List;

@Path("/template")
public class TemplateResource {

    @Inject
    TemplateService templateService;

    @GET
    @Produces("application/json")
    public Response getAllTemplates() {
        return Response.ok(templateService.getAllTemplates()).build();
    }
}
