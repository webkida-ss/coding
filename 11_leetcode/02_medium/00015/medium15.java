import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Medium15 {
    public static void main(String[] args) {
        Medium15 ins = new Medium15();
        int[][] testCases = {
            {-1, 0, 1, 2, -1, -4}
            ,{-4, -1, -1, 0, 1, 2}
            ,{-4, -1, -1, -1, -1, 2}
            ,{0, 1, 1}
            ,{0, 0, 0}
        };
        for (int[] nums :testCases) {
            System.out.println(ins.threeSum(nums));
        }
    }

    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);

        int length = nums.length;
        List<List<Integer>> result = new ArrayList<>();
        for (int i = 0; i < length -2; i++) {
            int left = i+1;
            int right = length-1;

            if( i >0 && nums[i-1] == nums[i]){
                continue;
            }

            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];
                if(sum == 0){
                    result.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    while(left < right && nums[left] == nums[left+1]){
                        left++;
                    }
                    while (left < right && nums[right-1] == nums[right]) {
                        right--;
                    }
                    left++;
                    right--;
                }else if (sum < 0){
                    left++;
                }else{
                    right--;
                }
            }
        }

        return result;
    }

}
