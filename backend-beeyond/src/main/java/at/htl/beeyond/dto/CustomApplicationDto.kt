package at.htl.beeyond.dto

import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.ValidKubernetes
import java.time.LocalDateTime
import javax.validation.GroupSequence
import javax.validation.constraints.NotBlank

@GroupSequence(value = [CustomApplicationDto::class, Checks.KubernetesContent::class])
@ValidKubernetes(groups = [Checks.KubernetesContent::class])
class CustomApplicationDto(
        id: Long? = null,
        note: String? = null,
        status: ApplicationStatus? = null,
        owner: UserDto? = null,
        createdAt: LocalDateTime? = null,
        content: String? = null
) : ApplicationDto(
        id,
        note,
        status,
        owner,
        createdAt
) {
    constructor(customApplication: CustomApplication) : this(
            customApplication.id,
            customApplication.note,
            customApplication.status,
            UserDto(customApplication.owner),
            customApplication.createdAt,
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
