package at.htl.beeyond.dto

import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.dto.UserDto
import org.hibernate.validator.constraints.Length
import java.time.LocalDateTime
import javax.json.bind.annotation.JsonbTransient
import javax.validation.constraints.Size

abstract class ApplicationDto {
    open var id: Long? = null
        protected set

    @field:Size(max = 255)
    var note: String? = null

    @set:JsonbTransient
    var status: ApplicationStatus? = null

    @set:JsonbTransient
    var owner: UserDto? = null

    @set:JsonbTransient
    var createdAt: LocalDateTime? = null

    constructor(id: Long?, note: String?, status: ApplicationStatus?, owner: UserDto?, createdAt: LocalDateTime?) {
        this.id = id
        this.note = note
        this.status = status
        this.owner = owner
        this.createdAt = createdAt
    }

    constructor() {}
}
