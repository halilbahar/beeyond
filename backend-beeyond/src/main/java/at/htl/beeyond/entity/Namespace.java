package at.htl.beeyond.entity;


import at.htl.beeyond.dto.NamespaceDto;
import io.quarkus.hibernate.orm.panache.PanacheEntityBase;

import javax.persistence.*;
import java.util.List;
import java.util.stream.Collectors;

@Entity
public class Namespace extends PanacheEntityBase {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String namespace;

    @ManyToMany
    private List<User> users;

    private boolean deleted;

    private boolean isDefault;

    public Namespace(String namespace) {
        this.namespace = namespace;
    }

    public Namespace(NamespaceDto namespaceDto) {
        this.namespace = namespaceDto.getNamespace();
        this.users = namespaceDto.getUsers()
                .stream()
                .map(user -> (User) User.find("name", user.getName()).firstResult())
                .collect(Collectors.toList());
    }

    public Namespace() {
    }

    public Long getId() {
        return id;
    }

    public String getNamespace() {
        return namespace;
    }

    public void setNamespace(String namespace) {
        this.namespace = namespace;
    }

    public List<User> getUsers() {
        return users;
    }

    public void setUsers(List<User> users) {
        this.users = users;
    }

    public boolean isDeleted() {
        return deleted;
    }

    public void setDeleted(boolean deleted) {
        this.deleted = deleted;
    }

    public boolean isDefault() {
        return isDefault;
    }

    public void setDefault(boolean aDefault) {
        isDefault = aDefault;
    }
}
