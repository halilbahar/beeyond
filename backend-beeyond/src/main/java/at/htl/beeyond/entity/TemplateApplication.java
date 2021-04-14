package at.htl.beeyond.entity;

import at.htl.beeyond.dto.TemplateApplicationDto;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.ManyToOne;
import javax.persistence.OneToMany;
import java.util.List;
import java.util.stream.Collectors;

@Entity
public class TemplateApplication extends Application {

    @ManyToOne
    private Template template;

    @OneToMany(mappedBy = "templateApplication", cascade = CascadeType.PERSIST)
    private List<TemplateFieldValue> fieldValues;

    public TemplateApplication(TemplateApplicationDto templateApplicationDto, User owner) {
        super(templateApplicationDto.getNote(), owner);
        this.template = Template.findById(templateApplicationDto.getTemplateId());
        List<TemplateFieldValue> templateFieldValues = templateApplicationDto.getFieldValues()
                .stream()
                .map(TemplateFieldValue::new)
                .collect(Collectors.toList());

        this.setFieldValues(templateFieldValues);
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

    @Override
    public String getContent() {
        List<TemplateFieldValue> fieldValues = this.fieldValues;
        String content = this.template.getContent();

        for (TemplateFieldValue fieldValue : fieldValues) {
            String wildcard = fieldValue.getField().getWildcard();
            content = content.replace("%" + wildcard + "%", fieldValue.getValue());
        }

        return content;
    }
}
