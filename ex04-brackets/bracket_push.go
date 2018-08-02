package brackets

func Bracket(s string) (bool, error) {
	var mas [10]uint8
	pos := 0

	for i := 0; i < len(s); i++ {
		if pos == 0 {
			mas[pos] = s[i]
			pos++
		} else {
			if mas[pos-1] == 123 && s[i] == 125 {
				pos--
			} else if mas[pos-1] == 91 && s[i] == 93 {
				pos--
			} else if mas[pos-1] == 40 && s[i] == 41 {
				pos--
			} else {
				mas[pos] = s[i]
				pos++
			}
		}
	}

	if pos == 0 {
		return true, nil
	} else {
		return false, nil
	}
	return true, nil
}
