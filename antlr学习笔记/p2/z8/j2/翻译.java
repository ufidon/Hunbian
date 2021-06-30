/*
java 翻译 < 测试.json
*/

import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.tree.*;


class 转XML extends JSONBaseListener{
    ParseTreeProperty<String> xml = new ParseTreeProperty<String>();
    String 取xml(ParseTree ctx) { return xml.get(ctx);}
    void 设xml(ParseTree ctx, String s) {xml.put(ctx, s);}

    @Override 
    public void exit粒值(JSONParser.粒值Context ctx) { 
        设xml(ctx, ctx.getText());
    }
    @Override 
    public void exit串值(JSONParser.串值Context ctx) { 
        设xml(ctx, 去引号(ctx.getText()));
    }
    @Override 
    public void exit物值(JSONParser.物值Context ctx) { 
        设xml(ctx, 取xml(ctx.物()));
    }

    @Override 
    public void exit队值(JSONParser.队值Context ctx) { 
        设xml(ctx, 取xml(ctx.队()));
    }


    @Override 
    public void exit双(JSONParser.双Context ctx) { 
        String 签名 = 去引号(ctx.C串().getText());
        JSONParser.值Context 值 = ctx.值();

        String 签 = String.format("<%s>%s</%s>\n", 签名, 取xml(值), 签名);
        设xml(ctx, 签);
    }

    @Override 
    public void exit某物(JSONParser.某物Context ctx) { 
        StringBuffer 区 = new StringBuffer();
        区.append("\n");
        for (JSONParser.双Context 双 : ctx.双()) {
            区.append(取xml(双));
        }
        设xml(ctx, 区.toString());
    }

    @Override 
    public void exit空物(JSONParser.空物Context ctx) { 
        设xml(ctx, "");
    }

    @Override 
    public void exit某队(JSONParser.某队Context ctx) { 
        StringBuffer 区 = new StringBuffer();
        区.append("\n");
        for (JSONParser.值Context 值 : ctx.值()) {
            区.append("<元素>");
            区.append(取xml(值));
            区.append("</元素>");
            区.append("\n");
        }
        设xml(ctx, 区.toString());
    }

    @Override 
    public void exit空队(JSONParser.空队Context ctx) { 
        设xml(ctx, "");
    }

    @Override 
    public void exitJson(JSONParser.JsonContext ctx) { 
        设xml(ctx, 取xml(ctx.getChild(0)));
    }

    public String 去引号(String s) {
        if ( s==null || s.charAt(0)!='"' ) return s;
        return s.substring(1, s.length() - 1);
    }
}


/**
 * 翻译
 */
public class 翻译 {
    public static void main(String[] 行参集) throws Exception {
        /** 新版antlr4 */
        CharStream 码流 = CharStreams.fromStream(System.in);
        JSONLexer 取词器 = new JSONLexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        JSONParser 树句器 = new JSONParser(词流);
        树句器.setBuildParseTree(true);
        ParseTree 树 = 树句器.json(); // 造树

        //System.out.println(树.toStringTree());

        ParseTreeWalker 行者 = new ParseTreeWalker();
        转XML 转 = new 转XML();
        行者.walk(转, 树);
        System.out.println(转.取xml(树));     
    }    
}