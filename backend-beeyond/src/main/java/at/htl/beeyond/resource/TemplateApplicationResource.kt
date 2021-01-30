package at.htl.beeyond.resource

import at.htl.beeyond.dto.TemplateApplicationDto
import at.htl.beeyond.entity.Template
import at.htl.beeyond.entity.TemplateApplication
import at.htl.beeyond.entity.User
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
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

@Path("/application/template")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class TemplateApplicationResource {

    @POST
    @RolesAllowed("student", "teacher")
    @Transactional
    fun createTemplateApplication(
            @Context context: SecurityContext,
            @Valid templateApplicationDto: TemplateApplicationDto?
    ): Response {
        val template = Template.findById<Template>(templateApplicationDto!!.templateId)
        if (template.deleted) {
            return Response.status(Response.Status.NOT_FOUND).build()
        }

        val owner = User.find<PanacheEntityBase>("name", context.userPrincipal.name).firstResult<User>()
        val templateApplication = TemplateApplication(templateApplicationDto, owner)
        templateApplication.persist()
        return Response.noContent().build()
    }
}
