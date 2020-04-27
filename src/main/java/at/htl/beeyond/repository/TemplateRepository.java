package at.htl.beeyond.repository;

import at.htl.beeyond.model.Template;
import io.quarkus.hibernate.orm.panache.PanacheRepository;

import javax.enterprise.context.ApplicationScoped;
import javax.transaction.Transactional;

@Transactional
@ApplicationScoped
public class TemplateRepository implements PanacheRepository<Template> {
}
