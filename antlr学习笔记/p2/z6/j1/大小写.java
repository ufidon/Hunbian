// 测试非英语文字大小写

/**
 * 大小写
 */
public class 大小写 {

    public static void main(String[] args) {
        String 中文 = "中文";
        System.out.println(Character.isUpperCase(中文.charAt(0)));
        System.out.println(Character.isLowerCase(中文.charAt(0)));
    }
}