package at.htl.beeyond.entity;

import at.htl.beeyond.dto.TemplateDto;
import io.quarkus.hibernate.orm.panache.PanacheEntityBase;

import javax.persistence.*;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

@Entity
public class Template extends PanacheEntityBase {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;

    private String description;

    @Lob
    private String content;

    @ManyToOne
    private User owner;

    @OneToMany(mappedBy = "template", cascade = CascadeType.ALL)
    private List<TemplateField> fields = new LinkedList<>();

    private Boolean deleted;

    public Template(TemplateDto templateDto, User owner) {
        this.name = templateDto.getName();
        this.description = templateDto.getDescription();
        this.content = templateDto.getContent();
        this.owner = owner;
        this.deleted = false;

        this.fields = templateDto.getFields()
                .stream()
                .map(templateFieldDto -> new TemplateField(
                        templateFieldDto.getLabel(),
                        templateFieldDto.getWildcard(),
                        templateFieldDto.getDescription(),
                        this)
                )
                .collect(Collectors.toList());
    }

    public Template() {
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

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public User getOwner() {
        return owner;
    }

    public void setOwner(User owner) {
        this.owner = owner;
    }

    public List<TemplateField> getFields() {
        return fields;
    }

    public void setFields(List<TemplateField> fields) {
        this.fields = fields;
    }

    public Boolean getDeleted() {
        return deleted;
    }

    public void setDeleted(Boolean deleted) {
        this.deleted = deleted;
    }
}
