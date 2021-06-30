/** 爪语类演示 */
import java.util.*;
/**
 * 类
 */
public class 类 {
    int 身高; // 厘米数
    String 名字; /* 姓+名 */
    List<Map<String /*宝贝名字*/, Integer /*宝贝数量*/>> 宝贝;

    // public 类(/* 构造函数 */String 名字, int 身高){
    //     this.名字 = 名字;
    //     this.身高 = 身高;
    // }

    public String 名(/* 取者无参 */){ return 名字; }
    public int 身高(){ return 身高; }
    public List<Map<String, Integer>> 宝贝(){ return 宝贝; };
    void 加宝贝(String 名, int 数){
        Map<String, Integer> 新 = new HashMap<String, Integer>();
        新.put(名, 数);
        宝贝.add(新);
    };
}