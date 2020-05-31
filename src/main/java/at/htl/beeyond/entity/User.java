package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntityBase;

import javax.persistence.*;
import java.util.List;

@Entity(name = "_user")
public class User extends PanacheEntityBase {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;

    @OneToMany(mappedBy = "user")
    private List<Namespace> namespaces;

    @OneToMany(mappedBy = "user")
    private List<CustomApplication> customApplications;

    @OneToMany(mappedBy = "user")
    private List<TemplateApplication> templateApplications;

    public User(String name) {
        this.name = name;
    }

    public User() {
    }

    public Long getId() {
        return id;
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
