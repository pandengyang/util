package StringUtils

func ExistsInSlice(slice []string, value string) (exists bool) {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
