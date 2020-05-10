package at.htl.beeyond.model;

import javax.persistence.*;

@Entity
public class TemplateApplication {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private Integer replica;
    private Boolean isApproved;
    @ManyToOne
    private Template template;
    @Transient
    private String templateName;

    public TemplateApplication(String name, Integer replica, Boolean isApproved) {
        this.name = name;
        this.replica = replica;
        this.isApproved = isApproved;
    }

    public TemplateApplication() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getReplica() {
        return replica;
    }

    public void setReplica(int replica) {
        this.replica = replica;
    }

    public Boolean getIsApproved() {
        return isApproved;
    }

    public void setIsApproved(Boolean approved) {
        isApproved = approved;
    }

    public Template getTemplate() {
        return template;
    }

    public void setTemplate(Template template) {
        this.template = template;
    }

    public String getTemplateName() {
        return templateName;
    }

    public void setTemplateName(String templateName) {
        this.templateName = templateName;
    }
}
