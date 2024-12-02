package validatiors

import "regexp"

// IsValidAddress проверяет валидность адреса.
// Если withScheme == true, учитывается схема (http:// или https://).
func IsValidAddress(addr string, withScheme bool) bool {
	var re *regexp.Regexp
	if withScheme {
		re = regexp.MustCompile(`^(http://|https://)[a-zA-Z0-9.-]+:[0-9]{1,5}$`)
	} else {
		re = regexp.MustCompile(`^[a-zA-Z0-9.-]+:[0-9]{1,5}$`)
	}
	return re.MatchString(addr)
}
