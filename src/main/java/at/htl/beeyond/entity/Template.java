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

    @OneToMany(mappedBy = "template", cascade = CascadeType.PERSIST)
    private List<TemplateField> fields = new LinkedList<>();

    public Template(String name, String description, String content, User owner) {
        this.name = name;
        this.description = description;
        this.content = content;
        this.owner = owner;
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

    public static List<TemplateDto> getDtos() {
        return Template.findAll().stream().map(o -> {
            Template template = (Template) o;
            return TemplateDto.map(template);
        }).collect(Collectors.toList());
    }
}
