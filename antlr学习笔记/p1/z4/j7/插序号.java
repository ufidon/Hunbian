/* 给爪语类定义插入一个成员变量——序列号
*/
import org.antlr.v4.runtime.*;

/**
 * 插序号
 */
public class 插序号 extends JavaBaseListener{
    TokenStreamRewriter 重写器;
    public 插序号(TokenStream 词流) {
        重写器 = new TokenStreamRewriter(词流);
    }

    @Override
    public void enterClassBody(JavaParser.ClassBodyContext ctx) {
        String 员量 = "\n\tpublic static final long 序列号 = 1L;";
        重写器.insertAfter(ctx.start, 员量);
    }    
}