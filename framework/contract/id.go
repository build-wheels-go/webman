package contract

const IDKey = "wm:id"

type ID interface {
	NewID() string
}
