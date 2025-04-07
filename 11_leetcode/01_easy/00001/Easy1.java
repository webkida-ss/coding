import java.util.HashMap;
import java.util.Map;

public class Easy1 {
    public static void main(String[] args) {
        Easy1 ins = new Easy1();
        int[] nums = {2, 7, 11, 15};
        int[] result = ins.twoSum(nums, 9);
        System.out.println(result[0] + " " + result[1]);
    }

    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < nums.length; i++) {
            int num = nums[i];
            int complement = target - num;
            if (map.containsKey(complement)) {
                return new int[] {map.get(complement), i};
            }
            map.put(num, i);
        }
        return null;
    }
}
