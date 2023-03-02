package errors

type Error struct {
	Code     int
	Message  string
	Original error
}

const (
	INTERNAL_SERVER_ERROR    = "500 Internal Server Error"
	TEMPLATE_CORRUPTED_ERROR = "500 Internal Server Error: Template Corrupted"
	NOT_FOUND_ERROR          = "404 Not Found"
	BAD_REQUEST_ERROR        = "400 Bad Request"
	UNAUTHROIZED_ERROR       = "401 Unauthorized"
)
