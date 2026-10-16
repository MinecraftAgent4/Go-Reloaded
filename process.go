package main

import (
	"strconv"
	"strings"
	"unicode"
)

func transform(text string) string {
	lines := strings.Split(text, "\n")
	for i, l := range lines {
		lines[i] = strings.Join(detectTransformInstructions(strings.Split(l, " ")), " ")

	}
	return strings.Join(lines, "\n")
}

func capitalize(w string) string {
	r := []rune(strings.ToLower(w))
	if len(r) == 0 {
		return w
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func detectTransformInstructions(text []string) []string {
	var oper []string
	for i, w := range text {
		if string(w[0]) == "(" {
			if string(w[len(w)-1]) == ")" {
				oper = []string{w[1 : len(w)-1], "1"}
			} else if string(w[len(w)-1]) == "," {
				oper = strings.Split(w[1:]+text[i+1][:len(text[i+1])-1], ",")
			}
			operCount, err := strconv.Atoi(oper[1])
			if err == nil {
				operCount = i - operCount
				switch oper[0] {
				case "up":
					for i2 := operCount; i2 < i; i2++ {
						text[i2] = strings.ToUpper(text[i2])
					}
				case "low":
					for i2 := operCount; i2 < i; i2++ {
						text[i2] = strings.ToLower(text[i2])
					}
				case "cap":
					for i2 := operCount; i2 < i; i2++ {
						text[i2] = capitalize(text[i2])
					}
				case "hex":
					for i2 := operCount; i2 < i; i2++ {
						num, err := strconv.ParseInt(text[i2], 16, 64)
						if err == nil {
							text[i2] = strconv.FormatInt(num, 10)
						}
					}
				case "bin":
					for i2 := operCount; i2 < i; i2++ {
						num, err := strconv.ParseInt(text[i2], 2, 64)
						if err == nil {
							text[i2] = strconv.FormatInt(num, 10)
						}
					}
				}
			}
		}
	}
	op := []string{}
	for _, i := range text {
		if string(i[0]) != "(" && string(i[len(i)-1]) != ")" {
			op = append(op, i)
		}
	}
	return op
}
