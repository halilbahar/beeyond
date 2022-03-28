package at.htl.beeyond;

import at.htl.beeyond.entity.Template;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.quarkus.runtime.StartupEvent;
import org.jboss.logging.Logger;

import javax.enterprise.context.ApplicationScoped;
import javax.enterprise.event.Observes;
import javax.inject.Inject;
import javax.transaction.Transactional;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Collectors;

@ApplicationScoped
class TemplateBean {
    private static final String[] FILENAMES = { "quarkus" };
    @Transactional
    void init(@Observes StartupEvent event) throws IOException {
        for(String filename : FILENAMES) {
            try (InputStream in = getClass()
                    .getResourceAsStream("/templates/json/" + filename + ".json")){
                assert in != null;

                BufferedReader reader = new BufferedReader(new InputStreamReader(in));
                var jsonString = reader.lines().collect(Collectors.joining("\n"));
                var newTemplate = new ObjectMapper().readValue(jsonString, Template.class);

                InputStream contentIn = getClass()
                        .getResourceAsStream("/templates/yml/" + newTemplate.getContent());
                assert contentIn != null;
                BufferedReader contentReader = new BufferedReader(new InputStreamReader(contentIn));
                newTemplate.setContent(
                        contentReader.lines().collect(Collectors.joining("\n"))
                );
                contentIn.close();
                newTemplate.getFields().forEach(fieldName ->{
                    fieldName.setTemplate(newTemplate);
                });
                newTemplate.persist();
            }
        }
    }
}