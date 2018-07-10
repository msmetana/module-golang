package cipher

type MyCaesar struct{}

type MyShift struct {
	shift int
}

type MyVigenere struct {
	key string
}

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func (c MyCaesar) Encode(p string) string {
	s := ReadyToEncode(p)
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] + 3

		if tmp > 122 || (tmp > 90 && tmp < 97) {
			tmp -= 26
		}
		str[i] = tmp
	}

	return string(str)
}

func (c MyCaesar) Decode(s string) string {
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] - 3

		if tmp < 65 || (tmp > 90 && tmp < 97) {
			tmp += 26
		}
		str[i] = tmp
	}

	return string(str)
}

func NewCaesar() Cipher {
	return MyCaesar{}
}

func (sh MyShift) Encode(p string) string {
	s := ReadyToEncode(p)
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] + uint8(sh.shift)

		if tmp > 122 || (tmp > 90 && tmp < 97) {
			if sh.shift > 0 {
				tmp -= 26
			} else {
				tmp += 26
			}
		}
		str[i] = tmp
	}

	return string(str)
}

func (sh MyShift) Decode(s string) string {
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] - uint8(sh.shift)

		if sh.shift > 0 {
			if tmp < 65 || (tmp > 90 && tmp < 97) {
				tmp += 26
			}
		} else {
			if tmp > 122 || (tmp > 90 && tmp < 97) {
				tmp -= 26
			}
		}
		str[i] = tmp
	}

	return string(str)
}

func NewShift(sh int) Cipher {
	if sh > 25 || sh < -25 || sh == 0 {
		return nil
	}

	return MyShift{sh}
}

func invalid_key(s string) bool {
	flag := false

	if len(s) < 1 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || (s[i] > 40 && s[i] < 91) {
			return false
		}

		if s[i] != 'a' {
			flag = true
		}
	}
	return flag
}

func (v MyVigenere) Encode(p string) string {
	s := ReadyToEncode(p)
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] + uint8(v.key[i%len(v.key)]) - uint8('a')

		if tmp > 122 {
			tmp -= 26
		}
		str[i] = tmp
	}

	return string(str)
}

func (v MyVigenere) Decode(s string) string {
	str := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		tmp := s[i] - uint8(v.key[i%len(v.key)]) + uint8('a')

		if tmp < 97 {
			tmp += 26
		}

		str[i] = tmp
	}

	return string(str)
}

func NewVigenere(key string) Cipher {
	if !invalid_key(key) {
		return nil
	}

	return MyVigenere{key}
}

func ReadyToEncode(s string) string {
	count := 0
	j := 0

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			count++
		}
	}

	str := make([]byte, count)

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			if s[i] >= 'A' && s[i] <= 'Z' {
				str[j] = s[i] - 'A' + 'a'
			} else {
				str[j] = s[i]
			}
			j++
		}
	}
	return string(str)
}
