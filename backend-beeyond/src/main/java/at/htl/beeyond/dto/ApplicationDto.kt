package at.htl.beeyond.dto

import at.htl.beeyond.entity.ApplicationStatus
import java.time.LocalDateTime
import javax.json.bind.annotation.JsonbTransient
import javax.validation.constraints.Size

abstract class ApplicationDto(
        var id: Long? = null,
        @field:Size(max = 255) var note: String? = null,
        @set:JsonbTransient var status: ApplicationStatus? = null,
        @set:JsonbTransient var owner: UserDto? = null,
        @set:JsonbTransient var createdAt: LocalDateTime? = null
)
