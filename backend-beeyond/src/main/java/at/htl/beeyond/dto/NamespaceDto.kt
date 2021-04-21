package at.htl.beeyond.dto

import at.htl.beeyond.entity.Namespace
import java.util.*
import javax.json.bind.annotation.JsonbTransient

class NamespaceDto(
        @set:JsonbTransient var id: Long? = null,
        var namespace: String? = null,
        var users: List<UserDto?> = LinkedList()
) {
    constructor(namespace: Namespace) : this(
            namespace.id,
            namespace.namespace,
            namespace.users.map { UserDto(it) }
    )
}
