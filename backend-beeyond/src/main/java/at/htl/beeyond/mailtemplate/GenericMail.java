package at.htl.beeyond.mailtemplate;

public class GenericMail {
    public String preheader;
    public String title;
    public String text;

    public GenericMail(String preheader, String title, String text) {
        this.preheader = preheader;
        this.title = title;
        this.text = text;
    }
}
