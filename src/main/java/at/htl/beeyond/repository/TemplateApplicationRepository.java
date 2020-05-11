package at.htl.beeyond.repository;

import at.htl.beeyond.model.TemplateApplication;
import io.quarkus.hibernate.orm.panache.PanacheRepository;

import javax.enterprise.context.ApplicationScoped;
import javax.transaction.Transactional;

@Transactional
@ApplicationScoped
public class TemplateApplicationRepository implements PanacheRepository<TemplateApplication> {


    public boolean persistApplication(TemplateApplication templateApplication) {
        if (this.isPersistent(templateApplication)) {
            return false;
        }
        this.persist(templateApplication);
        return true;
    }

    public void setApproval(Long id, boolean isApproved) {
        this.findById(id).setIsApproved(isApproved);
    }
}
