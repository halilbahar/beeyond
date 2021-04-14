package at.htl.beeyond.dto

import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Exists
import java.util.*
import javax.validation.GroupSequence
import javax.validation.constraints.NotNull
import javax.validation.constraints.Size

@GroupSequence(value = [UserListDto::class])
data class UserListDto(
        @NotNull
        @Size(min = 1)
        @field:Exists(
                entity = User::class,
                fieldName = "name"
        ) var users: List<String> = LinkedList()
)
