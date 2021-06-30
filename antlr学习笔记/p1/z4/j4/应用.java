/**
 * 应用ANTLR4生成的取词器与树句器
 * 造：javac *.java
 * 跑：java 应用 客户群.csv 1
*/

import org.antlr.v4.runtime.*;

public class 应用 {
    public static void main(String[] 行参集) throws Exception {
        String 程档名 = "./客户群.csv";
        int 列号 = 1;

        if (行参集.length > 0) {
            程档名 = 行参集[0];
        }
        if (行参集.length > 1) {
            列号 = Integer.valueOf(行参集[1]);
        }

        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromFileName(程档名);
        逗隔项Lexer 取词器 = new 逗隔项Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        逗隔项Parser 树句器 = new 逗隔项Parser(词流, 列号);        

        树句器.setBuildParseTree(false);
        树句器.行集();
        
    }
}