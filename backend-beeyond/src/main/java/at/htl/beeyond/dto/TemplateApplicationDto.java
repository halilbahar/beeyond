package at.htl.beeyond.dto;

import at.htl.beeyond.entity.*;
import at.htl.beeyond.validation.Checks;
import at.htl.beeyond.validation.Exists;
import at.htl.beeyond.validation.TemplateFieldsComplete;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;


@TemplateFieldsComplete(groups = Checks.TemplateField.class)
public class TemplateApplicationDto extends ApplicationDto {

    @NotNull
    @Exists(entity = Template.class, fieldName = "id")
    private Long templateId;

    @Valid
    private List<TemplateFieldValueDto> fieldValues = new LinkedList<>();

    public TemplateApplicationDto(Long id, String note, ApplicationStatus status, UserDto owner, Long templateId, List<TemplateFieldValueDto> fieldValues) {
        super(id, note, status, owner);
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
        this.note = note.trim();
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

    @Override
    public String toString() {
        return "";
    }

    public TemplateApplication map(User owner) {
        Template template = Template.findById(this.templateId);
        List<TemplateFieldValue> templateFieldValues = this.fieldValues.stream()
                .map(o -> o.map(template))
                .collect(Collectors.toList());

        return new TemplateApplication(this.note, owner, template, templateFieldValues);
    }

    public static TemplateApplicationDto map(TemplateApplication templateApplication) {
        List<TemplateFieldValueDto> fieldValues = templateApplication.getFieldValues().stream()
                .map(TemplateFieldValueDto::map)
                .collect(Collectors.toList());

        return new TemplateApplicationDto(
                templateApplication.getId(),
                templateApplication.getNote(),
                templateApplication.getStatus(),
                UserDto.map(templateApplication.getOwner()),
                templateApplication.getTemplate().getId(),
                fieldValues
        );
    }
}
