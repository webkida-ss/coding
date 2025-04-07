from typing import List

# Two Pointer でやるべし
def pair_sum_sorted(nums: List[int], target: int) -> List[int]:
    left, right = 0, len(nums) - 1
    while left < right:
        sum = nums[left] + nums[right]
        if sum == target:
            return [left, right]
        elif sum < target:
            left += 1
        elif sum > target:
            right -= 1
    return []


if __name__ == "__main__":
    test_cases = [
        {"nums": [-5, -2, 3, 4, 6], "target": 7, "expected": [2, 3]},
        {"nums": [1,1,1], "target": 2, "expected": [0, 1]}, # any index order
    ]
    for i, test_case  in enumerate(test_cases):
        nums = test_case["nums"]
        target = test_case["target"]
        expected = test_case["expected"]
        result = pair_sum_sorted(nums, target)

        print(f"Test case {i+1}: nums={nums}, target={target}")
        print(f"Got: {result}")
        # print(f"Expected: {expected}, Got: {result}")
        # print("Pass" if result == expected else "Fail")
        print()