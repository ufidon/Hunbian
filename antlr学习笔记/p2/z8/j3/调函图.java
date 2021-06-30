/*
java 调函图 <测试.c
*/
import java.util.*;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;

class 图{
    // 函为结调为边
    Set<String> 结集 = new OrderedHashSet<String>();
    MultiMap<String,String> 边集 = new MultiMap<String,String>();

    public void 连边(String 起点, String 终点){
        边集.map(起点, 终点);
    }
    public String 转串(){
        return "边集："+边集.toString()+"\n函集："+结集;
    }
    public String 转点图(){
        StringBuilder 区 = new StringBuilder();
        区.append("\ndigraph 矢图{\n");
        区.append(" //图形配置\n");
        区.append(" ranksep=.2;\n");
        区.append(" edge [arrowsize=.4]\n");
        区.append(" node [shape=circle]\n");
        区.append(" ");
        区.append("\n //结集，函数集\n ");
        for(String 结: 结集){
            区.append(结);
            区.append("; ");            
        }
        区.append("\n");
        区.append("\n //边集，函数调用\n");
        for(String 起点 : 边集.keySet()){
            for(String 终点:边集.get(起点)){
                区.append(" ");
                区.append(起点);
                区.append(" -> ");
                区.append(终点);
                区.append(";\n");
            }
        }
        区.append("}\n");
        return 区.toString();
    }
}

class 听函者 extends 子CBaseListener{
   图 函图 = new 图();
   String 现函名 = null;

   @Override 
   public void enter函宣(子CParser.函宣Context ctx) { 
       现函名 = ctx.M名().getText();
       函图.结集.add(现函名);

   }
   @Override 
   public void exit调函式(子CParser.调函式Context ctx) { 
       String 函名 = ctx.M名().getText();
       函图.连边(现函名, 函名);
   }
}


/**
 * 调函图
 */
public class 调函图 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        子CLexer 取词器 = new 子CLexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        子CParser 树句器 = new 子CParser(词流);
        树句器.setBuildParseTree(true);
        ParseTree 树 = 树句器.档(); // 造树

        //System.out.println(树.toStringTree());

        ParseTreeWalker 行者 = new ParseTreeWalker();
        听函者 听 = new 听函者();
        行者.walk(听, 树);
        System.out.println(听.函图.转串());  
        System.out.println(听.函图.转点图());   
    }    
}