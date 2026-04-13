// Package twofer determines what you will say as you give away a two-fer one cookie.
package twofer

import "fmt"

// ShareWith returns what you say when sharing depending on whether the inidividual is known or not.
func ShareWith(name string) string {
	if len(name) == 0 {
        return "One for you, one for me."
    }
	return fmt.Sprintf("One for %s, one for me.", name)
}
