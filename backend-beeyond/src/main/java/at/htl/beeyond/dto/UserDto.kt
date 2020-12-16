package at.htl.beeyond.dto

import at.htl.beeyond.entity.User

data class UserDto(var id: Long?, var name: String?) {

    fun map(): User {
        return User.findById(id)
    }

    companion object {
        @kotlin.jvm.JvmStatic
        fun map(user: User): UserDto {
            return UserDto(user.id, user.name)
        }
    }
}
