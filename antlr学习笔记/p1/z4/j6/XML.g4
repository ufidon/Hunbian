lexer grammar XML;
/**
antlr4 XML.g4
javac *.java
grun XML tokens -tokens 测试.xml
*/

// 两种取词模式：默认（非签）和签
Q签头   :   '<'     ->  pushMode(Q签内);    // 进入签内模式
Z注释   :   '<!--' .*? '-->'    -> skip;
T特符   :   '&' [a-z]+ ';';
W文本   :   ~('<'|'&')+;

// 签内模式
mode Q签内;
Q签尾   :   '>'     -> popMode; // 返回签外模式
X斜签尾 :   '/>'    -> popMode;
D等于:      '=';
C串:        '"' .*? '"';
X斜名:      '/' M名;
M名 :       (Z字母|H汉字) (H汉字|Z字母|S数字)*;
K空格:      [ \t\r\n]   ->skip;

fragment
H汉字   :   [\u2e80-\u9fff];
fragment
Z字母   :   [a-zA-Z];
fragment
S数字   :   [0-9];
