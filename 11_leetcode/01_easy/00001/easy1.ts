class Solution {
  twoSum(nums: number[], target: number): number[] | null {
    const map = new Map<number, number>();
    for (let i = 0; i < nums.length; i++) {
      const num = nums[i];
      const complement = target - num;
      if (map.has(complement)) {
        return [map.get(complement)!, i];
      }
      map.set(num, i);
    }

    return null;
  }
}

const solution = new Solution();
const nums = [2, 7, 11, 15];
const target = 9;
const result = solution.twoSum(nums, target);
console.log(result);
