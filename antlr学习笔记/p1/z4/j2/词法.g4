lexer grammar 词法;

/** 词由字组成: 词法需大写字母开头 
其中\u2e80-\u9fff为中日韩文字区，参考https://jrgraphix.net/r/Unicode/
可有否定式取所有文字，即非整数且非行尾符且非空格
*/    
L量名:  [a-zA-Z\u2e80-\u9fff]+;
Z整数:  [0-9]+;
H行尾符:    '\r'? '\n';
K空格:  [ \t]+ -> skip; // 忽略空格