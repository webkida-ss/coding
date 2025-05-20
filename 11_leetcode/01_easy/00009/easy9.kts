class Solution {
    fun isPalindrome(x: Int): Boolean {
        if(x < 0){
            return false;
        }
        val xString = x.toString()
        val length = xString.length

        for (i in 0..(length / 2)) {
            val start = xString[i]
            val end = xString[length-1-i]
            if(start != end){
                return false;
            }
        }

        return true;
    }
    fun isPalindrome2(x: Int): Boolean {
        val strX = x.toString()
        var left = 0
        var right = strX.length -1
        while (left < right) {
            if(strX[left] != strX[right]){
                return false
            }
            left++
            right--
        }
        return true
    }

}

val Solution = Solution()
println(Solution.isPalindrome(0))
println(Solution.isPalindrome(1))
println(Solution.isPalindrome(121))
println(Solution.isPalindrome(123))
println(Solution.isPalindrome(1221))
println()
println(Solution.isPalindrome2(0))
println(Solution.isPalindrome2(1))
println(Solution.isPalindrome2(121))
println(Solution.isPalindrome2(123))
println(Solution.isPalindrome2(1221))