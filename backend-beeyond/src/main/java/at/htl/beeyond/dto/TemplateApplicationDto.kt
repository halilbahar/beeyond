package at.htl.beeyond.dto

import at.htl.beeyond.entity.*
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.Exists
import at.htl.beeyond.validation.TemplateFieldsComplete
import java.util.*
import java.util.stream.Collectors
import javax.validation.GroupSequence
import javax.validation.Valid
import javax.validation.constraints.NotNull
import at.htl.beeyond.entity.TemplateFieldValue
import at.htl.beeyond.validation.ValidKubernetes
import java.time.LocalDateTime

@GroupSequence(value = [TemplateApplicationDto::class, Checks.TemplateField::class, Checks.KubernetesContent::class])
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
        @field:Valid var fieldValues: List<TemplateFieldValueDto> = LinkedList()
) : ApplicationDto(
        id,
        note,
        status,
        owner,
        createdAt,
        startedAt,
        finishedAt
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
            templateApplication.fieldValues.map { TemplateFieldValueDto(it) }.toList()
    )

    fun getContent(): String {
        val template = Template.findById<Template>(this.templateId)
        val fieldValues = this.fieldValues
        var content = template.content
        for (fieldValue in fieldValues) {
            val wildcard = TemplateField.findById<TemplateField>(fieldValue.fieldId).wildcard
            content = content.replace("%$wildcard%", fieldValue.value!!)
        }

        return content;
    }

    override fun toString(): String {
        return ""
    }
}
