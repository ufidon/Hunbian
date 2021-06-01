/* 狗类
*/
import java.io.*;

public class 狗 implements java.io.Serializable{
    String 名;
    String 主;
    int 编号;

    public 狗( String 主, String 名,int 编号) {
        this.主 = 主;
        this.名 = 名;
        this.编号 = 编号;
    }
    
    public String toString() {
        return String.format("%s是%s的狗，编号%d\n", 名,主,编号);
    }
}