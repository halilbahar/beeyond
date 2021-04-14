package at.htl.beeyond.resource

import at.htl.beeyond.dto.CustomApplicationDto
import at.htl.beeyond.dto.TemplateApplicationDto
import at.htl.beeyond.entity.Application
import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.TemplateApplication
import at.htl.beeyond.service.DeploymentService
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import java.time.LocalDateTime
import java.util.stream.Collectors
import javax.annotation.security.RolesAllowed
import javax.inject.Inject
import javax.transaction.Transactional
import javax.ws.rs.*
import javax.ws.rs.core.Context
import javax.ws.rs.core.MediaType
import javax.ws.rs.core.Response
import javax.ws.rs.core.SecurityContext

@Path("/application")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
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
            Application.streamAll<Application>().map(mapToDto).collect(Collectors.toList<Any>())
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
    fun approveApplication(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(Response.Status.NOT_FOUND).build()

        this.deploymentService.deploy(application)
        application.status = ApplicationStatus.RUNNING
        application.startedAt = LocalDateTime.now()

        return Response.noContent().build()
    }

    @PATCH
    @Path("/deny/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun denyApplication(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(404).build()

        application.status = ApplicationStatus.DENIED
        return Response.noContent().build()
    }

    @PATCH
    @Path("/stop/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun stopApplication(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(404).build()

        deploymentService.stop(application)
        application.status = ApplicationStatus.FINISHED
        application.finishedAt = LocalDateTime.now()

        return Response.noContent().build()
    }
}
