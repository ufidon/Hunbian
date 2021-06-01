/* 枚举
*/
class 周{
    enum 星期{日, 一,二,三,四,五,六}
    星期 几;
}
public class 枚举{
    public static void main(String[] args) {
        周 周几 = new 周();
        周几.几 = 周.星期.三;

        System.out.println("星期几？ "+周几.几);
    }
}