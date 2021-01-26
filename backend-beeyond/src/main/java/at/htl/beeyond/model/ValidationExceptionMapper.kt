package at.htl.beeyond.model

import java.util.stream.Collectors
import java.util.stream.StreamSupport
import javax.validation.ConstraintViolation
import javax.validation.ConstraintViolationException
import javax.validation.Path
import javax.ws.rs.core.Response
import javax.ws.rs.ext.ExceptionMapper
import javax.ws.rs.ext.Provider

@Provider
class ValidationExceptionMapper : ExceptionMapper<ConstraintViolationException> {
    override fun toResponse(exception: ConstraintViolationException): Response {
        return Response.status(422).entity(
                exception.constraintViolations.stream()
                        .map { o: ConstraintViolation<*> ->
                            FailedField(
                                    StreamSupport.stream(o.propertyPath.spliterator(), false).reduce { first: Path.Node?, second: Path.Node? -> second }.orElse(null).toString(),
                                    if (o.invalidValue != null) o.invalidValue.toString() else "",
                                    o.message
                            )
                        }.collect(Collectors.toList())).build()
    }
}
