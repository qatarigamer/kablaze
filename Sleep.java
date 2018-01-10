public class Sleep {
    public static void main(int x) {
        try {
            Thread.sleep(x);
        } catch ( java.lang.InterruptedException ie) {
    		System.out.println(ie);
        }
    }
}