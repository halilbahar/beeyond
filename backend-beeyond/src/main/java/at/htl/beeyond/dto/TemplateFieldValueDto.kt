package at.htl.beeyond.dto

import at.htl.beeyond.entity.TemplateFieldValue
import at.htl.beeyond.entity.Template
import at.htl.beeyond.entity.TemplateApplication
import at.htl.beeyond.entity.TemplateField
import at.htl.beeyond.validation.Checks
import javax.json.bind.annotation.JsonbTransient
import javax.validation.GroupSequence
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull

@GroupSequence(value = [TemplateFieldValueDto::class, Checks.TemplateField::class])
class TemplateFieldValueDto(
        @set:JsonbTransient var id: Long? = null,
        value: String? = null,
        @field:NotNull(groups = [Checks.TemplateField::class]) var fieldId: Long? = null
) {
    constructor(templateFieldValue: TemplateFieldValue) : this(
            templateFieldValue.id,
            templateFieldValue.value,
            templateFieldValue.field.id
    )

    @field:NotBlank(groups = [Checks.TemplateField::class])
    var value: String? = value
        set(value) {
            if (value != null) {
                field = value.trim()
            }
        }

    fun map(templateApplication: TemplateApplication?): TemplateFieldValue {
        val templateField = TemplateField.findById<TemplateField>(fieldId)
        return TemplateFieldValue(value, templateField, templateApplication)
    }
}
