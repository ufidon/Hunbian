/*
java 查符 <测试.c
*/

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;

// 入阶:符入域之阶
class 入阶 extends 子CBaseListener{
    ParseTreeProperty<符域> 域群 = new ParseTreeProperty<符域>();
    void 存域(ParserRuleContext ctx, 符域 域){ 域群.put(ctx, 域); }
    void 定义变量(子CParser.型Context ctx, Token 量名){ // 定义变量
        int 词类 = ctx.start.getType();
        符.符型 型 = 查符.查型(词类);
        量符 量 = new 量符(量名.getText(), 型);
        现域.入域(量);
    }

    全域 本全域;
    符域 现域; // 入符于现域

    @Override 
    public void enter档(子CParser.档Context ctx) { 
        本全域 = new 全域(null);
        现域 = 本全域;
    }
    @Override 
    public void exit档(子CParser.档Context ctx) { 
        System.out.println(本全域);
    }
    @Override 
    public void enter函宣(子CParser.函宣Context ctx) { 
        String 函名 = ctx.M名().getText();
        int 词类 = ctx.型().start.getType();
        符.符型 函型 = 查符.查型(词类);

        函符 本函符 = new 函符(函名, 函型, 现域);
        现域.入域(本函符);
        存域(ctx, 本函符); // 函符实现了域，亦域
        现域 = 本函符;
    }
    @Override 
    public void exit函宣(子CParser.函宣Context ctx) { 
        System.out.println(现域);
        现域 = 现域.父域();
    }
    @Override 
    public void enter块(子CParser.块Context ctx) { 
        现域 = new 局域(现域);
        存域(ctx, 现域);
    }
    @Override 
    public void exit块(子CParser.块Context ctx) { 
        System.out.println(现域);
        现域 = 现域.父域();
    }
    @Override 
    public void exit形参(子CParser.形参Context ctx) { 
        定义变量(ctx.型(), ctx.M名().getSymbol());
    }
    @Override 
    public void exit量宣(子CParser.量宣Context ctx) { 
        定义变量(ctx.型(), ctx.M名().getSymbol());
    }    
}

// 引阶：引符之阶
class 引阶 extends 子CBaseListener{
    ParseTreeProperty<符域> 域集;
    全域 本全域;
    符域 现域;

    public 引阶(全域 某全域, ParseTreeProperty<符域> 某域集){
        本全域 = 某全域; 域集 = 某域集;
    }
    @Override 
    public void enter档(子CParser.档Context ctx) { 
        现域 = 本全域;
    }
    @Override 
    public void enter函宣(子CParser.函宣Context ctx) { 
        现域 = 域集.get(ctx);
    }
    @Override 
    public void exit函宣(子CParser.函宣Context ctx) { 
        现域 = 现域.父域();
    }
    @Override 
    public void enter块(子CParser.块Context ctx) { 
        现域 = 域集.get(ctx);
    }
    @Override 
    public void exit块(子CParser.块Context ctx) { 
        现域 = 现域.父域();
    }
    @Override 
    public void exit变量式(子CParser.变量式Context ctx) { 
        String 量名 = ctx.M名().getSymbol().getText();
        符 录 = 现域.查(量名);
        if (录 == null) {
            查符.报错(ctx.M名().getSymbol(), 量名+"不存在。");
        }
        if (录 instanceof 函符) {
            查符.报错(ctx.M名().getSymbol(), 量名+"不是变量，乃函式。");
        }
    }
    @Override 
    public void exit调函式(子CParser.调函式Context ctx) { 
        String 函名 = ctx.M名().getSymbol().getText();
        符 录 = 现域.查(函名);
        if (录 == null) {
            查符.报错(ctx.M名().getSymbol(), 函名+"不存在。");
        }
        if (录 instanceof 量符) {
            查符.报错(ctx.M名().getSymbol(), 函名+"不是函式，乃变量。");
        }
    }
}

/**
 * 查符
 */
public class 查符 {
    public static 符.符型 查型(int 词型){
        switch (词型) {
            case 子CParser.X整型:
                return 符.符型.x整型;
            case 子CParser.X浮型:
                return 符.符型.x浮型;
            case 子CParser.X空型:
                return 符.符型.x空型;
            case 子CParser.X贰型:
                return 符.符型.x贰型;
            default:
                return 符.符型.x无效;
        }
    }
    public static void 报错(Token 词, String 信){
        System.err.printf("错'%s'出于第%d行第%d列\n", 信, 词.getLine(), 词.getCharPositionInLine());
    }
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
        入阶 入 = new 入阶();
        行者.walk(入, 树);
        引阶 引 = new 引阶(入.本全域, 入.域群);
        行者.walk(引, 树);
    }    
}