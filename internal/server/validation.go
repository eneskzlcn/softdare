package server

func validateServerAddress(s string) error {
	if len(s) != 5 {
		return ErrServerAddressLength
	}
	if s[0] != ':' {
		return ErrServerAddressSyntax
	}
	if !isDigit(s[1]) || !isDigit(s[2]) || !isDigit(s[3]) || !isDigit(s[4]) {
		return ErrServerAddressSyntax
	}
	return nil
}
func validateSessionKey(s string) error {
	if len(s) != 32 {
		return ErrSessionKeyStringLength
	}
	return nil
}
func isDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}
