import java.util.*;

public abstract class 基域 implements 符域 {
    符域 父域; // 全域无父域，即为空
    Map<String, 符> 符集 = new LinkedHashMap<String, 符>();

    public 基域(符域 父域) { this.父域 = 父域;  }

    public 符 查(String 名) {
        符 录 = 符集.get(名);
        if ( 录!=null ) return 录;
        // 不在本域查父域
        if ( 父域 != null ) return 父域.查(名);
        return null; // 无记录
    }

    public void 入域(符 某符) {
        符集.put(某符.名, 某符);
        某符.域 = this; 
    }

    public 符域 父域() { return 父域; }

    public String toString() { return 域名()+":"+符集.keySet().toString(); }
}
