package database

import (
	"slices"
	"strings"
)

func JoinColumns(columns []string) string {
	return strings.Join(columns, ", ")
}

func GetValueHolder(count int) string {
	return strings.Join(slices.Repeat([]string{"?"}, count), ", ")
}
