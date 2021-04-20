package at.htl.beeyond.resource

import at.htl.beeyond.dto.UserDto
import at.htl.beeyond.entity.User
import javax.ws.rs.Consumes
import javax.ws.rs.GET
import javax.ws.rs.Path
import javax.ws.rs.Produces
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response

@Path("/user")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class UserResource {

    @GET
    fun getAllUsers(): Response {
        val users = User.streamAll<User>().map { UserDto(it) }.toArray()
        return Response.ok(users).build()
    }
}
