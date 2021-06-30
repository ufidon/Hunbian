/**
 * 应用ANTLR4生成的取词器与树句器
*/

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;
import java.io.*;

public class 应用 {
    public static void main(String[] 行参集) throws Exception {
        String 程档名 = "./测试.式";
        if (行参集.length > 0) {
            程档名 = 行参集[0];
        }

        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromFileName(程档名);
        S算式Lexer 取词器 = new S算式Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        S算式Parser 树句器 = new S算式Parser(词流);
        
        ParseTree 树 = 树句器.程式();
        System.out.println(树.toStringTree(树句器));
        

        /** 废弃，但依旧可用20210601
        InputStream 入流 = System.in;
        if (程档名 != null) {
            入流 = new FileInputStream(程档名);            
        }

        ANTLRInputStream 入 = new ANTLRInputStream(入流);
        S算式Lexer 取词器 = new S算式Lexer(入);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        S算式Parser 树句器 = new S算式Parser(词流);

        ParseTree 树 = 树句器.程式();
        System.out.println(树.toStringTree(树句器));
        */
    }
}