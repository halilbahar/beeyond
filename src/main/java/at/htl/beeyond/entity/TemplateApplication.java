package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Entity;
import javax.persistence.ManyToOne;

@Entity
public class TemplateApplication extends PanacheEntity {
    @ManyToOne
    Template template;

    public TemplateApplication() {
    }

    public Template getTemplate() {
        return template;
    }

    public void setTemplate(Template template) {
        this.template = template;
    }
}
