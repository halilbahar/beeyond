package at.htl.beeyond.resource

import at.htl.beeyond.dto.TemplateDto
import at.htl.beeyond.entity.Template
import at.htl.beeyond.entity.User
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import javax.annotation.security.RolesAllowed
import javax.transaction.Transactional
import javax.validation.Valid
import javax.ws.rs.*
import javax.ws.rs.core.*
import kotlin.streams.toList

@Path("/template")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class TemplateResource {

    @GET
    @RolesAllowed("student", "teacher")
    @Transactional
    fun getAllTemplates(): Response = Response.ok(Template.streamAll<Template>().map { TemplateDto(it) }.toList()).build()

    @POST
    @RolesAllowed("teacher")
    @Transactional
    fun createTemplate(
            @Context sc: SecurityContext,
            @Context uriInfo: UriInfo,
            @Valid templateDto: TemplateDto?
    ): Response {
        val owner = User.find<PanacheEntityBase>("name", sc.userPrincipal.name).firstResult<User>()
        val template = Template(templateDto, owner)
        template.persist()

        val uri = uriInfo.absolutePathBuilder.path(template.id.toString()).build()
        return Response.created(uri).build()
    }

    @GET
    @Path("/{id}")
    @RolesAllowed("student", "teacher")
    @Transactional
    fun getTemplateById(@PathParam("id") id: Long?): Response {
        val template = Template.findById<Template>(id)
                ?: return Response.status(Response.Status.NOT_FOUND).build()

        val templateDto = TemplateDto(template)
        return Response.ok(templateDto).build()
    }

    @DELETE
    @Path("/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun delete(@PathParam("id") id: Long?): Response {
        val template = Template.findById<Template>(id)
                ?: return Response.status(Response.Status.NOT_FOUND).build()

        template.deleted = true
        return Response.noContent().build()
    }
}
