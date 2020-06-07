package at.htl.beeyond.dto;

import at.htl.beeyond.entity.TemplateField;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotBlank;

public class TemplateFieldDto {

    private Long id;

    @NotBlank
    @Length(max = 255)
    private String label;

    @NotBlank
    @Length(max = 255)
    private String wildcard;

    @Length(max = 255)
    private String description;

    public TemplateFieldDto(Long id, String label, String wildcard, String description) {
        this.id = id;
        this.label = label;
        this.wildcard = wildcard;
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
        this.label = label.trim();
    }

    public String getWildcard() {
        return wildcard;
    }

    public void setWildcard(String wildcard) {
        this.wildcard = wildcard.trim();
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description.trim();
    }

    public static TemplateFieldDto map(TemplateField templateField) {
        return new TemplateFieldDto(
                templateField.getId(),
                templateField.getLabel(),
                templateField.getWildcard(),
                templateField.getDescription()
        );
    }
}
