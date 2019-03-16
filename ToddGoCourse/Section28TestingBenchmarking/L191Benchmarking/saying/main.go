// Package saying contains salutations
package saying

import "fmt"

// Greet incorproates th eusers names and says welcome
func Greet(s string) string {
	return fmt.Sprint("Welcome to the sayings package,", s)
}
