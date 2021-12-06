package at.htl.beeyond.resource

import at.htl.beeyond.dto.CustomApplicationDto
import at.htl.beeyond.dto.TemplateApplicationDto
import at.htl.beeyond.entity.*
import at.htl.beeyond.mailtemplate.GenericMail
import at.htl.beeyond.service.DeploymentService
import at.htl.beeyond.service.NamespaceService
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import io.quarkus.mailer.Mail
import io.quarkus.mailer.MailTemplate
import io.quarkus.mailer.Mailer
import io.quarkus.qute.Template
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

    @Inject
    lateinit var namespaceService: NamespaceService

    @Inject
    lateinit var mailer: Mailer

    @Inject
    lateinit var mail: Template

    @GET
    @RolesAllowed(value = ["student", "teacher"])
    @Transactional
    fun getAll(@Context ctx: SecurityContext): Response? {
        mailer.send(Mail.withHtml(
                "example@example.com",
                "Application #1124 approved",
                mail.data("content", GenericMail(
                        "Your application has been approved!",
                        "Application approved",
                        "The application for your project, 'SAMPLE_NAME', has been approved. You can now access stats and more - just visit the Beeyond dashboard."
                )).render()
        ))

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

        if(application.status == ApplicationStatus.PENDING){
            this.deploymentService.deploy(application)
            application.status = ApplicationStatus.RUNNING
            application.startedAt = LocalDateTime.now()

            application.namespace.users.forEach {
                val notification = Notification(it, "Application has been accepted!", NotificationStatus.POSITIVE, "application", application.id)
                notification.persist()
            }

            return Response.ok().build()
        }
        else{
            return Response.status(422).entity("Application is not in state "+ApplicationStatus.PENDING).build()
        }
    }

    @PATCH
    @Path("/deny/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun denyApplication(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(404).build()

        if(application.status == ApplicationStatus.PENDING){
            application.status = ApplicationStatus.DENIED

            application.namespace.users.forEach {
                val notification = Notification(it, "Application has been denied!", NotificationStatus.NEGATIVE, "application", application.id)
                notification.persist()
            }

            return Response.ok().build()
        }
        else{
            return Response.status(422).entity("Application is not in state "+ApplicationStatus.PENDING).build()
        }
    }

    @PATCH
    @Path("/stop/{id}")
    @RolesAllowed("teacher")
    @Transactional
    fun stopApplication(@PathParam("id") id: Long?): Response? {
        val application = Application.findById<Application>(id)
                ?: return Response.status(404).build()

        if(application.status == ApplicationStatus.RUNNING){

            deploymentService.stop(application)
            application.status = ApplicationStatus.FINISHED
            application.finishedAt = LocalDateTime.now()

            application.namespace.users.forEach {
                val notification = Notification(it, "Application has been stopped!", NotificationStatus.NEUTRAL, "application", application.id)
                notification.persist()
            }

            val isLastApplication = Application
                    .streamAll<Application>()
                    .filter {
                        it.status == ApplicationStatus.RUNNING && it.namespace == application.namespace
                    }.count() == 0L

            if (isLastApplication) {
                application.namespace.isDeleted = true
                namespaceService.deleteNamespace(application.namespace.namespace)
            }

            deploymentService.client.extensions().ingresses().withLabel("beeyond-application-id", application.id.toString()).delete()

            return Response.ok().build()
        } else{
            return Response.status(422).entity("Application is not in state "+ApplicationStatus.RUNNING).build()
        }

    }
}