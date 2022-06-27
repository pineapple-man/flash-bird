package locate

func Locate(name string) string {

	return ""
}

func Exist(name string) bool {
	return Locate(name) != ""
}
