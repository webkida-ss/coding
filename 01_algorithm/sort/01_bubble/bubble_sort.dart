import '../../../00_util/dart/generate.dart';

void bubbleSort(List<int> nums){
  for(var i=0; i< nums.length; i++){
    for(var j=0; j< nums.length-i-1; j++){
      if(nums[j] > nums[j+1]){
        var tmp = nums[j];
        nums[j] = nums[j+1];
        nums[j+1] = tmp;
      }
    }
  }
}

void main(){
  var nums = [12, 11, 13, 5, 6, 7];
  bubbleSort(nums);
  print(nums);

  List<int> testNums = GenerateRandomArray(6);
  bubbleSort(testNums);
  print(testNums);
}