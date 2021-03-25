package at.htl.beeyond.resource

import at.htl.beeyond.dto.UserListDto
import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.entity.TemplateField
import at.htl.beeyond.entity.User
import at.htl.beeyond.service.NamespaceService
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.Exists
import at.htl.beeyond.validation.NamespaceValid
import org.hibernate.validator.constraints.Length
import javax.inject.Inject
import javax.transaction.Transactional
import javax.validation.Valid
import javax.ws.rs.*
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response

@Path("/namespace")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class NamespaceResource {

    @Inject
    lateinit var namespaceService: NamespaceService

    @POST
    @Path("/{namespace}")
    @Transactional
    fun assignNamespace(
        @PathParam("namespace") @Length(min = 1, max = 50) @NamespaceValid namespaceName: String,
        @Valid userList: UserListDto
    ): Response {
        var namespace = Namespace.find<Namespace>("namespace", namespaceName).firstResult<Namespace>()

        if (namespace == null) {
            namespaceService.createNamespace(namespaceName)
            namespace = Namespace(namespaceName)
            Namespace.persist(namespace)
        }

        userList.users?.forEach {
            User.find<User>("name", it).firstResult<User>()?.namespaces?.add(namespace);
        }

        return Response.noContent().build();
    }
}
