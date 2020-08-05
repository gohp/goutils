package str

/**
* @Author: Jam Wong
* @Date: 2020/8/4
 */

/*
Longest Palindromic Substring
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"
*/
func LongestPalindromic(input string) string {
	if len(input) < 2 {
		return input
	}
	longest := input[0:1]
	for i := 1; i < len(input); i++ {
		for rightStep := 0; rightStep < 2; rightStep++ {
			for p, q := i-1, i+rightStep; p >= 0 && q < len(input) && input[p] == input[q]; {
				if q-p+1 > len(longest) {
					longest = input[p : q+1]
				}
				p--
				q++
			}
		}
	}
	return longest
}
