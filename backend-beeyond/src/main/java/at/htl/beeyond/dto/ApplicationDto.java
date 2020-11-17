package at.htl.beeyond.dto;

import at.htl.beeyond.entity.ApplicationStatus;
import org.hibernate.validator.constraints.Length;

import javax.json.bind.annotation.JsonbTransient;

public abstract class ApplicationDto {

    protected Long id;

    @Length(max = 255)
    protected String note;

    protected ApplicationStatus status;

    protected UserDto owner;

    public ApplicationDto(Long id, String note, ApplicationStatus status, UserDto owner) {
        this.id = id;
        this.note = note;
        this.status = status;
        this.owner = owner;
    }

    public ApplicationDto() {
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

    public ApplicationStatus getStatus() {
        return status;
    }

    @JsonbTransient
    public void setStatus(ApplicationStatus status) {
        this.status = status;
    }

    public UserDto getOwner() {
        return owner;
    }

    @JsonbTransient
    public void setOwner(UserDto owner) {
        this.owner = owner;
    }

}
