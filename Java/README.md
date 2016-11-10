# 中英混编
有容乃大, 中英混合编程之Java.

##11102016.Java包与对应结构

>├─classes
>│  └─cis
>│      └─edu
>│          └─edgewood
>│              └─animals
>│                      Animal.class
>│                      MammalInt.class
>│
>└─sources
>    └─cis
>        └─edu
>            └─edgewood
>                └─animals
>                        Animal.java
>                        MammalInt.java

```java
package cis.edu.edgewood.animals;
```
##编译:
	 javac -verbose -d classes -sourcepath sources sources\cis\edu\edgewood\animals\MammalInt.java -cp classes
##运行:
	 java -verbose -cp classes cis.edu.edgewood.animals.MammalInt
