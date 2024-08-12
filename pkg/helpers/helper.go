package helpers

import "strconv"

func GetIDFromParam(param string) (int, error) {
	return strconv.Atoi(param)
}
