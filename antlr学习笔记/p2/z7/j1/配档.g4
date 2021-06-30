grammar 配档;
@members{
    void 始档(){}
    void 终档(){}
    void 配置(Token 名, Token 值){}
}

/* 
档: {始档();} 属性+ {终档();};
属性: Mm '=' C1 '\n' {配置($Mm, $C1);};
Mm: [a-zA-Z\u2e80-\u9fff]+;
C1: '"' .*? '"';
K空格: [ \t\n\r] -> skip;
*/

// 在antlr4-4.9.2/tool/src/org/antlr/v4/parse/ActionSplitter.g
// 中增加中文支持即可。否则不认$M名, $C串
档: {始档();} 属性+ {终档();}
    | EOF
    ;
属性: M名 '=' C串 '\r'?'\n' {配置($M名, $C串);};

M名: [a-zA-Z\u2e80-\u9fff]+;
C串: '"' .*? '"';

K空格: [ \t\n\r] -> skip;

