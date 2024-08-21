package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"math"
)
//////////////////////////////////////////////////////////////////////////
func missingNumber(nums []int) int {
	little := nums[0]
	big := nums[0]
	for _, item := range nums {

		if item < little {
			little = item
		}
		if item > big {
			big = item
		}
	}
	slice := []int{}
	for i := little; i <= big; i++ {
		slice = append(slice, i)
	}
    for index,item:=range(slice) {
	   for _,items:=range(nums) {
		   if item==items{
			slice[index]=0
		   }
	    }
    }  
    for _,item:=range(slice) {
       if item!=0{
	      return item
        }
    }
	return 0
}
//////////////////////////////////////////////////////////////////////////
//136. Single Number
func singleNumber(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
				if count > 1 {
					nums[i] = 0
					nums[j] = 0
					count = 0
				}
			}
			if count > 1 {
				nums[i] = 0
				nums[j] = 0

			}
		}
		count = 0

	}
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			index = nums[i]
		}
	}
	fmt.Print(index)
}

//way2:

// func singleNumber(nums []int) int {
// 	numCount := make(map[int]int)


// 	for _, num := range nums {
// 		numCount[num]++
// 	}


// 	var result int
// 	for num, count := range numCount {
// 		if count == 1 {
// 			result = num
// 			break
// 		}
// 	}

// 	return result
// }

// //////////////////////////////////////////////////////////////////////////
// 75. Sort Colors
func SortColors(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for _, item := range nums {
		fmt.Println(item)
	}
}

// //////////////////////////////////////////////////////////////////////////
func countTestedDevices(batteryPercentages []int) int {
	max := batteryPercentages[0]
	for _, item := range batteryPercentages {
		if item > max {
			max = item
		}
	}
	return max
}

////////////////////////////////////////////////////////////////////////////
// func plusOne(digits []int) []int{
// 	num := 0
// 	for i := 0; i < len(digits); i++ {
// 		num = num*10 + digits[i]%10
// 	}
// 	num++
// 	slice:=[]int{}
// 	for num > 0 {
// 		digit := num % 10
// 		slice = append([]int{digit}, slice...)
// 		num /= 10
// 	}
// 	return slice
// }

func plusOne(digits []int) []int {

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}

//////////////////////////////////////////////////////////////////

//1: 88. Merge Sorted Array

func merge(nums1 []int, m int, nums2 []int, n int) {
	slice := []int{}
	slice2 := []int{}
	slice = nums1
	slice2 = nums2
	nums2 = nil
	nums1 = nil
	for i := 0; i < m; i++ {
		nums1 = append(nums1, slice[i])
	}
	for i := 0; i < n; i++ {
		nums1 = append(nums1, slice2[i])
	}
	sort.Ints(nums1)
	fmt.Printf("%v ", nums1)
}

//////////////////////////////////////////////////////////////////

//2:389. Find the Difference

func findTheDifference(s string, t string) byte {
	if s == "" {
		return t[0]
	}

	tmp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		tmp[s[i]-'a']++
		tmp[t[i]-'a']++
	}
	tmp[t[len(s)]-'a']++

	var found byte
	for b, c := range tmp {
		if c%2 != 0 {
			found = b + 'a'
			break
		}
	}

	return found
}

//////////////////////////////////////////////////////////////////

//3: 242. Valid Anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, c := range s {
		m[c]++
	}

	for _, c := range t {
		if _, ok := m[c]; !ok {
			return false
		} else {
			m[c]--
			if m[c] < 0 {
				return false
			}
		}
	}

	return true
}

//////////////////////////////////////////////////////////////////

//4:628. Maximum Product of Three Numbers

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	l := len(nums)
	max := nums[l-1] * nums[l-2] * nums[l-3]

	if nums[0] < 0 && nums[1] < 0 {
		c := nums[0] * nums[1] * nums[l-1]
		if c > max {
			max = c
		}
	}
	return max
}

//////////////////////////////////////////////////////////////////

//5: 258. Add Digits

func addDigits(num int) int {
//way1
 ren:=num
 reg:=0
 for i:=0;num>0;i++{
     reg+=num%10
 	num/=10
 }
 res:=0
 for i:=0;ren>0;i++{
 	res++
 	ren/=10
 }
 return res
//way2

//     if num == 0 {
//         return 0
//     }
//     if num % 9 == 0 {
//         return 9
//     }
//     return num % 9
 }

//////////////////////////////////////////////////////////////////

//6: 169. Majority Element

func majorityElement(nums []int) int {
//way 1
    	count:=0
 	repeater:=make(map[int]int)
 	for i:=0;i<len(nums);i++{
 		for j:=0;j<len(nums);j++{
 			if nums[i]==nums[j]{
 				count++;
 			}
 		}
 		repeater[nums[i]]=count
         count=0
 	}
 	var maxKey, maxValue int
 	for key, value := range repeater {
 		if value > maxValue {
 			maxValue = value
 			maxKey = key
 		}
 	}
    return maxKey

//way2
// 	candidate := nums[0]
// 	count := 1

// 	for i := 1; i < len(nums); i++ {
// 		if count == 0 {
// 			candidate = nums[i]
// 		}

// 		if nums[i] == candidate {
// 			count++
// 		} else {
// 			count--
// 		}
// 	}

// 	return candidate
 }

//////////////////////////////////////////////////////////////////

//7: 217. Contains Duplicate

func containsDuplicate(nums []int) bool {
count:=0
wg:=sync.WaitGroup{}
  for i:=0;i<len(nums);i++{
	wg.Add(1)
	go func(){
		for j:=0;j<len(nums);j++{
            if nums[i]==nums[j]{
               count++
           }
       }
		wg.Done()
	}()
   wg.Wait()
	if count>=2{
		return true
	}
	count=0
}
return false
//way 2
//     m:=make(map[int]bool)
//     for _,el:=range nums{
//         _,ok:=m[el]
//         if ok{
//             return true
//         }
//         m[el]=true
//     }
//     return false
 }
//////////////////////////////////////////////////////////////////

//8: 16. 3Sum Closest

func threeSumClosest(nums []int, target int) int {
	var ans int
	var imin int
	var imax int
	var difftemp int
	var sum int

	diff := 100000  //an unrechable diff value
	sort.Ints(nums) // sorted array
	leng := len(nums)

	for index := 0; index < leng-2; index++ {
		imin = index + 1
		imax = leng - 1
		for imin < imax {
			sum = nums[index] + nums[imin] + nums[imax]
			difftemp = sum - target

			//                 //make the diff aways positive
			if difftemp < 0 {
				difftemp = -difftemp
			}

			//                 //check if it is the smallest diff
			if difftemp < diff {
				diff = difftemp
				ans = sum
			}

			if sum == target {
				return target // best case

			} else if sum > target {
				imax-- // make the sum smaller
			} else {
				imin++ // make the sum bigger
			}
		}
	}
	return ans
}

//////////////////////////////////////////////////////////////////

//9: 28. Find the Index of the First Occurrence in a String

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	needleLength := len(needle)
	for i := 0; i <= len(haystack)-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}

	return -1
}

//////////////////////////////////////////////////////////////////
//10: 15. 3Sum

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicate elements
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

//////////////////////////////////////////////////////////////////

//11: 14. Longest Common Prefix

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[0 : len(prefix)-1]
		}
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

//////////////////////////////////////////////////////////////////

//12: 9. Palindrome Number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	part := 0
	for x > 0 {
		part = part*10 + x%10
		x /= 10

		if part == x || part == x/10 {
			return true
		}
	}

	return false
}

//////////////////////////////////////////////////////////////////

//13: 1. Two Sum

func twoSum(nums []int, target int) []int {
	keys := map[int]int{}
	for i, iteam := range nums {
		result := target - iteam
		if _, ok := keys[result]; ok {
			return []int{keys[result], i}
		}
		keys[iteam] = i
	}
	return []int{0, 0}
}

//////////////////////////////////////////////////////////////////

//14: 2114. Maximum Number of Words Found in Sentences

func mostWordsFound(sentences []string) int {
	max := 0
	for _, iteam := range sentences {
		len := len(strings.Split(iteam, " "))
		if len > max {
			max = len
		}
	}
	return max
}

// ////////////////////////////////////////////////////////////////
// 15: intersection-of-two-arrays
func intersection(nums1 []int, nums2 []int) []int {
	slice := []int{}
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				count++
				if count != 2 {
					slice = append(slice, nums2[j])
				}
			}
		}
		count = 0
	}
	uniqueNums := make(map[int]bool)
	result := []int{}
	for _, num := range slice {
		if !uniqueNums[num] {
			uniqueNums[num] = true
			result = append(result, num)
		}
	}

	return result
}
// ////////////////////////////////////////////////////////////////
//Submission Detail
func sumOfTheDigitsOfHarshadNumber(x int) int {
	sum := 0
	input := x
	for i := 1; x > 0; i++ {
		digit := x % 10
		sum += digit
		x /= 10
	}
	if input%sum != 0 {
		return -1
	}
	return sum
}
// ////////////////////////////////////////////////////////////////
//Find the Sum of Encrypted Integers
func sumOfEncryptedInt(nums []int) int {
    sum := 0
    for _, num := range nums {
        sum += encrypt(num)
    }
    return sum
}
func encrypt(num int) int {
    count, large := 0, 0
    for num > 0 {
        digit := num % 10
        if digit > large {
            large = digit
        }
        num /= 10
        count++
    }
    for i:=1; i<count; i++ {
        large = (large * 10) + (large % 10)
    }
    return large
}
// ////////////////////////////////////////////////////////////////
//69. Sqrt(x)
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }

    left, right := 1, x
    for left <= right {
        mid := left + (right - left) / 2
        if mid <= x / mid && (mid + 1) > x / (mid + 1) {
            return mid
        } else if mid > x / mid {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return -1  
}
// ////////////////////////////////////////////////////////////////
//Excel Sheet Column Title
func convertToTitle(columnNumber int) string {
	var result string

	for columnNumber > 0 {
		columnNumber-- // Adjust to 0-indexed
		remainder := columnNumber % 26
		result = string('A'+remainder) + result
		columnNumber /= 26
	}

	return result
}
// ////////////////////////////////////////////////////////////////
//Pow(x, n)
func myPow(x float64, n int) float64 {
   return math.Pow(x,float64(n))
}
////////////////////////////////////////////////////////////////
//power-of-two
func isPowerOfTwo(n int) bool {
	res:=n
	var count float64 =0
	if n == 1 {
		return true
	}
	for i := n; i >= 1; i/=2 {
		if i%2 == 0 {
             count++
		}
	}
	Result:=math.Pow(2, count)
	if Result==float64(res){
		return true
	}
	return false
}
////////////////////////////////////////////////////////////////
func Sum150(){
	Str1:="123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890"
	Str2:="123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890"
	var Result[]int

	for i:=0;i<len(Str1);i++{
		carryFlag := 0
		Num,err:=strconv.Atoi(string(Str1[i]))
		if err!=nil{
            fmt.Println(err)
        }
		Num2,err:=strconv.Atoi(string(Str2[i]))
		if err!=nil{
            fmt.Println(err)
        }
		Nums:=Num+Num2+carryFlag
		if Nums>=10{
			Result = append(Result,Nums%10)
			carryFlag = Nums / 10
		}else{
           Result=append(Result,Nums)
		}
	}
	fmt.Printf("%v \n",Result)
}
////////////////////////////////////////////////////////////////
func Sum150v2(){
		Str1 := "123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890"
	Str2 := "123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890123445678901234456789012344567890"
	str3 := ""

	if len(Str1) != len(Str2) {
		fmt.Println("Lengths of the input strings do not match")
		return
	}

	carryFlag := 0
	length := len(Str1)

	for i := length - 1; i >= 0; i-- {
		Num, err := strconv.Atoi(string(Str1[i]))
		if err != nil {
			fmt.Println(err)
			return
		}
		Num2, err := strconv.Atoi(string(Str2[i]))
		if err != nil {
			fmt.Println(err)
			return
		}
		Nums := Num + Num2 + carryFlag
		if Nums >= 10 {
			str3 = strconv.Itoa(Nums%10) + str3
			carryFlag = Nums / 10
		} else {
			str3 = strconv.Itoa(Nums) + str3
			carryFlag = 0
		}
	}

	// If there is any carry left, add it
	if carryFlag > 0 {
		str3 = strconv.Itoa(carryFlag) + str3
	}

	fmt.Println(str3)
}
////////////////////////////////////////////////////////////////
//Divide Two Integers
func divide(d int, di int) int {
     if d/di == 2147483648{
         return d/di-1
     }
     return d/di
}
////////////////////////////////////////////////////////////////
func moveZeroes(nums []int)  {
    count := 0
    for index, _ := range nums {
        if nums[index] != 0 {
            nums[count], nums[index] = nums[index], nums[count]
            count++
        }
    }
}
////////////////////////////////////////////////////////////////
