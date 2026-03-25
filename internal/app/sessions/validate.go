package sessions

func ValidateKey(key string) (isValid bool) {
	// *2 since hex encoded
	if KeyLen*2 != len(key) {
		return false
	}
	return true
}
