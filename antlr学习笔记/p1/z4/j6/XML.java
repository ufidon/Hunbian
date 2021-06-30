// Generated from XML.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class XML extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		Q签头=1, Z注释=2, T特符=3, W文本=4, Q签尾=5, X斜签尾=6, D等于=7, C串=8, X斜名=9, M名=10, 
		K空格=11;
	public static final int
		Q签内=1;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE", "Q签内"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"Q签头", "Z注释", "T特符", "W文本", "Q签尾", "X斜签尾", "D等于", "C串", "X斜名", "M名", 
			"K空格", "H汉字", "Z字母", "S数字"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'<'", null, null, null, "'>'", "'/>'", "'='"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "Q\u0001\u0002", "Z\u0001\u0002", "T\u0001\u0002", "W\u0001\u0002", 
			"Q\u0001\u0002", "X\u0001\u0002\u0003", "D\u0001\u0002", "C\u0001", "X\u0001\u0002", 
			"M\u0001", "K\u0001\u0002"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public XML(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "XML.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\ro\b\1\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\3\2\3\2\3\2\3\2\3\3\3\3\3\3"+
		"\3\3\3\3\3\3\7\3+\n\3\f\3\16\3.\13\3\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4\6"+
		"\48\n\4\r\4\16\49\3\4\3\4\3\5\6\5?\n\5\r\5\16\5@\3\6\3\6\3\6\3\6\3\7\3"+
		"\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\7\tP\n\t\f\t\16\tS\13\t\3\t\3\t\3\n\3\n"+
		"\3\n\3\13\3\13\5\13\\\n\13\3\13\3\13\3\13\7\13a\n\13\f\13\16\13d\13\13"+
		"\3\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\4,Q\2\20\4\3\6\4\b\5\n\6"+
		"\f\7\16\b\20\t\22\n\24\13\26\f\30\r\32\2\34\2\36\2\4\2\3\b\3\2c|\4\2("+
		"(>>\5\2\13\f\17\17\"\"\3\2\u2e82\ua001\4\2C\\c|\3\2\62;\2r\2\4\3\2\2\2"+
		"\2\6\3\2\2\2\2\b\3\2\2\2\2\n\3\2\2\2\3\f\3\2\2\2\3\16\3\2\2\2\3\20\3\2"+
		"\2\2\3\22\3\2\2\2\3\24\3\2\2\2\3\26\3\2\2\2\3\30\3\2\2\2\4 \3\2\2\2\6"+
		"$\3\2\2\2\b\65\3\2\2\2\n>\3\2\2\2\fB\3\2\2\2\16F\3\2\2\2\20K\3\2\2\2\22"+
		"M\3\2\2\2\24V\3\2\2\2\26[\3\2\2\2\30e\3\2\2\2\32i\3\2\2\2\34k\3\2\2\2"+
		"\36m\3\2\2\2 !\7>\2\2!\"\3\2\2\2\"#\b\2\2\2#\5\3\2\2\2$%\7>\2\2%&\7#\2"+
		"\2&\'\7/\2\2\'(\7/\2\2(,\3\2\2\2)+\13\2\2\2*)\3\2\2\2+.\3\2\2\2,-\3\2"+
		"\2\2,*\3\2\2\2-/\3\2\2\2.,\3\2\2\2/\60\7/\2\2\60\61\7/\2\2\61\62\7@\2"+
		"\2\62\63\3\2\2\2\63\64\b\3\3\2\64\7\3\2\2\2\65\67\7(\2\2\668\t\2\2\2\67"+
		"\66\3\2\2\289\3\2\2\29\67\3\2\2\29:\3\2\2\2:;\3\2\2\2;<\7=\2\2<\t\3\2"+
		"\2\2=?\n\3\2\2>=\3\2\2\2?@\3\2\2\2@>\3\2\2\2@A\3\2\2\2A\13\3\2\2\2BC\7"+
		"@\2\2CD\3\2\2\2DE\b\6\4\2E\r\3\2\2\2FG\7\61\2\2GH\7@\2\2HI\3\2\2\2IJ\b"+
		"\7\4\2J\17\3\2\2\2KL\7?\2\2L\21\3\2\2\2MQ\7$\2\2NP\13\2\2\2ON\3\2\2\2"+
		"PS\3\2\2\2QR\3\2\2\2QO\3\2\2\2RT\3\2\2\2SQ\3\2\2\2TU\7$\2\2U\23\3\2\2"+
		"\2VW\7\61\2\2WX\5\26\13\2X\25\3\2\2\2Y\\\5\34\16\2Z\\\5\32\r\2[Y\3\2\2"+
		"\2[Z\3\2\2\2\\b\3\2\2\2]a\5\32\r\2^a\5\34\16\2_a\5\36\17\2`]\3\2\2\2`"+
		"^\3\2\2\2`_\3\2\2\2ad\3\2\2\2b`\3\2\2\2bc\3\2\2\2c\27\3\2\2\2db\3\2\2"+
		"\2ef\t\4\2\2fg\3\2\2\2gh\b\f\3\2h\31\3\2\2\2ij\t\5\2\2j\33\3\2\2\2kl\t"+
		"\6\2\2l\35\3\2\2\2mn\t\7\2\2n\37\3\2\2\2\13\2\3,9@Q[`b\5\7\3\2\b\2\2\6"+
		"\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}