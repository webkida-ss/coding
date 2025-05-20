def is_palindrome_valid(s: str) -> bool:    
    s = s.lower()
    left,right = 0, len(s)-1
    while left < right:
        while left < right and not s[left].isalnum():
            left += 1
        while left < right and not s[right].isalnum():
            right -= 1
        if s[left] != s[right]:
            return False
        left += 1
        right -=1
    return True

if __name__ == "__main__":
    test_cases = [
        {"s":"a dog! a panic in a pagoda.", "expected": True},
        {"s":"abc123", "expected": False},
    ]
    for test_case in test_cases:
        result = is_palindrome_valid(test_case["s"])
        print(result == test_case["expected"])