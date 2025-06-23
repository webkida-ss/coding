import '../../../00_util/dart/generate.dart';

List<int> selectionSort(List<int> nums){
  var l = nums.length;
  for(var i=0; i < l; i++){
    var tmpIndex = i;
    for(var j=i; j < l; j++){
      if (nums[j]< nums[tmpIndex]){
        tmpIndex = j;
      }
    }
    var tmp = nums[i];
    nums[i] = nums[tmpIndex];
    nums[tmpIndex] = tmp;
  }
  return nums;
}

void main(){
  var nums = [12, 11, 13, 5, 6, 7];
  selectionSort(nums);
  print(nums);

  List<int> testNums = GenerateRandomArray(6);
  selectionSort(testNums);
  print(testNums);
}