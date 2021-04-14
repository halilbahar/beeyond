package at.htl.beeyond.service;

import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;

import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.core.Response;

@RegisterRestClient
public interface ValidationRestClient {

    @POST
    @Path("/api/validate")
    Response validateKubernetesYaml(String content);
}
