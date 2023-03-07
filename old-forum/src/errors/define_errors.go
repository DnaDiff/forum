package errors

type Error struct {
	Code     int
	Message  string
	Original error
}

// Custom error messages for specific error cases
const (
	TEMPLATE_CORRUPTED_ERROR = "500 Internal Server Error: Template Corrupted"
)
