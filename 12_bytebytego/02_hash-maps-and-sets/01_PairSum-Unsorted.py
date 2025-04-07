from typing import List

def pair_sum_unsorted(nums: List[int], target: int) -> List[int]:
    num_dict = {}
    for i in range(len(nums)):
        num = nums[i]
        complement = target - num
        if complement in num_dict:
            return [num_dict[complement], i]
        num_dict[num] = i
    return []

if __name__ == "__main__":
    test_cases = [
        {"nums": [-1, 3, 4, 2], "target": 3, "expected": [0, 2]},
    ]
    for i, test_case  in enumerate(test_cases):
        nums = test_case["nums"]
        target = test_case["target"]
        expected = test_case["expected"]
        result = pair_sum_unsorted(nums, target)

        print(f"Test case {i+1}: nums={nums}, target={target}")
        print(f"Expected: {expected}, Got: {result}")
        print("Pass" if result == expected else "Fail")
        print()