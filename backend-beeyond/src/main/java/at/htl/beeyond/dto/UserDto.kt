package at.htl.beeyond.dto

import at.htl.beeyond.dto.UserDto
import at.htl.beeyond.entity.User

class UserDto {
    var id: Long? = null
    var name: String? = null

    constructor(id: Long?, name: String?) {
        this.id = id
        this.name = name
    }

    constructor() {}

    fun map(): User {
        return User.findById(id)
    }

    companion object {
        fun map(user: User): UserDto {
            return UserDto(user.id, user.name)
        }
    }
}
