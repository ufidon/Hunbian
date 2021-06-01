// 雇员类
import java.io.*;

public class 雇员{
    String 名;
    int 龄;
    String 职;
    double 薪;

    // 造函
    public 雇员(String 名, String 职, double 薪, int 龄) {
        this.名 = 名;
        this.职 = 职;
        this.薪 = 薪;
        this.龄 = 龄;
    }

    // 方法
    public void 设龄(int 龄) {
        this.龄 = 龄;
    }
    public void 任职(String 职) {
        this.职 = 职;
    }
    public void 给薪(double 薪){
        this.薪 = 薪;
    }

    public String 信息() {
        return String.format("雇员%s的信息：\n\t年龄：%d\n\t职位：%s\n\t年薪：%.2f\n", 
        名, 龄, 职, 薪);
    }

}
