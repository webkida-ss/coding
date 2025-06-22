import '../../../00_util/dart/generate.dart';

void mergeSort(List<int> nums){
  var l = nums.length;
  if (l <= 1){
    return;
  }

  var mid = l ~/2;
  // var list = List<Type>.filled(size, initialValue);
  var(left,right) = (nums.sublist(0, mid),nums.sublist(mid, l));

  mergeSort(left);
  mergeSort(right);

  merge(nums,left,right);
}

void merge(List<int> nums,List<int> left, List<int> right){
  var (i,j,k) = (0,0,0);
  while (i < left.length && j < right.length){
    var(lv,rv) = (left[i],right[j]);
    if (lv < rv){
      nums[k] = lv;
      i++;
    }else{
      nums[k] = rv;
      j++;
    }
    k++;
  }
  while (i < left.length){
    nums[k] = left[i];
    i++;
    k++;
  }
  while (j < right.length){
    nums[k] = right[j];
    j++;
    k++;
  }
}

void main(){
  var nums1 = GenerateRandomArray(10);
  mergeSort(nums1);
  print(nums1);
}