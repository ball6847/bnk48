package flag

import "flag"

// Secret value from flag
var Secret *string

// Init app flag
func Init() {
	Secret = flag.String("secret", "mysecret", "-secret=yoursecret")
}

// Parse just Parse
func Parse() {
	flag.Parse()
}
