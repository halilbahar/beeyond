package at.htl.beeyond.entity;

import at.htl.beeyond.dto.TemplateFieldValueDto;
import io.quarkus.hibernate.orm.panache.PanacheEntityBase;

import javax.persistence.*;

@Entity
public class TemplateFieldValue extends PanacheEntityBase {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String value;

    @ManyToOne
    private TemplateField field;

    @ManyToOne
    private TemplateApplication templateApplication;

    public TemplateFieldValue(String value, TemplateField field, TemplateApplication templateApplication) {
        this.value = value;
        this.field = field;
        this.templateApplication = templateApplication;
    }

    public TemplateFieldValue(TemplateFieldValueDto templateFieldValueDto) {
        this.value = templateFieldValueDto.getValue();
        this.field = TemplateField.findById(templateFieldValueDto.getFieldId());
    }

    public TemplateFieldValue() {
    }

    public Long getId() {
        return id;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }

    public TemplateField getField() {
        return field;
    }

    public void setField(TemplateField field) {
        this.field = field;
    }

    public TemplateApplication getTemplateApplication() {
        return templateApplication;
    }

    public void setTemplateApplication(TemplateApplication templateApplication) {
        this.templateApplication = templateApplication;
    }
}
