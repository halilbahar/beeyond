package at.htl.beeyond.dto;

import at.htl.beeyond.entity.CustomApplication;
import at.htl.beeyond.entity.User;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotBlank;

public class CustomApplicationDto {

    private Long id;

    @Length(max = 255, message = "The note of the custom application cannot be longer than 255 characters")
    private String note;

    @NotBlank(message = "The content of the custom application cannot be empty")
    private String content;

    public CustomApplicationDto(Long id, String note, String content) {
        this.id = id;
        this.note = note;
        this.content = content;
    }

    public CustomApplicationDto() {
    }

    public Long getId() {
        return id;
    }

    public String getNote() {
        return note;
    }

    public void setNote(String note) {
        this.note = note;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public CustomApplication map(User owner) {
        return new CustomApplication(note, owner, content);
    }
}
