//the files for testing and implementing two-fer for exercism's go path
package twofer

// The core function of the two-fer package.  Should respond "One for X, one for me."
func ShareWith(name string) (product string) {
	product = "One for "
	if name != "" {
		product += name
	} else {
		product += "you"
	}
	product += ", one for me."
	return product
}
