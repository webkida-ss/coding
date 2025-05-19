def binarySearch(list, target):
    left = 0
    right = len(list) - 1

    while left <= right:
        mid = left + (right - left) // 2
        midVal = list[mid]
        if midVal == target:
            return mid
        elif midVal < target:
            left = mid + 1
        else:
            right = mid - 1

    # 基底ケース
    return -1


print(binarySearch([1], 1))  # 0
print(binarySearch([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 3))  # 2
print(binarySearch([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 9))  # 8
print(binarySearch([1, 3, 5, 7, 9], 5))  # 2
print(binarySearch([2, 4, 6, 8, 10], 2))  # 0
print(binarySearch([1, 2, 3, 4, 5], 5))  # 4
print(binarySearch([10, 20, 30, 40, 50], 25))  # -1
print(binarySearch([], 1))  # -1
print(binarySearch([4], 4))  # 0
print(binarySearch([4], 3))  # -1
