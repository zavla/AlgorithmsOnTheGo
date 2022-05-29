//Knuth-Morris-Pratt - a substring search of complexcity O(n+m).
//Note: strait search has complexity of O(n*m).
package main

func getPiSlice(s string) []int {
	pi := make([]int, len(s))
	i := 0 //used for prefix length (prefix that equals suffix)
	j := 1 //every char in s

	for j < len(s) {
		if s[i] == s[j] {
			pi[j] = i + 1 //sets prefix length to i+1 for the substring in the image
			i++
			j++
		} else if i == 0 {
			pi[j] = 0 //even prefix of length 1 is not equal suffix of length 1
			j++
		} else {
			// the last symbol of prefix that matched the suffix was i-1
			i = pi[i-1]
		}
	}
	return pi
}

func Substring(s, what string) int {
	pi := getPiSlice(what)

	j := 0
	for i := 0; i < len(s); {
		if s[i] == what[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = pi[j-1]
		}
		if j == len(what) {
			return (i - len(what))
		}
	}

	return -1
}

// func main() {
// 	//    01234567890123456789012345678
// 	s := "I say abc and abcabcc abcabcd"
// 	what := "abcabcd"
// 	fmt.Println(Substring(s, what))
// }
