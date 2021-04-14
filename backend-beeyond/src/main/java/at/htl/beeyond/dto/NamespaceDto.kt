package at.htl.beeyond.dto

import at.htl.beeyond.entity.Namespace
import at.htl.beeyond.entity.User
import org.hibernate.validator.constraints.Length
import java.util.*
import javax.json.bind.annotation.JsonbTransient
import javax.validation.constraints.NotNull

class NamespaceDto(
    @set:JsonbTransient var id: Long? = null,
    @Length(min = 1, max = 50) var namespace: String? = null,
    @NotNull var users: List<UserDto?> = LinkedList()
) {
    constructor(namespace: Namespace) : this(
        namespace.id,
        namespace.namespace,
        namespace.users.map { UserDto(it) }
    )
}
