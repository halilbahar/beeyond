package at.htl.beeyond.entity;

import at.htl.beeyond.dto.CustomApplicationDto;

import javax.lang.model.element.Name;
import javax.persistence.Entity;
import javax.persistence.Lob;

@Entity
public class CustomApplication extends Application {

    public CustomApplication(CustomApplicationDto customApplicationDto, User owner) {
        super(
                customApplicationDto.getNote(),
                owner,
                Namespace.find("namespace", customApplicationDto.getNamespace()).firstResult(),
                customApplicationDto.getSchoolClass(),
                customApplicationDto.getToDate(),
                customApplicationDto.getPurpose(),
                customApplicationDto.getContent()
        );
    }

    public CustomApplication() {
    }
}
