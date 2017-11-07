package db

// DB acts a fake demo database object
type DB struct{}

// New returns an empty instance od DB
func New(_ string) (*DB, error) {
	return &DB{}, nil
}

// Read does nothing
func (d *DB) Read(_ string) (string, error) {
	return "root", nil
}
