package at.htl.beeyond.dto

import at.htl.beeyond.entity.Template
import at.htl.beeyond.validation.Checks.TemplateContent
import at.htl.beeyond.validation.TemplateFieldsMatching
import java.util.*
import javax.json.bind.annotation.JsonbTransient
import javax.validation.GroupSequence
import javax.validation.Valid
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull
import javax.validation.constraints.Size

@GroupSequence(TemplateDto::class, TemplateContent::class)
@TemplateFieldsMatching(groups = [TemplateContent::class])
class TemplateDto(
        @set:JsonbTransient var id: Long? = null,
        @field:NotNull @field:Size(min = 1, max = 255) var name: String? = null,
        @field:Size(max = 255) var description: String? = null,
        @field:NotBlank var content: String? = null,
        @field:Valid var fields: List<TemplateFieldDto> = LinkedList(),
        @set:JsonbTransient var deleted: Boolean? = null
) {
    constructor(template: Template) : this(
            template.id,
            template.name,
            template.description,
            template.content,
            template.fields.map { TemplateFieldDto(it) }.toList(),
            template.deleted
    )

    override fun toString(): String {
        return ""
    }
}
