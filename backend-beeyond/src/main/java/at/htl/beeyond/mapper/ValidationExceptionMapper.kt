package at.htl.beeyond.mapper

import at.htl.beeyond.dto.FailedFieldDto
import java.util.stream.StreamSupport
import javax.validation.ConstraintViolationException
import javax.validation.Path
import javax.ws.rs.core.Response
import javax.ws.rs.ext.ExceptionMapper
import javax.ws.rs.ext.Provider

@Provider
class ValidationExceptionMapper : ExceptionMapper<ConstraintViolationException> {
    override fun toResponse(exception: ConstraintViolationException): Response {
        // TODO: find a way to display the whole json / object path
        return Response.status(422).entity(
                exception.constraintViolations.map {
                    FailedFieldDto(
                            StreamSupport.stream(it.propertyPath.spliterator(), false)
                                    .reduce { _: Path.Node?, second: Path.Node? -> second }
                                    .orElse(null).toString(),
                            if (it.invalidValue != null) it.invalidValue.toString() else "",
                            it.message
                    )
                }.toList()
        ).build()
    }
}
