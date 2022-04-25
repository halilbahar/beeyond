package at.htl.beeyond.interceptor

import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.entity.User
import org.eclipse.microprofile.config.inject.ConfigProperty
import javax.enterprise.context.ApplicationScoped
import javax.transaction.Transactional
import javax.ws.rs.container.ContainerRequestContext
import javax.ws.rs.container.ContainerRequestFilter
import javax.ws.rs.core.Context
import javax.ws.rs.core.SecurityContext
import javax.ws.rs.ext.Provider

@Provider
@ApplicationScoped
class UserFilter : ContainerRequestFilter {

    @Context
    lateinit var securityContext: SecurityContext

    @ConfigProperty(name = "beeyond.namespace.prefix")
    lateinit var namespacePrefix: String

    @Transactional
    override fun filter(requestContext: ContainerRequestContext?) {
        val userPrincipal = this.securityContext.userPrincipal
        if (userPrincipal != null) {
            val name = userPrincipal.name
            var user: User? = User.find<User>("name", name).firstResult()

            if (user == null) {
                try {
                    user = User(name)
                    user.persist()

                    val namespaceName = namespacePrefix + "-" + name.split('@')[0].replace('.', '-');
                    val namespace = Namespace(namespaceName)
                    namespace.isDefault = true
                    namespace.users = listOf(user)
                    namespace.persist()
                } catch (e: Exception) {
                } // Prevent duplicate user and namespace
            }
        }
    }
}
