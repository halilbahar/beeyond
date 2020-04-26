package at.htl.beeyond.repository;

import at.htl.beeyond.model.Application;
import io.quarkus.hibernate.orm.panache.PanacheRepository;

import javax.enterprise.context.ApplicationScoped;
import javax.transaction.Transactional;

@Transactional
@ApplicationScoped
public class ApplicationRepository implements PanacheRepository<Application> {


    public boolean persistApplication(Application application) {
        if (this.isPersistent(application)) {
            return false;
        }
        this.persist(application);
        return true;
    }

    public void setApproval(Long id, boolean isApproved) {
        this.findById(id).setIsApproved(isApproved);
    }
}
