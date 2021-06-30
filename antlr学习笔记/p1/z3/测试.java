/**
 * 测试antlr4生成的取词器与树句器
 * 造：
 * javac 测试.java
 * 用：
 * java 测试
 * # 键入
 * {1,2,3}
 * ctrl+D
 * # 得输出
 * (队 { (值 1) , (值 2) , (值 3) })
 * 
 * # 错误输入
 * {1,2
 * ctrl+D
 * # 得输出
 * line 2:0 extraneous input '<EOF>' expecting {',', '}'}
 * (队 { (值 1) , (值 2) <missing '}'>)
*/
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;
/**
 * 测试
 */
public class 测试 {
    public static void main(String[] 行参集) throws Exception {

        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        队转节Lexer 取词器 = new 队转节Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        队转节Parser 树句器 = new 队转节Parser(词流);
        
        ParseTree 树 = 树句器.队();
        System.out.println(树.toStringTree(树句器));        
    }    
}