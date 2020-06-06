package at.htl.beeyond.dto;

import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateField;
import at.htl.beeyond.entity.User;
import at.htl.beeyond.validation.TemplateFieldsMatching;
import at.htl.beeyond.validation.checks.TemplateContentCheck;
import org.hibernate.validator.constraints.Length;

import javax.validation.Valid;
import javax.validation.constraints.NotBlank;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

@TemplateFieldsMatching(groups = TemplateContentCheck.class)
public class TemplateDto {

    private Long id;

    @NotBlank
    @Length(max = 255)
    private String name;

    @Length(max = 255)
    private String description;

    @NotBlank
    private String content;

    @Valid
    private List<TemplateFieldDto> fields = new LinkedList<>();

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
        this.name = name.trim();
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description.trim();
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content.trim();
    }

    public List<TemplateFieldDto> getFields() {
        return fields;
    }

    public void setFields(List<TemplateFieldDto> fields) {
        this.fields = fields;
    }

    @Override
    public String toString() {
        return "";
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
