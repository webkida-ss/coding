package algorithm.recursion.fibonacci;

public class Fibonacci {
    public static void main(String[] args) {
        System.out.println(fibonacci(0));
        System.out.println(fibonacci(1));
        System.out.println(fibonacci(2));
        System.out.println(fibonacci(3));
        System.out.println(fibonacci(4));
        System.out.println(fibonacci(5));
        System.out.println(fibonacci(10));
    }
    
    private static int fibonacci(int n){
        if (n <= 1) {
            return n;
        }
        return fibonacci(n-1)+fibonacci(n-2);
    }
}