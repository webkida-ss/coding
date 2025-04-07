package main

import "fmt"

// nextPermutationは入力配列をその次の辞書順に大きい順列に変更します。
// そのような配置が不可能な場合、それは最も低い順序（つまり、昇順にソートされた状態）に再配置しなければなりません。
// 置換はその場で行わなければならず、追加のメモリは定数しか使用できません。
func nextPermutation(nums []int) {
	i := len(nums) - 2

	// nums[i] < nums[i+1]を満たす最大のiを見つける
	// 左が右より大きい場合数値を探す
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	// nums[i] < nums[j]を満たす最大のjを見つける
	// さっき見つけたiに対して、右からiより大きい最小のjでを探す
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// i+1以降を反転する
	reverse(nums, i+1)
}

// reverseは配列を指定した位置から反転します。
func reverse(nums []int, start int) {
	left, right := start, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {

	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) //[1 3 2]

	nums = []int{1, 3, 5, 4, 2}
	nextPermutation(nums)
	fmt.Println(nums) // [1 4 2 3 5]
}
