package at.htl.beeyond.dto;

import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.TemplateField;
import at.htl.beeyond.entity.TemplateFieldValue;
import at.htl.beeyond.validation.checks.TemplateFieldChecks;

import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;

public class TemplateFieldValueDto {

    private Long id;

    @NotBlank(groups = TemplateFieldChecks.class)
    private String value;

    @NotNull(groups = TemplateFieldChecks.class)
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
        this.value = value.trim();
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
