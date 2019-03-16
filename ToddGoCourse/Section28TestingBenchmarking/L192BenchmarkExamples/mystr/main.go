// Package mystr pulls together words into sentences
package mystr

import "strings"

//Cat uses a resource intensive method to pull strings together
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

//Join uses an out the box more efficienct method to join words together
func Join(xs []string) string {
	return strings.Join(xs, " ")

}
