from typing import List

def shift_zeros_to_the_end(nums: List[int]) -> None:
    insert_position = 0
    for current in range(len(nums)):
        if nums[current] != 0:
            nums[insert_position] = nums[current]
            insert_position += 1

    for i in range(insert_position, len(nums)):
        nums[i] = 0

if __name__ == "__main__":
    test_cases = [
        {"nums":[0, 1, 0, 3, 2], "expected": [1, 3, 2, 0, 0]},
    ]
    for test_case in test_cases:
        shift_zeros_to_the_end(test_case["nums"])
        print(test_case["nums"])