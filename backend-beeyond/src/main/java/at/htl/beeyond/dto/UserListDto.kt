package at.htl.beeyond.dto

import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Exists
import at.htl.beeyond.validation.NamespaceValid
import org.hibernate.validator.constraints.Length
import java.util.*
import javax.validation.GroupSequence
import javax.validation.constraints.NotNull
import javax.validation.constraints.Pattern

@GroupSequence(UserListDto::class)
data class UserListDto(
        @field:Length(min = 1, max = 253)
        @field:Pattern(regexp = "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$")
        @field:NamespaceValid
        var namespace: String = "",

        @field:NotNull
        @field:Exists(entity = User::class, fieldName = "name")
        var users: List<String> = LinkedList()
)
