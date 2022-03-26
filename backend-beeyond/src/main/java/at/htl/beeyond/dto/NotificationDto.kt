package at.htl.beeyond.dto

import at.htl.beeyond.entity.Notification
import at.htl.beeyond.entity.NotificationStatus
import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Exists
import javax.validation.constraints.NotNull
import javax.validation.constraints.Size

class NotificationDto(
        var id: Long? = null,
        @field:Size(max = 255) var message: String? = null,
        @field:NotNull var details: String? = null,
        @field:Exists(entity = User::class, fieldName = "id") var userId: Long? = null,
        var status: NotificationStatus? = null,
        var entityName: String? = null,
        var entityId: Long? = null
){
    constructor(notification: Notification) : this(
            notification.id,
            notification.message,
            notification.details,
            notification.user.id,
            notification.status,
            notification.entityName,
            notification.entityId
    )
}