package function

func GetPaginationOffset(fromItem int) int {
	return (fromItem - 1) * 10
}
