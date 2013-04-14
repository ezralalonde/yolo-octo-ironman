package pat

import (
	"fmt"
	"strings"
)

func GeneratePatterns(nn, kk uint) (slice []string) {
	format := fmt.Sprintf("%%0%db", nn)
	for ii := (1 << nn) - 1; ii >= 0; ii-- {
		str := fmt.Sprintf(format, ii)
		if ValidPattern(str, kk) {
			slice = append(slice, str)
		}
	}
	return
}

func ValidPattern(str string, kk uint) bool {
	return strings.Count(str, "1") == int(kk)
}
