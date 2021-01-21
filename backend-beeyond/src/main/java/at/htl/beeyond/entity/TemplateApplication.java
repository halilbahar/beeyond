package at.htl.beeyond.entity;

import at.htl.beeyond.dto.TemplateApplicationDto;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.ManyToOne;
import javax.persistence.OneToMany;
import java.util.List;

@Entity
public class TemplateApplication extends Application {

    @ManyToOne
    private Template template;

    @OneToMany(mappedBy = "templateApplication", cascade = CascadeType.PERSIST)
    private List<TemplateFieldValue> fieldValues;

    public TemplateApplication(String note, User owner, Template template, List<TemplateFieldValue> fieldValues) {
        super(note, owner);
        this.template = template;
        this.setFieldValues(fieldValues);
    }

    public TemplateApplication() {
    }

    public Template getTemplate() {
        return template;
    }

    public void setTemplate(Template template) {
        this.template = template;
    }

    public List<TemplateFieldValue> getFieldValues() {
        return fieldValues;
    }

    public void setFieldValues(List<TemplateFieldValue> fieldValues) {
        fieldValues.forEach(templateFieldValue -> templateFieldValue.setTemplateApplication(this));
        this.fieldValues = fieldValues;
    }

    public static TemplateApplicationDto getDto(TemplateApplication templateApplication) {
        return new TemplateApplicationDto(templateApplication);
    }
}
