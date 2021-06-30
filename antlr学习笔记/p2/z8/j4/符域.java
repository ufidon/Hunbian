public interface 符域 {
    public String 域名();
    public 符域 父域();
    /** 符入域 */
    public void 入域(符 某符);
    /* 此域查名，无果查父域 */
    public 符 查(String 名);
}
