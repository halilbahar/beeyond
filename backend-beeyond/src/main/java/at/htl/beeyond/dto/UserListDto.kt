package at.htl.beeyond.dto

import at.htl.beeyond.entity.User
import at.htl.beeyond.validation.Checks
import at.htl.beeyond.validation.Exists
import javax.validation.GroupSequence

@GroupSequence(value = [Checks.UserListExists::class, UserListDto::class])
data class UserListDto(
    @field:Exists(
        entity = User::class,
        fieldName = "name",
        groups = [Checks.UserListExists::class]
    ) var users: List<String>?
){
    constructor():this(null)
}
