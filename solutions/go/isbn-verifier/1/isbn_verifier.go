package isbnverifier

func IsValidISBN(isbn string) bool {
	weight := 10
	sum := 0

	for _, r := range isbn {
		if r == '-' {
			continue
		}

		if weight == 0 {
			return false
		}

		switch {
		case r >= '0' && r <= '9':
			sum += int(r-'0') * weight
		case r == 'X' && weight == 1:
			sum += 10
		default:
			return false
		}

		weight--
	}

	return weight == 0 && sum%11 == 0
}