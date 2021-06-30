/**
 * 应用ANTLR4生成的取词器与树句器
*/

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;
import java.io.*;

public class 应用 {
    public static void main(String[] 行参集) throws Exception {
        String 程档名 = "./类.java";
        if (行参集.length > 0) {
            程档名 = 行参集[0];
        }

        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromFileName(程档名);
        Java9Lexer 取词器 = new Java9Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        Java9Parser 树句器 = new Java9Parser(词流);        
        ParseTree 树 = 树句器.compilationUnit(); // 语法分析

        ParseTreeWalker 行者 = new ParseTreeWalker();
        取接口 取 = new 取接口(树句器);
        行者.walk(取, 树);
        
    }
}