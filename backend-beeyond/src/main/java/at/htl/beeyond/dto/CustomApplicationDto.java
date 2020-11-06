package at.htl.beeyond.dto;

import at.htl.beeyond.entity.CustomApplication;
import at.htl.beeyond.entity.User;
import at.htl.beeyond.validation.Checks;
import org.hibernate.validator.constraints.Length;

import javax.validation.GroupSequence;
import javax.validation.constraints.NotBlank;

@GroupSequence({CustomApplicationDto.class, Checks.TemplateContent.class})
public class CustomApplicationDto {

    private Long id;

    @Length(max = 255)
    private String note;

    @NotBlank
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
        this.note = note.trim();
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content.trim();
    }

    public CustomApplication map(User owner) {
        return new CustomApplication(note, owner, content);
    }

    public static CustomApplicationDto map(CustomApplication customApplication) {
        return new CustomApplicationDto(
                customApplication.getId(),
                customApplication.getNote(),
                customApplication.getContent()
        );
    }
}
