package at.htl.beeyond.resource;

import at.htl.beeyond.service.NamespaceService;

import javax.inject.Inject;
import javax.ws.rs.DELETE;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.QueryParam;
import javax.ws.rs.core.Response;

@Path("/namespace")
public class NamespaceResource {

    @Inject
    NamespaceService namespaceService;

    @POST
    public Response createNamespace(@QueryParam("namespace") String namespace) {
        this.namespaceService.createNamespace(namespace);
        return Response.ok("Namespace " + namespace + " created").build();
    }

    @DELETE
    public Response deleteNamespace(@QueryParam("namespace") String namespace) {
        this.namespaceService.deleteNamespace(namespace);
        return Response.ok("Namespace " + namespace + " deleted").build();
    }
}
