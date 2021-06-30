grammar 配档;
/* 测：
antlr4 配档.g4
javac *.java
grun 配档 档 -gui <测试.配
*/

档: 属性+ 
    | EOF
    ;
属性: M名 '=' C串 '\r'?'\n' ;

M名: [a-zA-Z\u2e80-\u9fff]+;
C串: '"' .*? '"';

K空格: [ \t\n\r] -> skip;

