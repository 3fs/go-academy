// Package bucket demonstrates an interface for a fictional bucket storage
// where data can be sequentially appended in multiple buckets.
package bucket

type (
	// Reader supports methods required to read data
	Reader interface {
		Get(int) (string, error)
		GetAll() ([]string, error)
	}

	// Writer supports methods required to modify data
	Writer interface {
		Append(...string) error
		Remove(int) error
	}

	// ReadWriter joins the Reader and Writer interfaces
	ReadWriter interface {
		Reader
		Writer
	}

	// Storage describes methods required to handle a set of buckets
	Storage interface {
		Create(string) (ReadWriter, error)
		Remove(string) error
	}
)
