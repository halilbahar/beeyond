package at.htl.beeyond.dto

import at.htl.beeyond.entity.TemplateField
import at.htl.beeyond.entity.TemplateFieldValue
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.Exists
import javax.json.bind.annotation.JsonbTransient
import javax.validation.GroupSequence
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull

@GroupSequence(value = [Checks.TemplateFieldExists::class, TemplateFieldValueDto::class])
class TemplateFieldValueDto(
        @set:JsonbTransient var id: Long? = null,
        value: String? = null,
        @field:Exists(
                entity = TemplateField::class,
                fieldName = "id",
                groups = [Checks.TemplateFieldExists::class]
        )
        @field:NotNull var fieldId: Long? = null
) {
    constructor(templateFieldValue: TemplateFieldValue) : this(
            templateFieldValue.id,
            templateFieldValue.value,
            templateFieldValue.field.id
    )

    @field:NotBlank
    var value: String? = value
        set(value) {
            if (value != null) {
                field = value.trim()
            }
        }
}
