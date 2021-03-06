/* Terence Parr提取的一个C语言子集
测：
antlr4 子C.g4
javac *.java
grun 子C 档 -gui < 测试.c
*/
grammar 子C;

档:   (函宣 | 量宣)+ ;

量宣
    :   型 M名 ('=' 式)? ';'
    ;
型:   '浮型' | '整型' | '空型' ; // 变量类型

函宣
    :   型 M名 '(' 形参集? ')' 块 // "整型 加一(整型 甲) {...}"
    ;
形参集
    :   形参 (',' 形参)*
    ;
形参
    :   型 M名
    ;

块:  '{' 句* '}' ;   // 块可空
句:   块
    |   量宣
    |   '若' 式 '则' 句 ('否' 句)?
    |   '返' 式? ';' 
    |   式 '=' 式 ';' // 赋值句
    |   式 ';'          // 调函句
    ;

式:   M名 '(' 式列? ')'    // 算序最高的调函式，如：f(), f(x), f(1,2)
    |   M名 '[' 式 ']'         // 号式 a[i], a[i][j]
    |   '-' 式                // 负号
    |   '!' 式                // 逻辑非
    |   式 '*' 式
    |   式 ('+'|'-') 式
    |   式 '==' 式          // 算序最低的比较符
    |   M名                      // 变量
    |   Z整数
    |   '(' 式 ')'
    ;
式列 : 式 (',' 式)* ;   // 参列

M名  :   Z字 (Z字 | [0-9])* ;
fragment
Z字 : [a-zA-Z\u2e80-\u9fff] ;

Z整数 :   [0-9]+ ;

K空格  :   [ \t\n\r]+ -> skip ;

H行注释
    :   '//' .*? '\n' -> skip
    ;