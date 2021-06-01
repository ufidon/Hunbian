/* 爪语古典数据结构
*/
import java.util.*;

/**
 * 古构
 */
public class 古构 {
    public static <古型> void 印集(古型[] 集){
        for (古型 元素 : 集) {
            System.out.printf("%s ", 元素.toString());
        }

        // 编译错误, 不能确定集[i]是啥?
        // for (int i=0; i<集.length; i++) {
        //     System.out.printf("位集%d信息：\n\t长度：%d\n\t尺寸:%d\n\t值:%s\n", 
        //     i, 集[i].length(), 集[i].size(), 集[i].toString());            
        // }   
        System.out.println();     
    }
    public static void main(String[] 行参集) {
        // 枚举
        System.out.printf("\n枚举演示：");
        Vector<String> 周名 = new Vector<String>(Arrays.asList(new String[]{"气","木","火","土","金","水","空"}));
        Enumeration 周 =周名.elements();
        int i = 0;
        while (周.hasMoreElements()) {            
            System.out.printf("%s日是一周的第%d天\n", 周.nextElement(), ++i);
        }


        // 位集
        System.out.printf("\n位集演示：");
        BitSet[] 位集 = {new BitSet(),new BitSet(32)};
        印集(位集);
        位集[0].set(5, 10);
        位集[1].set(23);
        印集(位集);

        // 矢量

    }
}
