package at.htl.beeyond.model;

import com.sun.org.apache.xpath.internal.operations.Bool;

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
    private Boolean isApproved;

    public Application() {
    }

    public Application(int replica) {
        this(replica, false);
    }

    public Application(int replica, boolean isApproved) {
        this.replica = replica;
        this.isApproved = isApproved;
    }


    public Long getId() {
        return id;
    }

    /*
    public void setId(Long id) {

        this.id = id;
    } */

    public int getReplica() {
        return replica;
    }

    public void setReplica(int replica) {
        this.replica = replica;
    }

    public Boolean getApproved() {
        return isApproved;
    }

    public void setApproved(Boolean approved) {
        isApproved = approved;
    }
}
