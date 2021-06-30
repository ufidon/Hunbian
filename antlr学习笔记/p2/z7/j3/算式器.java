/*
java 算式器 <<<$(echo 4*5)
*/

import java.util.Stack;

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

/**
 * 算式
 */
class 算式 extends 式BaseListener {
    Stack<Integer> 叠 = new Stack<Integer>();

    public void exit算式(式Parser.算式Context ctx) { 
        if (ctx.getChildCount() == 3) {
            int 右数 = 叠.pop();
            int 左数 = 叠.pop();
            if (ctx.算符.getType() == 式Parser.C乘) {
                叠.push(左数*右数);
            }else{
                叠.push(左数+右数);
            }
        }        
    }
    public void visitTerminal(TerminalNode 结){
        Token 词 = 结.getSymbol();
        if (词.getType() == 式Parser.Z整数) {
            叠.push(Integer.valueOf(词.getText()));
        }
    }
}

class 估式 extends 式BaseListener{
    ParseTreeProperty<Integer> 值集 = new ParseTreeProperty<Integer>();

    @Override
    public void exit句(式Parser.句Context ctx) { 
        int 句数 = ctx.getChildCount();
        for (int i = 0; i < 句数; i++) {
            值集.put(ctx, 值集.get(ctx.getChild(i)));
        }
    }

    @Override 
    public void exit算式(式Parser.算式Context ctx) { 
        if (ctx.getChildCount() == 3) {
            int 左数 = 值集.get(ctx.算式(0));
            int 右数 = 值集.get(ctx.算式(1));
            if (ctx.算符.getType() == 式Parser.C乘) {
                值集.put(ctx, 左数*右数);
            }else{
                值集.put(ctx, 左数+右数);
            }
        }
        else{
            值集.put(ctx, 值集.get(ctx.getChild(0)));
        }
    }

    public void visitTerminal(TerminalNode 结) {
        Token 词 = 结.getSymbol();
        if (词.getType() == 式Parser.Z整数) {
            值集.put(结, Integer.valueOf(词.getText()));
        }
    }
}

/**
 * 算式器
 */
public class 算式器 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        式Lexer 取词器 = new 式Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        式Parser 树句器 = new 式Parser(词流);
        树句器.setBuildParseTree(true);
        ParseTree 树 = 树句器.句(); // 分析

        System.out.println(树.toStringTree());

        // 造树行者
        ParseTreeWalker 行者 = new ParseTreeWalker();

        算式 算 = new 算式();
        行者.walk(算, 树);
        System.out.println("算式结果："+算.叠.pop());  
        
        估式 估 = new 估式();
        行者.walk(估, 树);
        System.out.println("估式结果："+估.值集.get(树));
    }    
}