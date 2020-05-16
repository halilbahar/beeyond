package at.htl.beeyond.model;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Entity;
import javax.persistence.OneToMany;
import java.util.List;

@Entity(name = "_user")
public class User extends PanacheEntity {
    private String name;
    @OneToMany
    private List<Namespace> namespaces;
    @OneToMany
    private List<CustomApplication> customApplications;
    @OneToMany
    private List<TemplateApplication> templateApplications;

    public User(String name) {
        this.name = name;
    }

    public User() {
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public List<Namespace> getNamespaces() {
        return namespaces;
    }

    public void setNamespaces(List<Namespace> namespaces) {
        this.namespaces = namespaces;
    }

    public List<CustomApplication> getCustomApplications() {
        return customApplications;
    }

    public void setCustomApplications(List<CustomApplication> customApplications) {
        this.customApplications = customApplications;
    }

    public List<TemplateApplication> getTemplateApplications() {
        return templateApplications;
    }

    public void setTemplateApplications(List<TemplateApplication> templateApplications) {
        this.templateApplications = templateApplications;
    }
}
