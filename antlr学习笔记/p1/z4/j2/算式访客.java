/**
 * 实现算式访客
*/
import java.util.*;

/**
 * 算式访客
 */
public class 算式访客 extends 算式BaseVisitor<Integer>{
    // 用于存储变量名与值
    Map<String, Integer> 内存 = new HashMap<String, Integer>();

    	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit程式(算式Parser.程式Context ctx) {
        return visitChildren(ctx); 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit印式(算式Parser.印式Context ctx) {
        Integer 值 = visit(ctx.式());
        System.out.println(值);
        return 0; 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit赋(算式Parser.赋Context ctx) {
        String 量名 = ctx.L量名().getText();
        int 值 = visit(ctx.式());
        内存.put(量名, 值);
        return 值; 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit空行(算式Parser.空行Context ctx) {
        return visitChildren(ctx); 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit整数(算式Parser.整数Context ctx) {
        return Integer.valueOf(ctx.Z整数().getText()); 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit加减(算式Parser.加减Context ctx) {
        int 左数 = visit(ctx.式(0));
        int 右数 = visit(ctx.式(1));
        if (ctx.算符.getType() == 算式Parser.J加) {
            return 左数+右数;
        }
        return 左数-右数; 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit乘除(算式Parser.乘除Context ctx) {
        int 左数 = visit(ctx.式(0));
        int 右数 = visit(ctx.式(1));
        if (ctx.算符.getType() == 算式Parser.C乘) {
            return 左数*右数;
        }
        return 左数/右数; 
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit量名(算式Parser.量名Context ctx) {
        String 量名 = ctx.L量名().getText();
        if (内存.containsKey(量名)) {
            return 内存.get(量名);
        }
        return 0; 
    }
    /**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit清仓(算式Parser.清仓Context ctx) {
        内存.clear();
        return 0; 
    }
	
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation returns the result of calling
	 * {@link #visitChildren} on {@code ctx}.</p>
	 */
	@Override 
    public Integer visit括式(算式Parser.括式Context ctx) {
        return visit(ctx.式()); 
    }
    
}