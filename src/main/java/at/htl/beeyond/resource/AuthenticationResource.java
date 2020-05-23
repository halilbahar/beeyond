package at.htl.beeyond.resource;

import at.htl.beeyond.model.LoginData;
import at.htl.beeyond.service.AuthenticationService;
import org.eclipse.microprofile.config.inject.ConfigProperty;
import org.eclipse.microprofile.rest.client.inject.RestClient;

import javax.annotation.security.PermitAll;
import javax.inject.Inject;
import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Form;


@Path("/authentication")
@Consumes("application/json")
@Produces("application/json")
public class AuthenticationResource {

    @Inject
    @RestClient
    AuthenticationService authenticationService;

    @ConfigProperty(name = "beeyond.keycloak-client-id")
    String clientId;

    @ConfigProperty(name = "beeyond.keycloak-secret")
    String clientSecret;

    @POST
    @Path("/login")
    @PermitAll
    public Object login(LoginData loginData) {
        Form form = new Form()
                .param("grant_type", "password")
                .param("client_id", this.clientId)
                .param("client_secret", this.clientSecret)
                .param("username", loginData.getUsername())
                .param("password", loginData.getPassword());

        return this.authenticationService.login(form);
    }

}
