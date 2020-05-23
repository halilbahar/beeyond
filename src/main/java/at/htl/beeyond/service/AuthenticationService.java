package at.htl.beeyond.service;

import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;

import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Form;

@RegisterRestClient
@Produces("application/json")
public interface AuthenticationService {

    @POST
    @Consumes("application/x-www-form-urlencoded")
    Object login(Form loginForm);
}
