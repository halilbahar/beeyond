package at.htl.beeyond.resource

import at.htl.beeyond.dto.NotificationDto
import at.htl.beeyond.entity.Notification
import javax.annotation.security.RolesAllowed
import java.util.stream.Collectors
import javax.ws.rs.*
import javax.ws.rs.core.Context
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response
import javax.ws.rs.core.SecurityContext

@Path("/notification")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class NotificationResource {
    @GET
    @RolesAllowed("student", "teacher")
    fun getNamespaces(@Context ctx: SecurityContext): Response {
        val mapToDto = { o: Notification -> NotificationDto(o) }

        val notifications = Notification.streamAll<Notification>().filter {
            it.user.name == ctx.userPrincipal.name
        }.map(mapToDto).collect(Collectors.toList<Any>())

        return Response.ok(notifications).build()
    }
}
