package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Entity;
import javax.persistence.Lob;

@Entity
public class Template extends PanacheEntity {

    private String name;
    private String description;
    @Lob
    private String content;

    public Template(String name, String content, String description) {
        this.name = name;
        this.content = content;
        this.description = description;
    }

    public Template() {
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
}
