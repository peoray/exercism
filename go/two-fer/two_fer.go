// Package twofer implements a way to satisfies two criteria or needs simultaneously
package twofer

// ShareWith returns a greeting with a name or a defaut value if a name does not exist
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}