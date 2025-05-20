// Define the Solution struct
struct Solution;

impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        let str_x = x.to_string();

        let mut left = 0;
        let mut right = str_x.len() - 1;
        while left < right {
            if str_x.chars().nth(left) != str_x.chars().nth(right) {
                return false
            }
            left += 1;
            right -= 1;
        }
        return true
    }
}

fn main() {
    let test_cases = vec![
        (121,true),
        (-121, false),
        (123, false),
        (7, true),
    ];
    for(number, expected) in test_cases {
        let result = Solution::is_palindrome(number);
        println!(
            "Is {} a palindrome? {} (expected: {})",
            number, result, expected
        );
    }
}