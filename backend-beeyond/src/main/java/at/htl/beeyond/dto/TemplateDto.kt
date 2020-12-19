package at.htl.beeyond.dto

import at.htl.beeyond.entity.Template
import at.htl.beeyond.entity.TemplateField
import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Checks.TemplateContent
import at.htl.beeyond.validation.TemplateFieldsMatching
import java.util.stream.Collectors
import javax.json.bind.annotation.JsonbTransient
import javax.validation.Valid
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull
import javax.validation.constraints.Size

@TemplateFieldsMatching(groups = [TemplateContent::class])
data class TemplateDto(
        @set:JsonbTransient var id: Long?,
        @field:NotNull @field:Size(min = 1, max = 255) var name: String?,
        @field:Size(max = 255) var description: String?,
        @field:NotBlank var content: String?,
        @field:Valid var fields: List<TemplateFieldDto>?
) {
    constructor() : this(null, null, null, null, null)

    override fun toString(): String {
        return ""
    }

    fun map(owner: User?): Template {
        val template = Template(name, description, content, owner)
        val templateFields = template.fields
        fields!!.stream()
                .map { fieldDto: TemplateFieldDto -> TemplateField(fieldDto.label, fieldDto.wildcard, fieldDto.description, template) }
                .forEach { e: TemplateField -> templateFields.add(e) }
        return template
    }

    companion object {
        @JvmStatic
        fun map(template: Template): TemplateDto {
            val fields = template.fields.stream()
                    .map { templateField: TemplateField? -> TemplateFieldDto.map(templateField) }
                    .collect(Collectors.toList())
            return TemplateDto(template.id, template.name, template.description, template.content, fields)
        }
    }
}
