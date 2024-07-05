package main

import "fmt"

func main(){
	// s1 := "abcdabcabah"
	// s1 := "dvdf"
	// s1 := "asjrgapa"
	s1 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	// s1 := "hhhhhhhhhhhhhhhhh"
	fmt.Println(lengthOfLongestSubstring(s1))
}

func lengthOfLongestSubstring(s string) int {
    stringArr := stringToList(s)
	subString := []string{}
	mainString := [][]int{}
	subStringIndex := []int{}
	for j := 0; j < len(stringArr);j++{
		for i,ch := range stringArr[j:]{
			fmt.Println(stringArr[j:])
		
			if !(isInsubString(string(ch),subString)){			
				subString = append(subString, string(ch))
				subStringIndex = append(subStringIndex, i+j)
			} else {
				// fmt.Println(subString)
				mainString = append(mainString, subStringIndex)
				subStringIndex = []int{i}
				subString = []string{ch}
			}
			if i == len(stringArr[j:])-1{
				mainString = append(mainString, subStringIndex)
				subString = []string{}
				subStringIndex = []int{}
			}
		}
		subStr := ""
		for i := 0; i<len(mainString[0]);i++{
			subStr += string(s[i])
		}
		fmt.Println(len(subStr))
		return 2
	}
	
	longerArr := []int{}
	for i,_ := range mainString{
		if len(mainString[i]) > len(longerArr){
			longerArr = mainString[i]
		}
	}

	fmt.Println(longerArr)

	// for i,ch := range subStringIndex {
	// 	if i != 0{
	// 		difference := subStringIndex[i] - subStringIndex[i-1]
	// 	}
	// }
	return len(mainString)
}

func isInsubString(s string, subArr []string) bool {
	for _,ch := range subArr{
		if ch == s{
			return true
		}
	}
	return false
}

func stringToList(s string) []string {
	res := []string{}
	for _,ch := range s{
		res = append(res, string(ch))
	}
	return res
}



// func lengthOfLongestSubstring(s string) int {
//     charIdx := make(map[rune]int, len(s))
//     maxLen, pos := 0, 0

//     for i, char := range s {
//         if idx, ok := charIdx[char]; ok && idx >= pos {
//            pos = idx + 1 
//         }

//         length := i - pos + 1
//         if maxLen < length {
//             maxLen = length
//         }

//         charIdx[char] = i
//     }

//     return maxLen
// }

