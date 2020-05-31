package at.htl.beeyond.resource;

import javax.annotation.security.RolesAllowed;
import javax.transaction.Transactional;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.core.Response;

@Path("/template")
@Transactional
public class TemplateApplicationResource {

    @POST
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response test() {
        return null;
    }
}
