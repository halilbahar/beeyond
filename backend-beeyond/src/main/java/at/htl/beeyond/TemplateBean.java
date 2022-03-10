package at.htl.beeyond;

import at.htl.beeyond.entity.Template;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.quarkus.runtime.StartupEvent;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.event.Observes;
import javax.json.JsonReader;
import javax.transaction.Transactional;
import java.io.IOException;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Paths;

@ApplicationScoped
class TemplateBean {
    @Transactional
    void init(@Observes StartupEvent event) throws URISyntaxException, IOException {
        var uri = getClass().getResource("/templates/json").toURI();
        var dirPath = Paths.get(uri);
        Files.list(dirPath).forEach(fileName -> {
            try {
                var jsonString = Files.readString(Paths.get(fileName.toString()));
                var newTemplate = new ObjectMapper().readValue(jsonString, Template.class);
                newTemplate.setContent(
                        Files.readString(Paths.get(getClass().getResource("/templates/yml/"+newTemplate.getContent()).toURI()))
                );
                newTemplate.getFields().forEach(fieldName ->{
                    fieldName.setTemplate(newTemplate);
                });
                newTemplate.persist();
            } catch (Exception e) {
                e.printStackTrace();
            }
        });
    }
}