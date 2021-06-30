/** 从爪语类生成接口 */
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.TokensStartState;
import org.antlr.v4.runtime.misc.*;

/**
 * 取接口
 */
public class 取接口 extends Java9BaseListener{
    Java9Parser 爪9树句器;

    public 取接口(Java9Parser 树句器){爪9树句器 = 树句器;}

/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override 
    public void enterNormalClassDeclaration(Java9Parser.NormalClassDeclarationContext ctx) { 
        System.out.println("interface I"+ctx.identifier().getText()+ " {");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override 
    public void exitNormalClassDeclaration(Java9Parser.NormalClassDeclarationContext ctx) { 
        System.out.println("}");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void enterMethodDeclaration(Java9Parser.MethodDeclarationContext ctx) { 
        System.out.print("\t");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void exitMethodDeclaration(Java9Parser.MethodDeclarationContext ctx) { 
        System.out.print(";\n");
    }
    
    /**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void enterMethodModifier(Java9Parser.MethodModifierContext ctx) { 
        System.out.print(ctx.getText());
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void exitMethodModifier(Java9Parser.MethodModifierContext ctx) { 
        System.out.print(" ");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void enterMethodHeader(Java9Parser.MethodHeaderContext ctx) { 
        System.out.print(ctx.result().getText()+" ");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void exitMethodHeader(Java9Parser.MethodHeaderContext ctx) { 
        System.out.print(" ");
    }
    /**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void enterMethodDeclarator(Java9Parser.MethodDeclaratorContext ctx) { 
        System.out.print(ctx.identifier().getText() +"( ");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void exitMethodDeclarator(Java9Parser.MethodDeclaratorContext ctx) { 
        System.out.print(" )");
    }
    /**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void enterFormalParameters(Java9Parser.FormalParametersContext ctx) {
        System.out.print("形参列表，待做");
    }
	/**
	 * {@inheritDoc}
	 *
	 * <p>The default implementation does nothing.</p>
	 */
	@Override public void exitFormalParameters(Java9Parser.FormalParametersContext ctx) { }
	
	
    
}