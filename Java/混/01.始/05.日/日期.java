/* 日期与时间
*/

import java.util.*;
import java.text.*;

public class 日期{
    public static void main(String[] 行参集) {
        SimpleDateFormat 日式 = new SimpleDateFormat("yyyy-MM-dd");
        String 输入 = 行参集.length == 0? "2020-12-21":行参集[0];

        System.out.println(输入+"被分析为");
        Date 日;
        try {
            日 = 日式.parse(输入);
            System.out.println(日);
        } catch (Exception e) {
            System.out.println("日期格式不符合"+日式);
        }
    }
}