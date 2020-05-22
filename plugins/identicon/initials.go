package identicon

import (
	"regexp"
	"strings"
)

var spec = regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>_]`)

// This controls how parsing of initials is handled.
type opts struct {
	allCaps   bool
	allowSpec bool
	limit     int
}

func GetInitials(s string, o opts) ([]rune, error) {
	nameOrInitials := strings.TrimSpace(s)

	if o.allCaps {
		nameOrInitials = strings.ToUpper(nameOrInitials)
	}
	if o.allowSpec {
		nameOrInitials = spec.ReplaceAllString(nameOrInitials, "")
	}

	names := strings.Split(nameOrInitials, " ")

	initials := []rune(nameOrInitials)
	assignedNames := 0

	if len(names) > 1 {
		initials = []rune("")
		start := 0

		for i := 0; i < o.limit; i++ {
			index := i

			if ((index == o.limit-1) && index > 0) || (index > (len(names) - 1)) {
				index = len(names) - 1
			}

			if assignedNames >= len(names) {
				start += 1
			}

			initials = append(initials, []rune(names[index])[start])
			assignedNames += 1
		}
	}

	if len(initials) > o.limit {
		return initials[0:o.limit], nil
	}

	return initials, nil
}
