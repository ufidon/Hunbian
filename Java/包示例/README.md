# 中英混编
有容乃大, 中英混合编程之Java.

##11102016.Java包与对应结构
```bash
├─源码
│  └─中国
│      └─广东
│          └─深圳
│              └─动物园
│                      动物.java
│                      哺乳动物.java
│
└─类群
    └─中国
        └─广东
            └─深圳
                └─动物园
                        动物.class
                        哺乳动物.class
```

```java
package 中国.广东.深圳.动物园;
```
##编译:
```bash
	 javac -encoding "UTF-8" -d 类群 -sourcepath 源码 源码\中国\广东\深圳\动物园\动物.java
	 javac -encoding "UTF-8" -d 类群 -sourcepath 源码 源码\中国\广东\深圳\动物园\哺乳动物.java
```	 
##运行:
```bash
	java -cp 类群 中国.广东.深圳.动物园.哺乳动物
```