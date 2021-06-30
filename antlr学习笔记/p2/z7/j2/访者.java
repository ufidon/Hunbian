/*
java 访者 < 测试.配
*/
import java.util.Map;

import org.antlr.v4.misc.OrderedHashMap;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

/**
 * 访档者
 */
class 访档者 extends 配档BaseVisitor<Void> {
    Map<String,String> 属性集 = new OrderedHashMap<String,String>();

    public Void visit属性(配档Parser.属性Context ctx) { 
        属性集.put(ctx.M名().getText(), ctx.C串().getText());
        return null;
    }


    public void 印屏属性(){
        for (String 名 : 属性集.keySet()) {
            System.out.println(名+'='+属性集.get(名));
        }
    }
}

/**
 * 访者
 */
public class 访者 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        配档Lexer 取词器 = new 配档Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        配档Parser 树句器 = new 配档Parser(词流);
        
        ParseTree 树 = 树句器.档();

        // 造仿树者
        访档者 访者 = new 访档者();
        访者.visit(树);
        访者.印屏属性();
        System.out.println();      
    }
    
}