/** 从爪语类生成接口 */
import org.antlr.v4.runtime.*;


/**
 * 取接口
 */
public class 取接口 extends JavaBaseListener{
    JavaParser 树句器;

    public 取接口(JavaParser 树句器){this.树句器 = 树句器;}
	/** 侦听 classDeclaration匹配 */
	@Override
	public void enterClassDeclaration(JavaParser.ClassDeclarationContext ctx){
		System.out.println("interface I"+ctx.Identifier()+" {");
	}
	@Override
	public void exitClassDeclaration(JavaParser.ClassDeclarationContext ctx) {
		System.out.println("}");
	}

	/** 侦听 methodDeclaration 匹配*/
	@Override
	public void enterMethodDeclaration(
		JavaParser.MethodDeclarationContext ctx
	)
	{
		// 用树句器取词流
		TokenStream 词流 = 树句器.getTokenStream();
		String 返型 = "void";
		if ( ctx.type()!=null ) {
			返型 = 词流.getText(ctx.type());
		}
		String 形参集 = 词流.getText(ctx.formalParameters());
		System.out.println("\t"+返型+" "+ctx.Identifier()+形参集+";");
	}   
}