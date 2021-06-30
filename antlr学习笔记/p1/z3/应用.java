/**
 * 将整数队列转换成字节
*/
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

/**
 * 转换
 */
class 转换 extends 队转节BaseListener {
    // 将‘{’转成 "
    @Override
    public void enter队(队转节Parser.队Context ctx) {
        System.out.print('"');
    }

    // 将‘}’转成 "
    @Override 
    public void exit队(队转节Parser.队Context ctx) {
        System.out.println('"');
    }

    // 将整数值转换成字节
    // 待做：队中队不能处理
    @Override 
    public void enter值(队转节Parser.值Context ctx) { 
        int 值 = Integer.valueOf(ctx.Z整数().getText());
        System.out.printf("\\u%04x", 值);
    }
    
}

/**
 * 应用
 */
public class 应用 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        队转节Lexer 取词器 = new 队转节Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        队转节Parser 树句器 = new 队转节Parser(词流);
        
        ParseTree 树 = 树句器.队();

        // 造树行者
        ParseTreeWalker 行者 = new ParseTreeWalker();
        行者.walk(new 转换(), 树);
        System.out.println();      
    }
    
}