package at.htl.beeyond.model;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.Lob;

@Entity
public class CustomApplication extends PanacheEntity {
    @Lob
    private String content;
    private String note;
    @Enumerated(EnumType.STRING)
    private ApplicationStatus status;

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

    public void setStatus(ApplicationStatus status) {
        this.status = status;
    }
}
