function isPalindrome(x: number): boolean {
  let strX = x.toString();

  let lett = 0;
  let right = strX.length - 1;
  while (lett < right) {
    if (strX[lett] != strX[right]) {
      return false;
    }

    lett++;
    right--;
  }

  return true;
}

console.log(isPalindrome(121));
console.log(isPalindrome(-121));
console.log(isPalindrome(10));
