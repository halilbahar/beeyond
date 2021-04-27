package at.htl.beeyond.entity;

import at.htl.beeyond.dto.CustomApplicationDto;

import javax.naming.Name;
import javax.persistence.Entity;
import javax.persistence.Lob;

@Entity
public class CustomApplication extends Application {

    @Lob
    private String content;

    public CustomApplication(CustomApplicationDto customApplicationDto, User owner) {
        super(customApplicationDto.getNote(), owner, new Namespace(customApplicationDto.getNamespace()));
        this.content = customApplicationDto.getContent();
    }

    public CustomApplication() {
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }
}
