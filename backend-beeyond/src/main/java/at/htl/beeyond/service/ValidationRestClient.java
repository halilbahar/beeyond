package at.htl.beeyond.service;

import at.htl.beeyond.dto.FailedFieldDto;
import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import java.util.Set;

@RegisterRestClient
public interface ValidationRestClient {

    @POST
    @Produces({MediaType.APPLICATION_JSON})
    @Path("/api/validate")
    Set<FailedFieldDto> validateKubernetesYaml(String content);
}
