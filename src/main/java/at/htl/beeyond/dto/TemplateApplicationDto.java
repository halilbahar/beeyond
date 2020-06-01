package at.htl.beeyond.dto;

import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateApplication;
import at.htl.beeyond.entity.TemplateFieldValue;
import at.htl.beeyond.entity.User;
import at.htl.beeyond.validation.Exists;
import org.hibernate.validator.constraints.Length;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

public class TemplateApplicationDto {

    private Long id;

    @Length(max = 255, message = "The note of the custom application cannot be longer than 255 characters")
    private String note;

    @NotNull(message = "The templateId of the template cannot be empty")
    @Exists(entity = Template.class, fieldName = "id", message = "The templateId of the template application does not exist")
    private Long templateId;

    @Valid
    private List<TemplateFieldValueDto> fieldValues = new LinkedList<>();

    public TemplateApplicationDto(Long id, String note, Long templateId, List<TemplateFieldValueDto> fieldValues) {
        this.id = id;
        this.note = note;
        this.templateId = templateId;
        this.fieldValues = fieldValues;
    }

    public TemplateApplicationDto() {
    }

    public Long getId() {
        return id;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }

    public Long getTemplateId() {
        return templateId;
    }

    public void setTemplateId(Long templateId) {
        this.templateId = templateId;
    }

    public List<TemplateFieldValueDto> getFieldValues() {
        return fieldValues;
    }

    public void setFieldValues(List<TemplateFieldValueDto> fieldValues) {
        this.fieldValues = fieldValues;
    }

    public TemplateApplication map(User owner) {
        Template template = Template.findById(templateId);
        List<TemplateFieldValue> templateFieldValues = fieldValues.stream()
                .map(o -> o.map(template))
                .collect(Collectors.toList());

        return new TemplateApplication(note, owner, template, templateFieldValues);
    }
}
