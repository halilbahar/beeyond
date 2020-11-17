package at.htl.beeyond.dto;

import at.htl.beeyond.entity.ApplicationStatus;
import at.htl.beeyond.entity.CustomApplication;
import at.htl.beeyond.entity.User;
import at.htl.beeyond.validation.Checks;
import javax.validation.GroupSequence;
import javax.validation.constraints.NotBlank;

@GroupSequence({CustomApplicationDto.class, Checks.TemplateContent.class})
public class CustomApplicationDto extends ApplicationDto {

    @NotBlank
    private String content;

    public CustomApplicationDto(Long id, String note, ApplicationStatus status, UserDto owner, String content) {
        super(id, note, status, owner);
        this.content = content;
    }

    public CustomApplicationDto() {
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
                customApplication.getStatus(),
                UserDto.map(customApplication.getOwner()),
                customApplication.getContent()
        );
    }
}
