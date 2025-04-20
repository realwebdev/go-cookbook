package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	len1 := len(str1)
	len2 := len(str2)
	commonLength := gcd(len1, len2)

	return str1[:commonLength]

}

func gcd(a, b int) int {
	for b != 0 {
		aa := a % b
		fmt.Println(aa)
		a, b = b, a%b
	}
	return a

}

func main() {
	str1 := "ABCABC"
	str2 := "ABCABCABCABCABCABCABC"

	ans := gcdOfStrings(str1, str2)
	fmt.Printf("gcd Of string --> %v\n", ans)
}
