package at.htl.beeyond.resource

import at.htl.beeyond.dto.NamespaceDto
import at.htl.beeyond.dto.UserListDto
import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.entity.User
import at.htl.beeyond.service.NamespaceService
import at.htl.beeyond.validation.NamespaceValid
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import org.hibernate.validator.constraints.Length
import java.util.*
import java.util.stream.Collectors
import javax.inject.Inject
import javax.transaction.Transactional
import javax.validation.Valid
import javax.ws.rs.*
import javax.ws.rs.core.*

@Path("/namespace")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class NamespaceResource {

    @Inject
    lateinit var namespaceService: NamespaceService

    @GET
    fun getNamespaces(@Context ctx: SecurityContext, @QueryParam("all") all: Int): Response {
        val mapToDto = { o: PanacheEntityBase? ->
            NamespaceDto(o as Namespace)
        }

        val namespaces = if (ctx.isUserInRole("teacher") && all == 1) {
            Namespace.streamAll<Namespace>().map(mapToDto).collect(Collectors.toList<Any>())
        } else {
            Namespace.streamAll<Namespace>().filter {
                it.users.contains(User.find<User>("name", ctx.userPrincipal.name).firstResult<User>())
            }.map(mapToDto).collect(Collectors.toList<Any>())
        }

        return Response.ok(namespaces).build()
    }

    @GET
    @Path("/{namespace}")
    fun getNamespace(
        @PathParam("namespace") namespaceName: String
    ): Response {
        val namespace = Namespace.find<Namespace>("namespace", namespaceName).firstResultOptional<Namespace>();

        return if(namespace.isEmpty){
            Response.status(Response.Status.NOT_FOUND).build()
        } else{
            Response.ok(NamespaceDto(namespace.get())).build()
        }
    }

    @PUT
    @Path("/{namespace}")
    @Transactional
    fun assignNamespace(
        @Context uriInfo: UriInfo,
        @PathParam("namespace") @Length(min = 1, max = 50) @NamespaceValid namespaceName: String,
        @Valid userList: UserListDto
    ): Response {
        var namespace = Namespace.find<Namespace>("namespace", namespaceName).firstResult<Namespace>()

        if (namespace == null) {
            if (namespaceService.allNamespaces.find { it.metadata.name == namespaceName } == null) {
                namespaceService.createNamespace(namespaceName)
            }

            namespace = Namespace(namespaceName)
            Namespace.persist(namespace)
        }

        namespace.users = userList.users.distinct().map {
            User.find<User>("name", it).firstResult<User>()
        }

        if (namespace.users.isEmpty()) {
            namespaceService.deleteNamespace(namespaceName)
            Namespace.deleteById(namespace.id)
        }

        return Response.noContent().build()
    }
}
