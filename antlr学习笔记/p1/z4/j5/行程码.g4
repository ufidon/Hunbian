grammar 行程码;

/**
antlr4 -no-listener 行程码.g4
javac *.java
grun 行程码 档 -tree 行程档
*/

档: 行程+;

行程: Ju 队[$Ju.int];

队[int j]
locals [int i=1;]
    : ( {$i<=$j}? Ju {$i++;} )*
    ;

Ju: [0-9]+; // 行程距
K空格:  [ \t\n\r]+ -> skip;    