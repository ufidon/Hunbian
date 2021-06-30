/**
 * 印屏配置
*/
import java.util.Map;

import org.antlr.v4.misc.OrderedHashMap;
import org.antlr.v4.runtime.*;

/**
 * 印屏
 */
class 印屏配置 extends 配档Parser {
    Map<String,String> 属性集 = new OrderedHashMap<String,String>();

    public 印屏配置(TokenStream input) {
		super(input);
	}

    void 配置(Token 名, Token 值){
        属性集.put(名.getText(), 值.getText());
        for (String 键 : 属性集.keySet()) {
            System.out.println(键+"="+属性集.get(键));  
        }
        //System.out.println(名.getText()+"="+值.getText());
    }   
}

/**
 * 应用
 */
public class 应用 {
    public static void main(String[] 行参集) throws Exception {
        CharStream 码流 = CharStreams.fromStream(System.in);
        配档Lexer 取词器 = new 配档Lexer(码流);
        CommonTokenStream 词流 = new CommonTokenStream(取词器);
        配档Parser 树句器 = new 印屏配置(词流);
        树句器.档();    
    }
    
}