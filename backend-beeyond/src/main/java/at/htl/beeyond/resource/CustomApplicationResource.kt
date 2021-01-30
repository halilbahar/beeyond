package at.htl.beeyond.resource

import at.htl.beeyond.dto.CustomApplicationDto
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.User
import javax.annotation.security.RolesAllowed
import javax.transaction.Transactional
import javax.validation.Valid
import javax.ws.rs.Consumes
import javax.ws.rs.POST
import javax.ws.rs.Path
import javax.ws.rs.Produces
import javax.ws.rs.core.Context
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response
import javax.ws.rs.core.SecurityContext

@Path("/application/custom")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class CustomApplicationResource {

    @POST
    @RolesAllowed("student", "teacher")
    @Transactional
    fun createCustomApplication(
            @Context context: SecurityContext,
            @Valid customApplicationDto: CustomApplicationDto?
    ): Response {
        val owner = User.find<User>("name", context.userPrincipal.name).firstResult<User>()
        val customApplication = CustomApplication(customApplicationDto, owner)

        customApplication.persist()
        return Response.noContent().build()
    }
}
