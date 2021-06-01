/* 雇员类演示
*/
import java.io.*;
public class 演示{
    public static void main(String[] 行参集) {
        // 创建两个雇员
        雇员 小张 = new 雇员("张天师", "程序员", 70000, 25);
        雇员 小李 = new 雇员("李世民", "老板", 900000, 50);

        System.out.println(小张.信息());
        System.out.println(小李.信息());
        char 雪人 ='☃';
        System.out.println(雪人);
    }
}