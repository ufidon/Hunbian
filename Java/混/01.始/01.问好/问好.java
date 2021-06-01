/* 爪语问好
*/

public class 问好{
    public static void main(String []行参集) {
        if (行参集.length < 1) {
            System.out.println("用法： java 问好 谁");
            return;
        }
        System.out.println("您好，"+行参集[0]);
    }
}