package at.htl.beeyond.dto

import at.htl.beeyond.entity.*
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.Exists
import at.htl.beeyond.validation.TemplateFieldsComplete
import at.htl.beeyond.validation.ValidKubernetes
import java.time.LocalDate
import java.time.LocalDateTime
import java.util.*
import javax.validation.GroupSequence
import javax.validation.Valid
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull

@GroupSequence(TemplateApplicationDto::class, Checks.TemplateField::class, Checks.KubernetesContent::class)
@TemplateFieldsComplete(groups = [Checks.TemplateField::class])
@ValidKubernetes(groups = [Checks.KubernetesContent::class])
class TemplateApplicationDto(
    id: Long? = null,
    note: String? = null,
    status: ApplicationStatus? = null,
    owner: UserDto? = null,
    createdAt: LocalDateTime? = null,
    startedAt: LocalDateTime? = null,
    finishedAt: LocalDateTime? = null,
    @field:NotNull @field:Exists(entity = Template::class, fieldName = "id") var templateId: Long? = null,
    @field:Valid var fieldValues: List<TemplateFieldValueDto> = LinkedList(),
    namespace: String = "",
    schoolClass: String? = null,
    toDate: LocalDate? = null,
    purpose: String? = null,
    content: String? = null
) : ApplicationDto(
    id,
    note,
    status,
    owner,
    createdAt,
    startedAt,
    finishedAt,
    namespace,
    schoolClass,
    toDate,
    purpose
) {

    constructor(templateApplication: TemplateApplication) : this(
        templateApplication.id,
        templateApplication.note,
        templateApplication.status,
        UserDto(templateApplication.owner),
        templateApplication.createdAt,
        templateApplication.startedAt,
        templateApplication.finishedAt,
        templateApplication.template.id,
        templateApplication.fieldValues.map { TemplateFieldValueDto(it) }.toList(),
        templateApplication.namespace.namespace,
        templateApplication.schoolClass,
        templateApplication.toDate,
        templateApplication.purpose
    )

    override fun toString(): String {
        return ""
    }

    @field:NotBlank
    var content: String? = content
        set(value) {
            if (value != null) {
                field = value.trim()
            }
        }
}
