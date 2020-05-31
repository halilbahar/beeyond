package at.htl.beeyond.dto;

import at.htl.beeyond.entity.Template;
import at.htl.beeyond.entity.User;
import org.hibernate.validator.constraints.Length;

import javax.validation.constraints.NotBlank;

public class TemplateDto {

    @NotBlank(message = "The name of the template cannot be blank")
    @Length(max = 255, message = "The name cannot be longer than 255 characters")
    private String name;

    @Length(max = 255, message = "The description cannot be longer than 255 characters")
    private String description;

    @NotBlank(message = "The content of the template cannot be blank")
    private String content;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Template map(User owner) {
        return new Template(name, description, content, owner);
    }
}
