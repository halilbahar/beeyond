package at.htl.beeyond.dto;

import at.htl.beeyond.entity.TemplateField;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotBlank;

public class TemplateFieldDto {

    private Long id;

    @NotBlank(message = "The label of the field cannot be blank")
    @Length(max = 255, message = "The label of the field cannot be longer than 255 characters")
    private String label;

    @Length(max = 255, message = "The description of the field cannot be longer than 255 characters")
    private String description;

    public TemplateFieldDto(Long id, String label, String description) {
        this.id = id;
        this.label = label;
        this.description = description;
    }

    public TemplateFieldDto() {
    }

    public Long getId() {
        return id;
    }

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public static TemplateFieldDto map(TemplateField templateField) {
        return new TemplateFieldDto(templateField.getId(), templateField.getLabel(), templateField.getDescription());
    }
}
