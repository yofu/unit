package unit

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func parseUnit(str string) ([]Unit, error) {
	name := strings.TrimSuffix(strings.TrimRightFunc(str, unicode.IsNumber), "^")
	pow := strings.TrimPrefix(strings.TrimLeftFunc(str, unicode.IsLetter), "^")
	var p int
	if pow == "" {
		p = 1
	} else {
		val, err := strconv.ParseInt(pow, 10, 64)
		if err != nil {
			return nil, err
		}
		p = int(val)
	}
	switch strings.ToLower(name) {
	default:
		return nil, fmt.Errorf("unknown unit")
	case "t", "tf":
		return Power(T, p), nil
	case "kg", "kgf":
		return Power(Kg, p), nil
	case "g", "gf":
		return Power(G, p), nil
	case "m":
		return Power(M, p), nil
	case "cm":
		return Power(Cm, p), nil
	case "mm":
		return Power(Mm, p), nil
	}
}

func Parse(str string) ([]Unit, []Unit, error) {
	lis := strings.Split(str, "/")
	if len(lis) > 2 {
		return nil, nil, fmt.Errorf("invalid format")
	}
	num := make([]Unit, 0)
	den := make([]Unit, 0)
	for _, s := range strings.Fields(lis[0]) {
		u, err := parseUnit(s)
		if err != nil {
			return nil, nil, err
		}
		num = append(num, u...)
	}
	if len(lis) < 2 {
		return num, den, nil
	}
	for _, s := range strings.Fields(lis[1]) {
		u, err := parseUnit(s)
		if err != nil {
			return nil, nil, err
		}
		den = append(den, u...)
	}
	return num, den, nil
}
