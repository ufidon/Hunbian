/* 爪语集(Collection)之演示
* 不指定模参，则出错： compile with -Xlint:unchecked
*/
import java.util.*;

/**
 * 集
 */
public class 集 {

    public static <型 extends Collection> void 印集(String 集名, 型 集) {
        System.out.printf("\n%s共有%d个元素：\n", 集名, 集.size());
        System.out.printf("%s\n", 集.toString());
    }
    public static <型 extends Map> void 印地图(String 图名, 型 地图) {
        System.out.printf("\n%s共有%d个元素：\n", 图名, 地图.size());
        System.out.printf("%s\n", 地图.toString());
    }    
    public static void main(String[] 行参集) {
        // 天、地、水、火、雷、风、山、泽
        String[] 卦名 = {"乾","坤","坎","离","震","巽","艮","兑"};
        String[] 卦象 = {"天","地","水","火","雷","风","山","泽"};

        // ArrayList
        List<String> 列 = new ArrayList<String>(卦名.length);
        for (int i = 0; i < 卦名.length; i++) {
            列.add(卦名[i]);
        }        
        印集("列", 列);

        // HashSet
        Set<String> 集合 = new HashSet<String>(卦名.length);
        for (String 卦 : 列) {  
            集合.add(卦);
        }
        // for(Iterator<String> i=列.iterator(); i.hasNext();){
        //     集合.add(i.next());
        // }
        // for (int i = 0; i < 卦名.length; i++) {
        //     集合.add(卦名[i]);
        // }
        印集("集合", 集合);

        // HashMap
        
        HashMap<String,String> 地图 = new HashMap<String,String>(卦名.length);
        for (int i=0; i<卦象.length; i++) {
            地图.put(卦名[i], 卦象[i]);
        }
        印地图("散列地图", 地图);
        
    }
}