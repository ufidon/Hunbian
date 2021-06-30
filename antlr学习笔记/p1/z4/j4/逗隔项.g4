grammar 逗隔项;
/**
antlr4 -no-listener 逗隔项.g4  不生成听者
*/
// 向树句器添加成员变量及定制构函
@parser::members{
    int 列;
    public 逗隔项Parser(TokenStream 词流, int 列){
        this(词流);
        this.列 = 列;
    } 
}

行集: (行 H行尾符)+;

行
locals [int liehao=0]
    :(  XIANG
        {
            $liehao++;
            if($liehao == 列) System.out.println($XIANG.text);
        }
    )+
    ;

// 词法
D逗号: [,，] ->skip; 
H行尾符:    '\r'? '\n';
XIANG:    ~[,，\r\n]+;