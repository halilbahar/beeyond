package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntity;
import org.hibernate.validator.constraints.Length;

import javax.json.bind.annotation.JsonbTransient;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.Lob;
import javax.validation.constraints.NotBlank;

@Entity
public class CustomApplication extends PanacheEntity {
    @NotBlank(message = "The content cannot be empty")
    @Lob
    private String content;
    @Length(max = 255, message = "The note length must be between 0 and 255")
    private String note;
    @Enumerated(EnumType.STRING)
    private ApplicationStatus status = ApplicationStatus.PENDING;

    public CustomApplication(String content, String note, ApplicationStatus status) {
        this.content = content;
        this.note = note;
        this.status = status;
    }

    public CustomApplication() {
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }

    public ApplicationStatus getStatus() {
        return status;
    }

    @JsonbTransient
    public void setStatus(ApplicationStatus status) {
        this.status = status;
    }
}
