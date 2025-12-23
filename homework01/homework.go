package homework01

import (
	"fmt"
	"strconv"
)

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	existMap := make(map[int]int, len(nums))
	for _, e := range nums {
		existMap[e]++
	}
	var once int
	for k, v := range existMap {
		if v == 1 {
			once = k
			break
		}
	}
	existMap = nil
	return once
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	if 0 <= x && x <= 9 {
		return true
	}
	str := strconv.Itoa(x)
	strlen := len(str)
	res := true
	for i := 0; i <= (strlen / 2); i++ {
		if str[i] != str[strlen-i-1] {
			res = false
			break
		}
	}
	return res
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	slic := make([]rune, len(s)+1)
	var index int
	for _, v := range s {
		switch v {
		case '(':
			index++
			slic[index] = '('
		case ')':
			if index > 0 && slic[index] == '(' {
				index--
			} else {
				return false
			}
		case '{':
			index++
			slic[index] = '{'
		case '}':
			if index > 0 && slic[index] == '{' {
				index--
			} else {
				return false
			}
		case '[':
			index++
			slic[index] = '['
		case ']':
			if index > 0 && slic[index] == '[' {
				index--
			} else {
				return false
			}
		}
	}
	return index == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}
	slic := []rune(strs[0])
	var index = len(slic)
	for i, ele := range strs {
		if i > 0 {
			eleSlic := []rune(ele)
			j := 0
			for ; j < len(eleSlic); j++ {
				if j < index && eleSlic[j] != slic[j] {
					break
				}
			}
			if j < index {
				index = j
			}
			if j == 0 {
				return ""
			}
		}
	}
	return string(slic[:index])
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	num := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 && num == 1 {
			digits[i] = 0
		} else {
			digits[i]++
			num = 0
			break
		}
	}
	if num == 1 {
		digits = append([]int{1}, digits[:]...)
	}
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	nl := len(nums)
	if nl == 1 {
		return 1
	}
	index := 0
	indexN := 1
	step := 0
	for {
		if indexN >= nl {
			break
		}
		if nums[index] == nums[indexN] {
			step++
			indexN++
		} else {
			if step > 0 {
				for j := index + 1; j < nl; j++ {
					if step+j < nl {
						nums[j] = nums[step+j]
					} else {
						nums[j] = nums[nl-1]
					}

				}
				step = 0
			}
			index++
			indexN = index + 1
		}
	}
	index++
	return index
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	ilen := len(intervals)
	if ilen == 1 {
		return intervals
	}
	// 选择排序
	var min int
	for i := 0; i < ilen-1; i++ {
		min = i
		for j := i + 1; j < ilen; j++ {
			if intervals[min][0] > intervals[j][0] {
				min = j
			}
		}
		if min != i {
			temp := intervals[i]
			intervals[i] = intervals[min]
			intervals[min] = temp
		}
	}
	fmt.Println("sort res=", intervals)
	res := [][]int{}
	sindex := 0
	nindex := 1
	start := intervals[sindex]
	temp := [2]int{start[0], start[1]}
	for {
		next := intervals[nindex]
		// fmt.Println("before, temp=", temp,", next=", next,", nindex=", nindex)
		if next[0] >= temp[0] && next[0] <= temp[1] {
			if next[1] > temp[1] {
				temp[1] = next[1]
			} else {
				temp[1] = temp[1]
			}
			nindex++
		} else {
			// fmt.Println("res append 1, temp=", temp)
			res = append(res, []int{temp[0], temp[1]})
			// fmt.Println("res append 1 after, res=", res)
			sindex = nindex
			nindex = sindex + 1
			temp = [2]int{intervals[sindex][0], intervals[sindex][1]}
		}
		if nindex >= ilen {
			// fmt.Println("res append 2, temp=", temp)
			res = append(res, []int{temp[0], temp[1]})
			// fmt.Println("res append 2 after, res=", res)
			break
		}
	}
	return res
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	ilen := len(nums)
	for i := 0; i < ilen-1; i++ {
		for j := i + 1; j < ilen; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
