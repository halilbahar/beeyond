package at.htl.beeyond.resource;

import at.htl.beeyond.entity.CustomApplication;
import at.htl.beeyond.entity.User;

import javax.annotation.security.RolesAllowed;
import javax.inject.Inject;
import javax.transaction.Transactional;
import javax.validation.ConstraintViolation;
import javax.validation.Validator;
import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.SecurityContext;
import java.util.Set;

@Path("/application/custom")
@Consumes("application/json")
@Produces("application/json")
public class CustomApplicationResource {

    @Inject
    Validator validator;

    @POST
    @RolesAllowed({"student", "teacher"})
    @Transactional
    public Response create(@Context SecurityContext context, CustomApplication customApplication) {
        Set<ConstraintViolation<CustomApplication>> violations = this.validator.validate(customApplication);
        if (!violations.isEmpty()) {
            return Response.status(422).build();
        }

        User user = User.find("name", context.getUserPrincipal().getName()).firstResult();
        customApplication.setOwner(user);
        customApplication.persist();

        return Response.noContent().build();
    }
}
