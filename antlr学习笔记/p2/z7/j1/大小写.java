// 测试非英语文字大小写

/**
 * 大小写
 */
public class 大小写 {
    public 大小写(String 名){this.名  = 名;}
    String 名;
    String 名(){return 名;}
    public static void main(String[] args) {
        String 中文 = "中文";
        System.out.println(Character.isUpperCase(中文.charAt(0)));
        System.out.println(Character.isLowerCase(中文.charAt(0)));
        System.out.println(Character.isLetter(中文.charAt(0)));
        System.out.println(Character.isLetterOrDigit(中文.charAt(0)));

        大小写 某大小写 = new 大小写("张三");
        System.out.println(某大小写.名());

    }
}