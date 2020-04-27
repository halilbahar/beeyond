package at.htl.beeyond.resource;

import at.htl.beeyond.service.NamespaceService;

import javax.inject.Inject;
import javax.ws.rs.*;
import javax.ws.rs.core.Response;

@Path("/namespace")
public class NamespaceResource {

    @Inject
    NamespaceService namespaceService;

    @GET
    public Response getAllNamespaces() {
        return Response.ok(this.namespaceService.getAllNamespaces()).build();
    }

    @POST
    public Response createNamespace(@QueryParam("namespace") String namespace) {
        this.namespaceService.createNamespace(namespace);
        return Response.noContent().build();
    }

    @DELETE
    public Response deleteNamespace(@QueryParam("namespace") String namespace) {
        this.namespaceService.deleteNamespace(namespace);
        return Response.noContent().build();
    }
}
