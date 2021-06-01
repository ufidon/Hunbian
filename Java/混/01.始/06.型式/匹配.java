/* 型式匹配
*/
import java.util.regex.*;

public class 匹配 {

    public static void main(String[] args) {
        String 行 = "你今天中奖了3000块，对吧？(⊙o⊙)?";
        String 型 = "(.*)(\\d+)(.*)";

        Pattern 模 = Pattern.compile(型);

        Matcher 配 = 模.matcher(行);
        if (配.find()) {
            for (int i=0; i< 配.groupCount(); i++) {
                System.out.printf("第%d匹配组:%s\n", i, 配.group(i));
            } 
            
        } else {
            System.out.println("未找到配型！");
        }
    }
}