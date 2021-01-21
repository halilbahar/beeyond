package at.htl.beeyond.resource

import at.htl.beeyond.dto.CustomApplicationDto
import at.htl.beeyond.dto.TemplateApplicationDto
import at.htl.beeyond.entity.Application
import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.TemplateApplication
import at.htl.beeyond.service.DeploymentService
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import java.util.stream.Collectors
import javax.annotation.security.RolesAllowed
import javax.inject.Inject
import javax.transaction.Transactional
import javax.ws.rs.*
import javax.ws.rs.core.Context
import javax.ws.rs.core.Response
import javax.ws.rs.core.SecurityContext

@Path("/application")
@Consumes("application/json")
@Produces("application/json")
class ApplicationResource {

    @Inject
    lateinit var deploymentService: DeploymentService

    @GET
    @RolesAllowed(value = ["student", "teacher"])
    @Transactional
    fun getAll(@Context ctx: SecurityContext): Response? {
        val mapToDto = { o: PanacheEntityBase? ->
            if (o is CustomApplication) {
                CustomApplicationDto(o)
            } else {
                TemplateApplicationDto(o as TemplateApplication)
            }
        }

        val applications = if (ctx.isUserInRole("teacher")) {
            Application.streamAll<PanacheEntityBase>().map(mapToDto).collect(Collectors.toList<Any>())
        } else {
            Application.streamAll<Application>().filter {
                it.owner.name == ctx.userPrincipal.name
            }.map(mapToDto).collect(Collectors.toList<Any>())
        }

        return Response.ok(applications).build()
    }

    @GET
    @Path("/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun getApplicationById(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(Response.Status.NOT_FOUND).build()
        return if (application is CustomApplication) {
            Response.ok(CustomApplicationDto(application)).build()
        } else {
            Response.ok(TemplateApplicationDto((application as TemplateApplication))).build()
        }
    }

    @PATCH
    @Path("/approve/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun approve(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(404).build()

        deploymentService.deploy(application)
        application.status = ApplicationStatus.RUNNING
        return Response.noContent().build()
    }

    @PATCH
    @Path("/deny/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun deny(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(404).build()

        application.status = ApplicationStatus.DENIED
        return Response.noContent().build()
    }
}
