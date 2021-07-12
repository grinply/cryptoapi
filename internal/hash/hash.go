package hash

import "crypto/sha1"

// Generate a sha1 hash string from the provided string information
// usefull to generate fingerprint of data
func SHA1(content string) string {
	var h = sha1.New()
	h.Write([]byte(content))
	return string(h.Sum(nil))
}
