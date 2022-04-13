package at.htl.beeyond.dto

import at.htl.beeyond.entity.ApplicationStatus
import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.validation.Exists
import java.time.LocalDate
import java.time.LocalDateTime
import java.util.*
import javax.json.bind.annotation.JsonbDateFormat
import javax.json.bind.annotation.JsonbProperty
import javax.json.bind.annotation.JsonbTransient
import javax.validation.constraints.NotNull
import javax.validation.constraints.Size

abstract class ApplicationDto(
    var id: Long? = null,
    @field:Size(max = 255) var note: String? = null,
    @set:JsonbTransient var status: ApplicationStatus? = null,
    @set:JsonbTransient var owner: UserDto? = null,
    @set:JsonbTransient var createdAt: LocalDateTime? = null,
    @set:JsonbTransient var startedAt: LocalDateTime? = null,
    @set:JsonbTransient var finishedAt: LocalDateTime? = null,
    @field:NotNull
    @field:Exists(entity = Namespace::class, fieldName = "namespace")
    var namespace: String? = null,
    @field:JsonbProperty("class")
    @field:NotNull var schoolClass: String? = null,
    @set:JsonbDateFormat(value = "dd.MM.yyyy")
    @field:JsonbProperty("to")
    @field:NotNull var toDate: LocalDate? = null,
    @field:Size(max = 255)
    @field:NotNull var purpose: String? = null
){
    abstract var content: String?
}
