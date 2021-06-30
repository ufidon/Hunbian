/*
java 听者 < 测试.配
*/
import java.util.Map;

import org.antlr.v4.misc.OrderedHashMap;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

/**
 * 听档者
 */
class 听档者 extends 配档BaseListener {
    Map<String,String> 属性集 = new OrderedHashMap<String,String>();

    public void exit属性(配档Parser.属性Context ctx) { 
        属性集.put(ctx.M名().getText(), ctx.C串().getText());
    }

    public void exit档(配档Parser.档Context ctx) { 
        印屏属性();
    }

    public void 印屏属性(){
        for (String 名 : 属性集.keySet()) {
            System.out.println(名+'='+属性集.get(名));
        }
    }
}

/**
 * 听者
 */
public class 听者 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        配档Lexer 取词器 = new 配档Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        配档Parser 树句器 = new 配档Parser(词流);
        
        ParseTree 树 = 树句器.档();

        // 造树行者
        ParseTreeWalker 行者 = new ParseTreeWalker();
        行者.walk(new 听档者(), 树);
        System.out.println();      
    }
    
}