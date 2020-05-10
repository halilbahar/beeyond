package at.htl.beeyond.model;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;

@Entity
public class TemplateApplication {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private Integer replica;
    private Boolean isApproved;

    public TemplateApplication(Integer replica, boolean isApproved) {
        this.replica = replica;
        this.isApproved = isApproved;
    }

    public TemplateApplication(Integer replica) {
        this(replica, false);
    }

    public TemplateApplication() {
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

    public Boolean getIsApproved() {
        return isApproved;
    }

    public void setIsApproved(Boolean approved) {
        isApproved = approved;
    }
}
