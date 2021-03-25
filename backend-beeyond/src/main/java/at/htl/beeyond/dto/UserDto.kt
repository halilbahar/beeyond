package at.htl.beeyond.dto

import at.htl.beeyond.entity.User

data class UserDto(
    var id: Long? = null,
    var name: String? = null
) {
    constructor(user: User) : this(user.id, user.name)
    constructor() : this(null, null)
}
