package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	if len(s) == 0 {
		return result
	}
	findAddressPart([]byte(s), "", 0, &result)
	return result
}

func findAddressPart(s []byte, address string, count int, result *[]string) {
	if len(s) == 0 && count == 4 {
		address = string([]byte(address)[:len(address)-1])
		*result = append(*result, address)
		return
	}
	if len(s) == 0 || count == 4 {
		return
	}
	part := ""
	for i, b := range s {
		part += string(b)
		partNum, _ := strconv.Atoi(part)
		if partNum > 255 || (len(part) > 1) && part[0] == '0' {
			break
		}
		findAddressPart(s[i+1:], address+part+".", count+1, result)
	}
}

func main() {
	s := "010010"
	fmt.Println(restoreIpAddresses(s))
}
