package pat

import (
	"fmt"
	"testing"
)

func tGen(nn, kk int, out []string, t *testing.T) {
	f := GeneratePatterns
	ans := f(nn, kk)
	fmt.Println(len(ans))
	if fmt.Sprintf("%v", ans) != fmt.Sprintf("%v", out) {
		t.Errorf("GeneratePatterns(%v, %v) = %v, wanted %v", nn, kk, ans, out)
	}
}

func TestGeneratePatterns0(t *testing.T) {
	pat := []string{"10", "01"}
	tGen(2, 1, pat, t)
}
func TestGeneratePatterns1(t *testing.T) {
	pat := []string{"00"}
	tGen(2, 0, pat, t)
}
func TestGeneratePatterns2(t *testing.T) {
	pat := []string{"1100", "1010", "1001", "0110", "0101", "0011"}
	tGen(4, 2, pat, t)
}
