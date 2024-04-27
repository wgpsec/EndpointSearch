package Compare

// IsStringInStringArray test if str string is contained in strArray []string
// return true when contains
func IsStringInStringArray(str string, strArray []string) bool {
	for _, v := range strArray {
		if v == str {
			return true
		}
	}
	return false
}

func RemoveDuplicates(reqList []string) []string {
	encountered := map[string]bool{}
	var result []string
	for v := range reqList {
		if !encountered[reqList[v]] {
			encountered[reqList[v]] = true
			result = append(result, reqList[v])
		}
	}
	return result
}
