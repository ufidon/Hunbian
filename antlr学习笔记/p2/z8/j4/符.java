public class 符 { // 函符、量符
    public static enum 符型 {x无效, x空型, x整型, x浮型, x贰型}

    String 名;      // 符皆有名
    符型 型;
    符域 域;      // 本符所在域

    public 符(String 名) { this.名 = 名; }
    public 符(String 名, 符型 型) { this.名 = 名; this.型 = 型; }

    public String 名() { return 名; }

    public String toString() {
        if ( 型!=符型.x无效 ) return '<'+名()+":"+型+'>';
        return 名();
    }
}
