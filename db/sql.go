package db

import (
	"fmt"
	"strings"
)

func ConvertParams(sql string) string {
	n := 1

	for {
		i := strings.Index(sql, "?")

		if i == -1 {
			break
		}

		sql = strings.Replace(sql, "?", fmt.Sprintf("$%v", n), 1)
		n++
	}

	return sql
}
