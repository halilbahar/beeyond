package at.htl.beeyond.entity;

import javax.persistence.Entity;
import javax.persistence.ManyToOne;

@Entity
public class TemplateApplication extends Application {

    @ManyToOne
    private Template template;

    @ManyToOne
    private User user;

    public TemplateApplication(String note, Template template, User user) {
        super(note);
        this.template = template;
        this.user = user;
    }

    public TemplateApplication() {
    }

    public Template getTemplate() {
        return template;
    }

    public void setTemplate(Template template) {
        this.template = template;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
