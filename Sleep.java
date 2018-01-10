package kablazegame;

/*
* Author: cmyui#5585 // https://akatsuki.pw
* Just a little thing so i don't have to catch the fucking exception every time I use Thread.sleep lol
*/
public class Sleep {
    public static void main(int x) {
        try {
            Thread.sleep(x);
        } catch ( java.lang.InterruptedException ie) {
    		System.out.println(ie);
        }
    }
}