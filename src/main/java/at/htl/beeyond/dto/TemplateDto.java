package at.htl.beeyond.dto;

import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateField;
import at.htl.beeyond.entity.User;
import org.hibernate.validator.constraints.Length;

import javax.validation.Valid;
import javax.validation.constraints.NotBlank;
import java.util.List;
import java.util.stream.Collectors;

public class TemplateDto {

    private Long id;

    @NotBlank(message = "The name of the template cannot be blank")
    @Length(max = 255, message = "The name of the template cannot be longer than 255 characters")
    private String name;

    @Length(max = 255, message = "The description of the template cannot be longer than 255 characters")
    private String description;

    @NotBlank(message = "The content of the template cannot be blank")
    private String content;

    @Valid
    private List<TemplateFieldDto> fields;

    public TemplateDto(Long id, String name, String description, String content, List<TemplateFieldDto> fields) {
        this.id = id;
        this.name = name;
        this.description = description;
        this.content = content;
        this.fields = fields;
    }

    public TemplateDto() {
    }

    public Long getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public List<TemplateFieldDto> getFields() {
        return fields;
    }

    public void setFields(List<TemplateFieldDto> fields) {
        this.fields = fields;
    }

    public Template map(User owner) {
        Template template = new Template(name, description, content, owner);
        List<TemplateField> templateFields = template.getFields();
        fields.stream()
                .map(fieldDto -> new TemplateField(fieldDto.getLabel(), fieldDto.getWildcard(), fieldDto.getDescription(), template))
                .forEach(templateFields::add);

        return template;
    }

    public static TemplateDto map(Template template) {
        List<TemplateFieldDto> fields = template.getFields().stream()
                .map(TemplateFieldDto::map)
                .collect(Collectors.toList());

        return new TemplateDto(template.getId(), template.getName(), template.getDescription(), template.getContent(), fields);
    }
}
