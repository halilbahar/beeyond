package at.htl.beeyond.dto

import at.htl.beeyond.entity.Template
import at.htl.beeyond.entity.TemplateField
import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Checks.TemplateContent
import at.htl.beeyond.validation.TemplateFieldsMatching
import java.util.*
import java.util.stream.Collectors
import javax.json.bind.annotation.JsonbTransient
import javax.validation.GroupSequence
import javax.validation.Valid
import javax.validation.constraints.NotBlank
import javax.validation.constraints.NotNull
import javax.validation.constraints.Size

@TemplateFieldsMatching(groups = [TemplateContent::class])
class TemplateDto(
        @set:JsonbTransient var id: Long? = null,
        @field:NotNull @field:Size(min = 1, max = 255) var name: String? = null,
        @field:Size(max = 255) var description: String? = null,
        @field:NotBlank var content: String? = null,
        @field:Valid var fields: List<TemplateFieldDto> = LinkedList(),
        var deleted: Boolean?
) {

    constructor(template: Template):this(template.id, template.name, template.description, template.content, template.fields.stream().map { TemplateFieldDto(it) }!!.collect(Collectors.toList()), template.deleted)

    constructor():this(null, null, null, null, LinkedList(), null)

    override fun toString(): String {
        return ""
    }

    fun map(owner: User?): Template {
        val template = Template(name, description, content, owner, deleted)
        val templateFields = template.fields
        fields.stream()
                .map{ fieldDto -> TemplateField(fieldDto.label, fieldDto.wildcard, fieldDto.description, template) }
                .forEach { e: TemplateField -> templateFields.add(e) }
        return template
    }
}
