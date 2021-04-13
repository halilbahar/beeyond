package at.htl.beeyond.resource

import at.htl.beeyond.dto.UserListDto
import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.entity.User
import at.htl.beeyond.service.NamespaceService
import at.htl.beeyond.validation.NamespaceValid
import org.hibernate.validator.constraints.Length
import javax.inject.Inject
import javax.json.Json
import javax.json.JsonArrayBuilder
import javax.json.JsonObject
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

    @GET
    @Path("")
    fun getNamespaces(): Response {
        val user = Json.createObjectBuilder()
        user.add("id", 1)
        user.add("name", "user-1")

        val json = Json.createObjectBuilder()
        json.add("namespace", "namespace-1")
        json.add("id", 1)
        json.add("users", Json.createArrayBuilder().add(user).build())

        json.add("namespace", "namespace-2")
        json.add("id", 2)
        json.add("users", Json.createArrayBuilder().add(user).build())

        return Response.ok(listOf(json.build())).build();
    }

    @PUT
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

        userList.users.forEach {
            User.find<User>("name", it).firstResult<User>()?.namespaces?.add(namespace)
        }

        return Response.noContent().build()
    }
}
