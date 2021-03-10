package at.htl.beeyond.resource

import at.htl.beeyond.dto.TemplateDto
import at.htl.beeyond.entity.Template
import at.htl.beeyond.entity.User
import at.htl.beeyond.service.DeploymentService
import io.fabric8.kubernetes.api.model.apiextensions.v1.CustomResourceDefinitionBuilder
import io.quarkus.hibernate.orm.panache.PanacheEntityBase
import javax.annotation.security.RolesAllowed
import javax.inject.Inject
import javax.transaction.Transactional
import javax.validation.Valid
import javax.ws.rs.*
import javax.ws.rs.core.*
import kotlin.streams.toList

@Path("/template")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
class UserResource {

    @Inject
    lateinit var deploymentYamlService: DeploymentService

    @GET
    @RolesAllowed("student", "teacher")
    @Transactional
    fun getOwn() {

    }
}
