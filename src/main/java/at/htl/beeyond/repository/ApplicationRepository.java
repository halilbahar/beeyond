package at.htl.beeyond.repository;

import at.htl.beeyond.model.Application;

import javax.enterprise.context.ApplicationScoped;
import javax.persistence.EntityManager;
import javax.persistence.PersistenceContext;
import javax.transaction.Transactional;
import java.util.List;

@Transactional
@ApplicationScoped
public class ApplicationRepository {

    @PersistenceContext
    EntityManager em;

    public List<Application> getAllApplications() {
        return em.createNamedQuery("Application.getAll").getResultList();
    }

    public Application getApplicationById(Long id) {
        return em.find(Application.class, id);
    }

    public Application uploadApplication(Application application) {
        em.persist(application);
        return application;
    }
}
