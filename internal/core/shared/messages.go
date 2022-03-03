package shared

var (
	INVALID_MESSAGE_ERROR   = "The message format read from the given topic is invalid"
	VALIDATION_ERROR        = "The request has validation errors"
	REQUEST_NOT_FOUND       = "The requested resource was NOT found"
	DUPLICATE_REQUEST_ERROR = "A resource having the same identifier already exist"
	GENERIC_ERROR           = "Generic error occurred. See stacktrace for details"
	AUTHORISATION_ERROR     = "You do NOT have adequate permission to access this resource"
	NO_PRINCIPAL            = "Principal identifier NOT provided"
	NORESOURCEFOUND         = "this resource does not exist"
	NORECORDFOUND           = "sorry, no record found"
	NOERRORSFOUND           = "no errors at the moment"
)
