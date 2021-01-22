package at.htl.beeyond.dto

import at.htl.beeyond.entity.User

data class UserDto(
        var id: Long?=null,
        var name: String?=null
) {
    constructor(user: User) : this(user.id, user.name)

    fun map(): User {
        return User.findById(id)
    }
}
