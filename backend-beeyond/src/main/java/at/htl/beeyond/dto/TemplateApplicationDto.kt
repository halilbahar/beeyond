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
import java.time.LocalDateTime


@GroupSequence(value = [TemplateApplicationDto::class, Checks.TemplateField::class])
@TemplateFieldsComplete(groups = [Checks.TemplateField::class])
class TemplateApplicationDto(
        id: Long? = null,
        note: String? = null,
        status: ApplicationStatus? = null,
        owner: UserDto? = null,
        createdAt: LocalDateTime? = null,
        @field:NotNull @field:Exists(entity = Template::class, fieldName = "id") var templateId: Long? = null,
        fieldValues: List<TemplateFieldValue> = LinkedList()
) : ApplicationDto(
        id,
        note,
        status,
        owner,
        createdAt
) {
    @field:Valid
    var fieldValues: List<TemplateFieldValueDto> = fieldValues.stream().map { TemplateFieldValueDto(it) }.collect(Collectors.toList())

    constructor(templateApplication: TemplateApplication) : this(
            templateApplication.id,
            templateApplication.note,
            templateApplication.status,
            UserDto(templateApplication.owner),
            templateApplication.createdAt,
            templateApplication.template.id,
            templateApplication.fieldValues
    )

    override fun toString(): String {
        return ""
    }

    fun map(owner: User?): TemplateApplication {
        val template = Template.findById<Template>(templateId)
        return TemplateApplication(note, owner, template, fieldValues.map { TemplateFieldValue(it) }.toList())


//        val templateApplication = TemplateApplication.findById<TemplateApplication>(id)
//        val fieldValues = fieldValues.stream()
//                .map{ fieldDto -> TemplateFieldValue(fieldDto.value, TemplateField.findById(fieldDto.fieldId), templateApplication) }
//                .collect(Collectors.toList())
//        return TemplateApplication(note, owner, template, fieldValues)
    }
}
