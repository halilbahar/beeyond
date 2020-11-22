package at.htl.beeyond.interceptor;

import at.htl.beeyond.entity.User;

import javax.enterprise.context.ApplicationScoped;
import javax.transaction.Transactional;
import javax.ws.rs.container.ContainerRequestContext;
import javax.ws.rs.container.ContainerRequestFilter;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.SecurityContext;
import javax.ws.rs.ext.Provider;
import java.security.Principal;

// Seems like @Transactional does only work in a managed bean => @ApplicationScoped is required
@Provider
@ApplicationScoped
public class UserFilter implements ContainerRequestFilter {

    @Context
    SecurityContext securityContext;

    @Override
    @Transactional
    public void filter(ContainerRequestContext containerRequestContext) {
        Principal userPrincipal = this.securityContext.getUserPrincipal();
        if (userPrincipal != null) {
            String name = userPrincipal.getName();
            User user = User.find("name", name).firstResult();

            if (user == null) {
                user = new User(name);
                user.persist();
            }
        }
    }
}
