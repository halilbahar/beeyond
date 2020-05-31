package at.htl.beeyond.entity;

import javax.persistence.Entity;
import javax.persistence.Lob;
import javax.persistence.ManyToOne;

@Entity
public class CustomApplication extends Application {

    @Lob
    private String content;

    @ManyToOne
    private User user;

    public CustomApplication(String note, String content, User user) {
        super(note);
        this.content = content;
        this.user = user;
    }

    public CustomApplication() {
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }
}
