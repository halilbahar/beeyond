package at.htl.beeyond.model;

import javax.persistence.*;

@Entity
@NamedQueries({
        @NamedQuery(name = "Application.getAll", query = "select a from Application a")
})
public class Application {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private int replica;

    public Application() {
    }

    public Application(int replica) {
        this.replica = replica;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public int getReplica() {
        return replica;
    }

    public void setReplica(int replica) {
        this.replica = replica;
    }
}
