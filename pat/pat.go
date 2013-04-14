package pat

import (
    "strings"
)

func GeneratePatterns(nn, kk int) []string {
    aa := make([]string, nn)
    var bb []string
    Gen(nn, kk, aa, &bb)
    return bb
}

func Gen(nn, kk int, aa []string, bb *[]string) {
    if nn == 0 && kk == 0 {
        *bb = append(*bb, strings.Join(aa, ""))
    } else {
        if kk > 0 {
            aa[len(aa) - nn] = "1"
            Gen(nn-1, kk - 1, aa, bb)
        }
        if nn > kk {
            aa[len(aa) - nn] = "0"
            Gen(nn-1, kk , aa, bb)
        }
    }
}

