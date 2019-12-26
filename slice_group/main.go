package main

/*

Given a list of names, you need to organize each name within a slice based on
its length.
*/

func main() {

}

func groupByLength(nameSlice []string) [][]string {
	maxLen := getMaxLen(nameSlice)
	result := make([][]string, maxLen)

	for _, name := range nameSlice {
		result[len(name)-1] = append(result[len(name)-1], name)
	}

	return result
}

func getMaxLen(nameSlice []string) int {
	maxLen := 0
	for _, name := range nameSlice {
		if l := len(name); l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}
