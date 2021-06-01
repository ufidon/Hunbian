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
造罐： jar cf ../罐/生物.jar 生物/*   
Path = 罐/生物.jar
Type = zip
Physical Size = 1072

   Date      Time    Attr         Size   Compressed  Name
------------------- ----- ------------ ------------  ------------------------
2021-05-28 15:41:06 D....            0            2  META-INF
2021-05-28 15:41:06 .....           67           66  META-INF/MANIFEST.MF
2021-05-28 09:27:30 .....          134          117  生物/动物.class
2021-05-28 09:27:30 .....          528          343  生物/哺乳动物.class
------------------- ----- ------------ ------------  ------------------------

*/
package 生物;

interface 动物{
    public void 吃();
    public void 睡();
}

