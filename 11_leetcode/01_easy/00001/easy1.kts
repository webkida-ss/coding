class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val map = mutableMapOf<Int, Int>()

        for (i in nums.indices) {
            val num = nums[i]
            val complement = target - num
            if (map.containsKey(complement)){
                return intArrayOf(map[complement]!!, i)
            }
            map[num] = i
        }
        throw IllegalArgumentException("No two sum solution")
    }
}

val Solution = Solution()
val nums = intArrayOf(2, 7, 11, 15)
val target = 9
val result = Solution.twoSum(nums, target)
println(result.contentToString()) // [0, 1]
