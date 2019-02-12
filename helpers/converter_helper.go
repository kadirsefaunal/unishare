package helpers

import (
	"fmt"
	"strconv"
)

func StringToUint(str string) uint {
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		fmt.Println("Convertion error!")
	}

	return uint(u64)
}
