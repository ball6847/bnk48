package payload

import (
	"crypto/sha256"
	"fmt"
)

// Signup payload for signup request
type Signup struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// HashPassword hash password using sha256
func (p *Signup) HashPassword() string {
	sum := sha256.Sum256([]byte(p.Password))
	return fmt.Sprintf("%x", sum)
}
