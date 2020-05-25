package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Entity;
import javax.persistence.ManyToOne;

@Entity
public class Namespace extends PanacheEntity {

    private String name;

    @ManyToOne
    private User user;

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

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
