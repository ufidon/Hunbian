/* 类与对象
*/

class 狗{
    // 狗属性
    private int 龄;
    String 种;
    String 皮色;

    // 造函
    public 狗(int 龄1, String 种1, String 皮色1) {
        龄 = 龄1;
        种 = 种1;
        皮色 = 皮色1;
    }

    public void 设龄(int 龄1) {
        龄 = 龄1;
    }
    public int 龄(){
        return 龄;
    }
    public String 何种(){
        return 种;
    }
    public String 何色(){
        return 皮色;
    }    

}

public class 演示 {

    public static void main(String[] 行参集 ) {
        狗 我家狗 = new 狗(5, "看家狗", "黑");
        System.out.println("我家狗简介：");
        System.out.println("狗龄："+我家狗.龄());
        //System.out.println("狗龄："+我家狗.龄);    
        System.out.println("品种："+我家狗.种);
        System.out.println("皮色："+我家狗.皮色);
    }
}