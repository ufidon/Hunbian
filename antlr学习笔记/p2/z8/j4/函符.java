import java.util.*;

public class 函符 extends 符 implements 符域 {
    Map<String, 符> 参集 = new LinkedHashMap<String, 符>();
    符域 父域;

    public 函符(String 名, 符型 返型, 符域 父域) {
        super(名, 返型);
        this.父域 = 父域;
    }

    public 符 查(String 名) {
        符 录 = 参集.get(名);
        if ( 录!=null ) return 录;
        if ( 父域() != null ) {
            return 父域().查(名);
        }
        return null; 
    }

    public void 入域(符 某符) {
        参集.put(某符.名, 某符);
        某符.域 = this; 
    }

    public 符域 父域() { return 父域; }
    public String 域名() { return 名; } // 以函名为域名

    public String toString() { return "函"+super.toString()+":"+参集.values(); }
}
