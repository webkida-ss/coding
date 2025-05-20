public class Easy9 {
    public static void main(String[] args) {
        Easy9 ins = new Easy9();
        System.out.println(ins.isPalindrome(121));
        System.out.println(ins.isPalindrome(1221));
        System.out.println(ins.isPalindrome(-121));
        System.out.println(ins.isPalindrome(10));
        System.out.println(ins.isPalindrome(1));
        System.out.println("aaa");
        System.out.println(ins.isPalindrome2(121));
        System.out.println(ins.isPalindrome2(1221));
        System.out.println(ins.isPalindrome2(-121));
        System.out.println(ins.isPalindrome2(10));
        System.out.println(ins.isPalindrome2(1));
    }

    public boolean isPalindrome(int x) {
        if(x<0){
            return false;
        }

        String str = Integer.toString(x);
        int length = str.length();

        for (int i=0;i<length/2;i++){
            // パースする必要はない
            // int start = Integer.parseInt(str.substring(i,i+1));
            // int end = Integer.parseInt(str.substring(length-i-1,length-i));
            char start = str.charAt(i);
            char end = str.charAt(length-i-1);
            if (start != end){
                return false;
            }
        }
        return true;
    }

    public boolean isPalindrome2(int x) {

        String strX = Integer.toString(x);

        int left = 0;
        int right = strX.length() - 1;

        while (left < right) {
            if (strX.charAt(left) != strX.charAt(right)) {
                return false;
            }
            left++;
            right--;
        }
        return true;
    }
}
