package validatiors

import "regexp"

func IsValidAddress(addr string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9.-]+:[0-9]{1,5}$`)
	return re.MatchString(addr)
}
