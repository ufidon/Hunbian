grammar 易;

程:   类定义+ ; 

类定义
    :   '类' M名 '{' 成员+ '}' 
        {System.out.println("类 "+$M名.text);}
    ;

成员
    :   '整型' M名 ';'                       // 成员变量定义
        {System.out.println("成员变量 "+$M名.text);}
    |   '整型' 方=M名 '(' M名 ')' '{' j句 '}' // 成员方法定义
        {System.out.println("成员方法: "+$方.text);}
    ;

j句:   式 ';'
        {System.out.println("式: "+$j句.text);}
    |   M名 '=' 式 ';'
        {System.out.println("赋式: "+$j句.text);}
    ;

式:   Z整 
    |   M名 '(' Z整 ')'
    ;

Z整 :   [0-9]+ ;
M名  :   [a-zA-Z\u2e80-\u9fff]+ ;
K空格  :   [ \t\r\n]+ -> skip ;