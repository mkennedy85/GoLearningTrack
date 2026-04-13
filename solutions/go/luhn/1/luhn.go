package luhn

import "unicode"

func Valid(id string) bool {
	sum := 0
    double := false
    digits := 0

    for i := len(id) - 1; i >= 0; i-- {
        r := rune(id[i])

        if r == ' ' {
            continue
        }

        if !unicode.IsDigit(r) {
			return false
		}

        d := int(r - '0')
        digits++

        if double {
            d *= 2
            if d > 9 {
                d -= 9
            }
        }

        sum += d
        double = !double
    }

    if digits <= 1 {
        return false
    }

    return sum%10 == 0
}
