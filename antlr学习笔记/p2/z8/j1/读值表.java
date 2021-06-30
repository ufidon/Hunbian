/*
java 读值表 < 数据.csv
*/

import java.util.*;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;


class 载值表 extends 值表BaseListener{
    public static final String 空 = "";
    List<Map<String, String>> 行集 = new ArrayList<Map<String, String>>();
    List<String> 题行 = new ArrayList<String>();
    List<String> 现行;

    @Override 
    public void exit串(值表Parser.串Context ctx) { 
        现行.add(ctx.C串().getText());
    }

    @Override 
    public void exit文本(值表Parser.文本Context ctx) { 
        现行.add(ctx.W文本().getText());
    }

    @Override 
    public void exit空(值表Parser.空Context ctx) { 
        现行.add(空);
    }

    @Override 
    public void exit标题(值表Parser.标题Context ctx) { 
        题行.addAll(现行);
    }

    @Override 
    public void enter行(值表Parser.行Context ctx) { 
        现行 = new ArrayList<String>();
    }
    @Override 
    public void exit行(值表Parser.行Context ctx) { 
        if (ctx.getParent().getRuleIndex() == 值表Parser.RULE_标题) {
            // 标题行已在exit标题中完成，此处勿重复
            return;
        }

        // 数据行
        Map<String,String> 行 = new LinkedHashMap<String,String>();
        int 列号 = 0;
        for (String s : 现行) {
            行.put(题行.get(列号), s);
            ++列号;
        }
        行集.add(行);
    }
}


/**
 * 读值表
 */
public class 读值表 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        值表Lexer 取词器 = new 值表Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        值表Parser 树句器 = new 值表Parser(词流);
        树句器.setBuildParseTree(true);
        ParseTree 树 = 树句器.档(); // 造树

        System.out.println(树.toStringTree());

        ParseTreeWalker 行者 = new ParseTreeWalker();
        载值表 载 = new 载值表();
        行者.walk(载, 树);
        System.out.println(载.行集);     
    }    
}