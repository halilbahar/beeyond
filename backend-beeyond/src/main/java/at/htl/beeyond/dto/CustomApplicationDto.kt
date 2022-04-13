package at.htl.beeyond.dto

import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.entity.CustomApplication
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.ValidKubernetes
import java.time.LocalDate
import java.time.LocalDateTime
import javax.validation.GroupSequence
import javax.validation.constraints.NotBlank

@GroupSequence(CustomApplicationDto::class, Checks.KubernetesContent::class)
@ValidKubernetes(groups = [Checks.KubernetesContent::class])
class CustomApplicationDto(
    id: Long? = null,
    note: String? = null,
    status: ApplicationStatus? = null,
    owner: UserDto? = null,
    createdAt: LocalDateTime? = null,
    startedAt: LocalDateTime? = null,
    finishedAt: LocalDateTime? = null,
    content: String? = null,
    namespace: String? = null,
    schoolClass: String? = null,
    toDate: LocalDate? = null,
    purpose: String? = null
) : ApplicationDto(
    id,
    note,
    status,
    owner,
    createdAt,
    startedAt,
    finishedAt,
    namespace,
    schoolClass,
    toDate,
    purpose
) {
    constructor(customApplication: CustomApplication) : this(
        customApplication.id,
        customApplication.note,
        customApplication.status,
        UserDto(customApplication.owner),
        customApplication.createdAt,
        customApplication.startedAt,
        customApplication.finishedAt,
        customApplication.content,
        customApplication.namespace.namespace,
        customApplication.schoolClass,
        customApplication.toDate,
        customApplication.purpose
    )

    @field:NotBlank
    override var content: String? = content
        set(value) {
            if (value != null) {
                field = value.trim()
            }
        }
}
