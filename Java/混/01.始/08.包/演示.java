/* 目录排布
.
├── 动物.java
├── 哺乳动物.java
├── 源
│   └── 生物
│       ├── 动物.java
│       └── 哺乳动物.java
├── 演示.class
├── 演示.java
├── 类
│   └── 生物
│       ├── 动物.class
│       └── 哺乳动物.class
└── 罐
    └── 生物.jar

编译: javac -cp ".:./类/" 演示.java
运行：
1) java -cp ".:./类" 演示
2) java -cp ".:./罐/生物.jar" 演示
*/
import 生物.*;

public class 演示{
    public static void main(String[] 行参集) {
        哺乳动物 哺 = new 哺乳动物();
        哺.吃();
        哺.睡();
    }
}