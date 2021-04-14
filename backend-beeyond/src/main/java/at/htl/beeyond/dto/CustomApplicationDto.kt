package at.htl.beeyond.dto

import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.User
import java.time.LocalDateTime
import javax.validation.constraints.NotBlank

class CustomApplicationDto(
        id: Long? = null,
        note: String? = null,
        status: ApplicationStatus? = null,
        owner: UserDto? = null,
        createdAt: LocalDateTime? = null,
        startedAt: LocalDateTime? = null,
        finishedAt: LocalDateTime? = null,
        content: String? = null
) : ApplicationDto(
        id,
        note,
        status,
        owner,
        createdAt,
        startedAt,
        finishedAt
) {
    constructor(customApplication: CustomApplication) : this(
            customApplication.id,
            customApplication.note,
            customApplication.status,
            UserDto(customApplication.owner),
            customApplication.createdAt,
            customApplication.startedAt,
            customApplication.finishedAt,
            customApplication.content
    )

    @field:NotBlank
    var content: String? = content
        set(value) {
            if (value != null) {
                field = value.trim()
            }
        }
}
