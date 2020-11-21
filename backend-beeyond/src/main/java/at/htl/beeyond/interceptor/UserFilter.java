package at.htl.beeyond.interceptor;

import at.htl.beeyond.entity.User;

import javax.enterprise.context.Dependent;
import javax.inject.Inject;
import javax.persistence.EntityManager;
import javax.transaction.Transactional;
import javax.ws.rs.container.ContainerRequestContext;
import javax.ws.rs.container.ContainerRequestFilter;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.SecurityContext;
import javax.ws.rs.ext.Provider;
import java.security.Principal;

@Provider
public class UserFilter implements ContainerRequestFilter {

    // @Transactional not working without EntityManager
    @Inject
    EntityManager entityManager;

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
                user.persistAndFlush();
            }
        }
    }
}
