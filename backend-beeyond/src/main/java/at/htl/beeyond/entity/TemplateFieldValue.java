package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntityBase;

import javax.persistence.*;

@Entity
public class TemplateFieldValue extends PanacheEntityBase {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String value;

    @OneToOne
    private TemplateField field;

    @ManyToOne
    private Template template;

    public TemplateFieldValue(String value, TemplateField field, Template template) {
        this.value = value;
        this.field = field;
        this.template = template;
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

    public Template getTemplate() {
        return template;
    }

    public void setTemplate(Template template) {
        this.template = template;
    }
}
