package common

// 增
func Add(slice []string, value string) []string {
	return append(slice, value)
}

// 删
func Remove(slice []interface{}, i int) []interface{} {
	return append(slice[:i], slice[i+1:]...)
}

// 改
func Update(slice []interface{}, index int, value interface{}) {
	slice[index] = value
}

// 查
func Find(slice []interface{}, index int) interface{} {
	return slice[index]
}