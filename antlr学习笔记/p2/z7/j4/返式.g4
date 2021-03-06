grammar 返式;

/* 给各选项加标签
antlr4 -visitor 签式.g4
grun 签式 句 -gui <<<$(echo -e "2+3*2\n 1+2*3")
*/

句 : 算式+ ;

算式 returns [int 值] // 返结值
  : 算式 算符=C乘 算式   #乘法
  | 算式 算符=J加 算式     #加法
  | Z整数                #整数
  ;  

C乘: '*' ;
J加 : '+' ;
Z整数 : [0-9]+ ;
H行尾符: '\r'?'\n';
K空格 : [ \t\r\n]+ -> skip ;