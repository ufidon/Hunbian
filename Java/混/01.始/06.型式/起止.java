/* 型式匹配
*/
import java.util.regex.*;

/**
 * 起止
 */
public class 起止 {

    private static final String 规则式 = "\\b狗\\b";
    private static final String 输入 = "狗 野狗 狗 黑狗 狗 狗啊";

    public static void main(String[] args) {
        Pattern 模式 = Pattern.compile(规则式);
        Matcher 配 = 模式.matcher(输入);

        int 配数 = 0;
        System.out.printf("输入：%s\n", 输入);
        System.out.printf("模式：%s\n", 模式.toString());
        System.out.println("匹配项起[含]止[不含][起,止)索引号:");

        while (配.find()) {
            配数++;
            System.out.println("匹配项："+配数);
            System.out.println("起："+配.start());
            System.out.println("止："+配.end());
        }
    }
}