/*
java 算式器 <<<$(echo -e "2+3*2\n 1+2*3")
*/

import java.util.Map;
import java.util.Stack;

import org.antlr.v4.misc.OrderedHashMap;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;


/**
 * 算式
 */
class 算式 extends 签式BaseVisitor<Integer> {
    public Integer visit整数(签式Parser.整数Context ctx){
        return Integer.valueOf(ctx.Z整数().getText());
    }

    public Integer visit乘法(签式Parser.乘法Context ctx){
        return visit(ctx.算式(0))*visit(ctx.算式(1));
    }

    public Integer visit加法(签式Parser.加法Context ctx){
        return visit(ctx.算式(0)) + visit(ctx.算式(1));
    }
}

class 估式 extends 签式BaseListener{
    // 沿路收集与计算
    Stack<Integer> 叠 = new Stack<Integer>();

    @Override 
    public void exit乘法(签式Parser.乘法Context ctx) { 
        int 左数 = 叠.pop();
        int 右数 = 叠.pop();
        叠.push(左数*右数);
    }

    @Override 
    public void exit加法(签式Parser.加法Context ctx) { 
        int 左数 = 叠.pop();
        int 右数 = 叠.pop();
        叠.push(左数+右数);        
    }

    @Override 
    public void exit整数(签式Parser.整数Context ctx) { 
        叠.push(Integer.valueOf(ctx.Z整数().getText()));
    }
}

class 注式 extends 签式BaseListener{
    // 沿路收集与计算
    ParseTreeProperty<Integer> 注 = new ParseTreeProperty<Integer>();
    Map<String,Integer> 句值集 = new OrderedHashMap<String, Integer>();

    public void 设值(ParseTree 结, int 值){注.put(结, 值);}
    public int 取值(ParseTree 结){return 注.get(结);}

    @Override 
    public void exit句(签式Parser.句Context ctx) {
        for (int i=0; i<ctx.算式().size(); ++i) {
            设值(ctx, 取值(ctx.算式(i)));
            句值集.put(ctx.算式(i).toStringTree(), 取值(ctx.算式(i)));
        }         
    }

    @Override 
    public void exit乘法(签式Parser.乘法Context ctx) { 
        int 左数 = 取值(ctx.算式(0));
        int 右数 = 取值(ctx.算式(1));
        设值(ctx, 左数*右数);
    }

    @Override 
    public void exit加法(签式Parser.加法Context ctx) { 
        int 左数 = 取值(ctx.算式(0));
        int 右数 = 取值(ctx.算式(1));
        设值(ctx, 左数+右数);        
    }

    @Override 
    public void exit整数(签式Parser.整数Context ctx) { 
        设值(ctx, Integer.valueOf(ctx.Z整数().getText()));
    }
}


/**
 * 算式器
 */
public class 算式器 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        签式Lexer 取词器 = new 签式Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        签式Parser 树句器 = new 签式Parser(词流);
        树句器.setBuildParseTree(true);
        ParseTree 树 = 树句器.句(); // 造树

        System.out.println(树.toStringTree());

        // 都只返了最后一句，如何返全句?
        算式 算 = new 算式();
        int 结果 = 算.visit(树);        
        System.out.println("算式结果："+结果);  
        
        ParseTreeWalker 行者 = new ParseTreeWalker();
        估式 估 = new 估式();
        行者.walk(估, 树);
        System.out.println("估式结果："+估.叠.pop());

        注式 注 = new 注式();
        行者.walk(注, 树);
        System.out.println("注式结果："+注.取值(树)); 
        System.out.println(注.句值集);     
    }    
}