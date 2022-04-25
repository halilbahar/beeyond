package at.htl.beeyond.resource

import at.htl.beeyond.dto.ApplicationDto
import at.htl.beeyond.dto.CustomApplicationDto
import at.htl.beeyond.dto.DenyMessageDto
import at.htl.beeyond.dto.TemplateApplicationDto
import at.htl.beeyond.entity.*
import at.htl.beeyond.service.DeploymentService
import at.htl.beeyond.service.NamespaceService
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import java.time.LocalDateTime
import java.util.stream.Collectors
import javax.annotation.security.RolesAllowed
import javax.inject.Inject
import javax.json.JsonObject
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
    fun getAll(@Context ctx: SecurityContext, @QueryParam("all") all: Int): Response? {
        val mapToDto = { o: PanacheEntityBase? ->
            if (o is CustomApplication) {
                CustomApplicationDto(o)
            } else {
                TemplateApplicationDto(o as TemplateApplication)
            }
        }

        val applications = if (ctx.isUserInRole("teacher") && all == 1) {
            Application.streamAll<Application>().map(mapToDto).collect(Collectors.toList<Any>())
        } else {
            Application.streamAll<Application>().filter {
                it.namespace.users.map { it.name }.contains(ctx.userPrincipal.name)
            }.map(mapToDto).collect(Collectors.toList<Any>())
        }

        return Response.ok(applications).build()
    }

    @GET
    @Path("/{id}")
    @RolesAllowed(value = ["teacher", "student"])
    @Transactional
    fun getApplicationById(@PathParam("id") id: Long?, @Context ctx: SecurityContext): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(Response.Status.NOT_FOUND).build()

        if (!ctx.isUserInRole("teacher") && !application.namespace.users.map { it.name }
                .contains(ctx.userPrincipal.name)) {
            return Response.status(Response.Status.FORBIDDEN).build();
        }

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

        return if (application.status == ApplicationStatus.PENDING) {
            deploy(application)

            application.namespace.users.forEach {
                val notification = Notification(
                    it,
                    "Application has been accepted!",
                    "",
                    NotificationStatus.POSITIVE,
                    "application",
                    application.id
                )
                notification.persist()
            }

            Response.ok().build()
        } else {
            Response.status(422).entity("Application is not in state " + ApplicationStatus.PENDING).build()
        }
    }

    @PATCH
    @Path("/start/{id}")
    @RolesAllowed(value = ["teacher", "student"])
    @Transactional
    fun startApplication(@PathParam("id") id: Long?, @Context ctx: SecurityContext): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(Response.Status.NOT_FOUND).build()

        if (!ctx.isUserInRole("teacher") && !application.namespace.users.map { it.name }
                .contains(ctx.userPrincipal.name)) {
            return Response.status(Response.Status.FORBIDDEN).build();
        }

        return if (application.status == ApplicationStatus.STOPPED) {
            deploy(application)

            application.namespace.users.forEach {
                val notification = Notification(
                    it,
                    "Application has been started!",
                    "",
                    NotificationStatus.POSITIVE,
                    "application",
                    application.id
                )
                notification.persist()
            }

            Response.ok().build()
        } else {
            Response.status(422).entity("Application is not in state " + ApplicationStatus.STOPPED).build()
        }
    }

    @PATCH
    @Path("/{id}")
    @RolesAllowed(value = ["teacher", "student"])
    @Transactional
    fun saveApplication(@PathParam("id") id: Long?, @Context ctx: SecurityContext, json: JsonObject): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(Response.Status.NOT_FOUND).build()

        if (!ctx.isUserInRole("teacher") && !application.namespace.users.map { it.name }
                .contains(ctx.userPrincipal.name)) {
            return Response.status(Response.Status.FORBIDDEN).build();
        }

        return if ((application.status == ApplicationStatus.STOPPED ||
                    application.status == ApplicationStatus.DENIED ||
                    application.status == ApplicationStatus.PENDING) && json.containsKey("content")
        ) {
            application.content = json.getString("content")

            Response.ok().build()
        } else {
            Response.status(422)
                .entity("Application is not in state " + ApplicationStatus.STOPPED + "," + ApplicationStatus.PENDING + " or " + ApplicationStatus.DENIED)
                .build()
        }
    }

    private fun deploy(application: Application) {
        this.deploymentService.deploy(application)
        application.status = ApplicationStatus.RUNNING
        application.startedAt = LocalDateTime.now()
    }

    @PATCH
    @Path("/deny/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun denyApplication(@PathParam("id") id: Long?, message: DenyMessageDto?): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(404).build()

        return if (application.status == ApplicationStatus.PENDING) {
            application.status = ApplicationStatus.DENIED

            application.namespace.users.forEach {
                val notification = Notification(
                    it,
                    "Application has been denied!",
                    message?.message,
                    NotificationStatus.NEGATIVE,
                    "application",
                    application.id
                )
                notification.persist()
            }

            Response.ok().build()
        } else {
            Response.status(422).entity("Application is not in state " + ApplicationStatus.PENDING).build()
        }
    }

    @PATCH
    @Path("/stop/{id}")
    @RolesAllowed(value = ["teacher", "student"])
    @Transactional
    fun stopApplication(@PathParam("id") id: Long?, @Context ctx: SecurityContext): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(404).build()

        if (!ctx.isUserInRole("teacher") && !application.namespace.users.map { it.name }
                .contains(ctx.userPrincipal.name)) {
            return Response.status(Response.Status.FORBIDDEN).build();
        }

        return if (application.status == ApplicationStatus.RUNNING) {
            finishStopApplication(application, ApplicationStatus.STOPPED)
            Response.ok().build()
        } else {
            Response.status(422).entity("Application is not in state " + ApplicationStatus.RUNNING).build()
        }
    }

    @PATCH
    @Path("/finish/{id}")
    @RolesAllowed(value = ["teacher", "student"])
    @Transactional
    fun finishApplication(@PathParam("id") id: Long?, @Context ctx: SecurityContext): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(404).build()

        if (!ctx.isUserInRole("teacher") && !application.namespace.users.map { it.name }
                .contains(ctx.userPrincipal.name)) {
            return Response.status(Response.Status.FORBIDDEN).build();
        }

        return if (application.status == ApplicationStatus.RUNNING || application.status == ApplicationStatus.STOPPED) {
            finishStopApplication(application, ApplicationStatus.FINISHED)
            Response.ok().build()
        } else {
            Response.status(422)
                .entity("Application is not in state " + ApplicationStatus.RUNNING + " or " + ApplicationStatus.STOPPED)
                .build()
        }
    }

    @PATCH
    @Path("/request/{id}")
    @RolesAllowed(value = ["teacher", "student"])
    @Transactional
    fun requestApplication(@PathParam("id") id: Long?, @Context ctx: SecurityContext): Response? {
        val application = Application.findById<Application>(id)
            ?: return Response.status(404).build()

        if (!ctx.isUserInRole("teacher") && !application.namespace.users.map { it.name }
                .contains(ctx.userPrincipal.name)) {
            return Response.status(Response.Status.FORBIDDEN).build();
        }

        return if (application.status == ApplicationStatus.DENIED) {
            application.status = ApplicationStatus.PENDING
            Response.ok().build()
        } else {
            Response.status(422).entity("Application is not in state " + ApplicationStatus.DENIED).build()
        }
    }

    private fun finishStopApplication(application: Application, status: ApplicationStatus) {
        deploymentService.stop(application)
        application.status = status
        if (status == ApplicationStatus.FINISHED) {
            application.finishedAt = LocalDateTime.now()
        }

        application.namespace.users.forEach {
            val notification = Notification(
                it,
                "Application has been ${status.toString().lowercase()}!",
                "",
                NotificationStatus.NEUTRAL,
                "application",
                application.id
            )
            notification.persist()
        }
    }
}