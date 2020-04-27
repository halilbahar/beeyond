package at.htl.beeyond.service;

import javax.enterprise.context.ApplicationScoped;
import java.io.File;
import java.net.URISyntaxException;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.stream.Collectors;

@ApplicationScoped
public class TemplateService {

    public List<String> getAllTemplates() {
        try {
            File directory = new File(this.getClass().getResource("/templates").toURI());
            return Arrays.asList(directory.list()).stream().map(s -> s.split("-")[0]).collect(Collectors.toList());
        } catch (URISyntaxException e) {
            e.printStackTrace();
        }
        return new LinkedList<>();
    }
}
