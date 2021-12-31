package helpers

import "strings"

func ByteArrayToString(ba []byte) string {
	return string(ba)
}

func SplitStringByByteArray(ba []byte, sep string) []string {
	str := ByteArrayToString(ba)
	return strings.Split(str, sep)
}
