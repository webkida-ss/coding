import '../../../00_util/dart/generate.dart';

void selectionSort(List<int> nums){

}

void main(){
  var nums = [12, 11, 13, 5, 6, 7];
  selectionSort(nums);
  print(nums);

  List<int> testNums = GenerateRandomArray(6);
  selectionSort(testNums);
  print(testNums);
}