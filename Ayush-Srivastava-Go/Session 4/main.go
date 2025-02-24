package main

import "fmt"

func main() {


}

// Q.1
func reverse(nums []int) []int {

	for i := 0; i < len(nums)/2; i++ {
		j := len(nums) - 1 - i
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}

// Q.2
func sumSlice(nums []int) int {

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i] 
	}

	return sum
}

// Q.3
func maxMin(nums []int) (int, int) {

	maxVal := nums[0]
	minVal := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	return maxVal, minVal
}

// Q.4

func DoesExist(mp map[int]int, num int) bool {

	_, exists := mp[num]

	if exists {
		return true
	} else {
		return false
	}

}

func removeDuplicates(nums []int) []int {
	mp := make(map[int]bool)
	var slc []int

	for _, num := range nums {
		if !mp[num] { 
			slc = append(slc, num)
			mp[num] = true
		}
	}

	return slc
}

// Q.5
func rotateByK(nums []int, k int) []int {
	n := len(nums)
	temp := make([]int, n)

	copy(temp, nums)
	k = k % n

	for i := 0; i < n; i++ {
		nums[(i+k)%n] = temp[i]
	}

	return nums
}

// Q.6
func findEven(nums []int) []int {

	n := len(nums)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			res = append(res, nums[i])
		}
	}

	return res
}

// Q.7
func findOccurence(nums []int) {

	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}

	for k, v := range mp {
		fmt.Printf("\n%d occured %d times.", k, v)
	}
}

// Q.8
func wordFreq(sent string) {

	mp := make(map[string]int)

	word := ""

	for i:=0; i<len(sent); i++ {

		if(sent[i] != ' '){
			word += string(sent[i])
			continue
		}
		mp[word]++
		word = ""
	}

	if(word != ""){
		mp[word]++
	}

	for word, freq := range mp{

		fmt.Printf("\n%s occured %d times.", word, freq)

	}
}	

// Q.9(considering ' ' is not a character)

func countChar(s string) map[string]int {

	mp := make(map[string]int)

	for i:=0; i<len(s); i++ {
		if(s[i] != ' '){
			mp[string(s[i])]++
		}
	}

	return mp
}

// Q.10
func isPresent(mp map[int]int, target int) bool{

	_, exists := mp[target]

	if(exists){
		return true
	}else{
		return false
	}
}

func twoSum(nums []int, target int)(int, int){

	mp := make(map[int]int)


	for idx, num := range nums{
		if index, ok := mp[target - num]; ok{
			return index, idx
		}
		mp[num] = idx
	}

	return -1, -1

}

// Q.11
func isDuplicate(nums []int) bool {

	n := len(nums)

	mp := make(map[int]bool)

	for i:=0; i<n; i++{

		_, val := mp[nums[i]]

		if(val){
			return true
		}else{
			mp[nums[i]] = true
		}
	}

	return false
}

// Q.12
func uniqueChar(s string) int {

	mp := make(map[rune]int)

	for _, ch := range s {
		mp[ch]++
	}

	for idx, ch := range s{

		_, exists := mp[ch]

		if(!exists){
			return idx
		}
	}

	return 0;
}

// Q.13
func fib(n int) []int {
    if n <= 0 {
        return []int{}
    }
    
    if n == 1 {
        return []int{0}
    }

    res := make([]int, n) 
    res[0], res[1] = 0, 1

    for i := 2; i < n; i++ {
        res[i] = res[i-1] + res[i-2]
    }

    return res
}

// Q.14
func isPalindrome(num int) bool {

	if(num < 0 || (num % 10 == 0 && num != 0)){
		return false
	}

	rev := 0
	orig := num

	for num > 0{
		rev = rev*10 + num%10
		num /= 10
	}

	return rev == orig
}

// Q.15
func missingNum(nums []int) int {

	n := len(nums)

	sum := 0
	expected := (n*(n+1))/2

	for i:=0; i<n; i++{
		sum += nums[i]
	}

	return expected - sum
}


