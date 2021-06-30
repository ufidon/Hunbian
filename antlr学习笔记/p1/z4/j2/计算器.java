/** 计算器
 * 应用ANTLR4生成的取词器与树句器
 * 造：
 * javac *.java
 * 跑:
 * java 计算器 测试.式
*/

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;


public class 计算器 {
    public static void main(String[] 行参集) throws Exception {
        String 程档名 = "./测试.式";
        if (行参集.length > 0) {
            程档名 = 行参集[0];
        }

        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromFileName(程档名);
        算式Lexer 取词器 = new 算式Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        算式Parser 树句器 = new 算式Parser(词流);
        
        ParseTree 树 = 树句器.程式();
        算式访客 访客 = new 算式访客();
        访客.visit(树);
        

        /** 废弃，但依旧可用20210601
        InputStream 入流 = System.in;
        if (程档名 != null) {
            入流 = new FileInputStream(程档名);            
        }

        ANTLRInputStream 入 = new ANTLRInputStream(入流);
        算式Lexer 取词器 = new 算式Lexer(入);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        算式Parser 树句器 = new 算式Parser(词流);

        ParseTree 树 = 树句器.程式();
        */
    }
}