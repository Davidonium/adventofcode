package util

import (
	"fmt"
	"strconv"
)

func ParseInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to parse \"%s\" to an integer.", s))
	}

	return v
}

func ParseUInt64(s string) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse \"%s\" to uint64.", s))
	}

	return v
}

