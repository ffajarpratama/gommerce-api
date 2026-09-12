package util

import (
	"fmt"
	"strings"
)

func CalculateOffset(page, limit int) (result int) {
	result = (page - 1) * limit
	return
}

func GetOrderByColumn(sort string) (result string) {
	if strings.Contains(sort, "-") {
		return fmt.Sprintf("%s DESC", sort[1:])
	}

	return sort
}

func TransformSortClause(column, sort string) (result string) {
	if sort == "latest" {
		return fmt.Sprintf("%s DESC", column)
	}

	return column
}
