package at.htl.beeyond.model;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Entity;

@Entity
public class Namespace extends PanacheEntity {
    private String name;

    public Namespace(String name) {
        this.name = name;
    }

    public Namespace() {
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}
