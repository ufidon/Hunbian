/*目录排布
.
├── 源
│   └── 生物
│       ├── 动物.java
│       └── 哺乳动物.java
├── 演示.class
├── 演示.java
└── 类
    └── 生物
        ├── 动物.class
        └── 哺乳动物.class
编译： javac -d 类 -sourcepath ./源 源/生物/*.java    
造罐：jar cf 生物.jar 生物/*   
*/
package 生物;

public class 哺乳动物 implements 动物{
    public void 吃(){
        System.out.println("哺乳动物在吃东西");
    }

    public void 睡(){
        System.out.println("哺乳动物在睡觉");
    }
}