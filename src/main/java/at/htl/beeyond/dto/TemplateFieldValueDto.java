package at.htl.beeyond.dto;

import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateField;
import at.htl.beeyond.entity.TemplateFieldValue;
import at.htl.beeyond.validation.Exists;
import at.htl.beeyond.validation.checks.TemplateFieldChecks;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;

public class TemplateFieldValueDto {

    private Long id;

    @NotBlank(message = "The value of the fieldvalue cannot be blank", groups = TemplateFieldChecks.class)
    private String value;

    @NotNull(message = "The fieldId of the fieldvalue cannot be empty", groups = TemplateFieldChecks.class)
    @Exists(entity = TemplateField.class, fieldName = "id", message = "The fieldId of the fieldvalue does not exist", groups = TemplateFieldChecks.class)
    private Long fieldId;

    public TemplateFieldValueDto(Long id, String value, Long fieldId) {
        this.id = id;
        this.value = value;
        this.fieldId = fieldId;
    }

    public TemplateFieldValueDto() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public Long getFieldId() {
        return fieldId;
    }

    public void setFieldId(Long fieldId) {
        this.fieldId = fieldId;
    }

    public TemplateFieldValue map(Template template) {
        TemplateField templateField = TemplateFieldValue.findById(this.fieldId);
        return new TemplateFieldValue(this.value, templateField, template);
    }
}
