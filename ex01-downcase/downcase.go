package downcase

func Downcase(s string) (string, error) {
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i]

		if tmp > 64 && tmp < 91 {
			tmp = s[i] + 32
		}

		str[i] = tmp
	}

	return string(str), nil
}
