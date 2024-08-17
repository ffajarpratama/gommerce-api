package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func AddPhoneCode(str string) string {
	if strings.HasPrefix(str, "0") {
		str = "62" + str[1:]
	}

	return str
}

func RemoveSpecialCharacters(str string) string {
	str = strings.ToLower(str)

	// replace multiple trailing hyphens with a single hyphen
	reg := regexp.MustCompile(`(\s*-|\s+-\s*|\s+-)\s*`)
	str = reg.ReplaceAllString(str, "-")

	// replace all non-alphabet characters with hyphens
	reg = regexp.MustCompile(`[^a-zA-Z]+`)
	str = reg.ReplaceAllString(str, "-")

	// remove leading and trailing hyphens
	str = strings.Trim(str, "-")

	// replace consecutive hyphens with a single hyphen
	reg = regexp.MustCompile("-+")
	str = reg.ReplaceAllString(str, "-")

	return str
}

func Slugify(str string) string {
	str = strings.ToLower(str)

	reg := regexp.MustCompile(`(\s*-|\s+-\s*|\s+-)\s*`)
	str = reg.ReplaceAllString(str, "-")

	reg = regexp.MustCompile(`[^a-zA-Z0-9-]+`)
	str = reg.ReplaceAllString(str, "-")
	str = strings.Trim(str, "-")

	reg = regexp.MustCompile("-+")
	str = reg.ReplaceAllString(str, "-")

	return fmt.Sprintf("%s-%s", str, RandomNumber(5))
}

func ParseZero(in float64, sep string) (res string) {
	res = strconv.FormatFloat(in, 'f', 0, 64)
	length := len(res)

	for i := length - 3; i > 0; i -= 3 {
		res = res[:i] + sep + res[i:]
	}

	return res
}
