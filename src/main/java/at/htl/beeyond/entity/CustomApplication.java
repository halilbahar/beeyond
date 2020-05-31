package at.htl.beeyond.entity;

import io.quarkus.hibernate.orm.panache.PanacheEntity;

import javax.persistence.*;

@Entity
public class CustomApplication extends PanacheEntity {
    @Lob
    private String content;
    private String note;
    @Enumerated(EnumType.STRING)
    private ApplicationStatus status = ApplicationStatus.PENDING;
    @ManyToOne
    private User user;

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

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
