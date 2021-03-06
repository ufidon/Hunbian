/** 简单的算式
测试：
javac S算式*.java
grun S算式 程式 -gui 测试.式
*/
grammar S算式;  // 语法名大写字母开头，需同文件名

/** 程式由句组成：句法需小写字母开头 */
程式: 句+;

/** 句由式构成 */
句:     式 H行尾符
    | L量名 '=' 式 H行尾符
    | H行尾符
    ;

/** 算式由子式、算符、变量、数和括式组成 */
式:     式 ('×'|'÷') 式
    | 式 ('+' | '-') 式
    | Z整数
    | L量名
    | '(' 式 ')'
    ;

/** 词由字组成: 词法需大写字母开头 
其中\u2e80-\u9fff为中日韩文字区，参考https://jrgraphix.net/r/Unicode/
可有否定式取所有文字，即非整数且非行尾符且非空格
*/    
L量名:  [a-zA-Z\u2e80-\u9fff]+;
Z整数:  [0-9]+;
H行尾符:    '\r'? '\n';
K空格:  [ \t]+ -> skip; // 忽略空格
