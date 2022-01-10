package utils

import "fmt"

func ReqFmt(name, method string) string {
	return fmt.Sprintf("req: %s-%s", name, method)
}

func ResFmt(name, method string) string {
	return fmt.Sprintf("res: %s-%s", name, method)
}
