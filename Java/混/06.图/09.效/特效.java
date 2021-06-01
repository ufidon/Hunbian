/* 爪语图架特效,滤镜
*/

import javafx.application.Application; 
import javafx.geometry.Insets; 
import javafx.geometry.Pos; 
import javafx.scene.Scene; 
import javafx.scene.control.Button; 
import javafx.scene.layout.GridPane;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.FontWeight;
import javafx.scene.text.Text; 
import javafx.scene.control.TextField;
import javafx.scene.effect.Blend;
import javafx.scene.effect.BlendMode;
import javafx.scene.effect.Bloom;
import javafx.scene.effect.BoxBlur;
import javafx.scene.effect.ColorAdjust;
import javafx.scene.effect.ColorInput;
import javafx.scene.effect.DropShadow;
import javafx.scene.effect.GaussianBlur;
import javafx.scene.effect.Glow;
import javafx.scene.effect.ImageInput;
import javafx.scene.effect.InnerShadow;
import javafx.scene.effect.Light;
import javafx.scene.effect.Lighting;
import javafx.scene.effect.MotionBlur;
import javafx.scene.effect.Reflection;
import javafx.scene.effect.SepiaTone;
import javafx.scene.effect.Shadow;
import javafx.scene.effect.Light.Spot;
import javafx.scene.image.Image;
import javafx.stage.Stage; 

/**
 * 特效
 */
public class 特效 extends Application{
    @Override
    public void start(Stage 舞台) {
        // 景物
        Text[][] 特效文 = new Text[3][6];
        String[][] 特效名={{"调色","入色","入图","混合","缀花","光晕"},
                          {"朦盒","斑朦","移朦","反射","调褐","阴影"},
                          {"外影","内影","布光","远光","点光","光点"}};
        Color[][] 字色 = {{Color.ALICEBLUE,Color.ANTIQUEWHITE, Color.AQUA, Color.AQUAMARINE, Color.GREEN,Color.VIOLET},
                        {Color.SADDLEBROWN,Color.ROYALBLUE, Color.BLACK, Color.YELLOWGREEN, Color.BLUE, Color.BROWN},
                        {Color.GOLDENROD, Color.CHOCOLATE, Color.CORAL, Color.CYAN, Color.RED, Color.YELLOW}};

        GridPane 网布 = new GridPane();
        网布.setMinSize(600, 300);
        网布.setPadding(new Insets(10, 10, 10, 10));
        网布.setHgap(50); 网布.setVgap(50);
        网布.setGridLinesVisible(true);
        网布.setAlignment(Pos.CENTER);

        for(int 行=0; 行<3; 行++){
            for(int 列=0; 列<6; 列++){
                特效文[行][列] = new Text(特效名[行][列]+String.format("\n（%d,%d）", 行,列));
                特效文[行][列].setFont(Font.font("Noto Sans CJK SC", FontWeight.BOLD, 30));
                特效文[行][列].setFill(字色[行][列]);
                网布.add(特效文[行][列], 列, 行);
            }
        }

        // 18种特效
        // "调色","入色","入图","混合","缀花","光晕"
        // （0，0）.调色:调节点色
        ColorAdjust 调色 = new ColorAdjust(0.08, 0.6, 0.8, 0.3);
        特效文[0][0].setEffect(调色);
        // （0，1）.色入：在节点上盖矩形色块
        ColorInput 色入 = new ColorInput(0, 0, 60, 40, Color.LIGHTCYAN);
        特效文[0][1].setEffect(色入);
        // （0，2）.入图：在节点上盖图
        ImageInput 入图 = new ImageInput(new Image("file:///home/x/workspace/java/06.图/00.图像/小鹰.png"), 0, 0);
        特效文[0][2].setEffect(入图);
        // (0,3).混合:节点覆盖部分混合
        Blend 混合=new Blend(BlendMode.MULTIPLY);
        混合.setTopInput(色入);
        特效文[0][3].setEffect(混合);
        // (0,4).缀花:节点发光
        Bloom 缀花 = new Bloom(0.2);
        特效文[0][4].setEffect(缀花);
        // (0,5).光晕:节点发光
        Glow 光晕 = new Glow(0.8);
        特效文[0][5].setEffect(光晕);
        // "朦盒","斑朦","移朦","反射","调褐","阴影"
        // (1,0)朦盒
        BoxBlur 朦盒 = new BoxBlur(5, 2, 1);
        特效文[1][0].setEffect(朦盒);
        // (1,1)斑朦
        GaussianBlur 斑朦 = new GaussianBlur(5);
        特效文[1][1].setEffect(斑朦);
        // (1,2)移朦
        MotionBlur 移朦 = new MotionBlur(30, 8);
        特效文[1][2].setEffect(移朦);
        // (1,3)反射
        Reflection 反射 = new Reflection(2, 0.4, 0.8, 0.1);
        特效文[1][3].setEffect(反射);
        // (1,4)调褐
        SepiaTone 调褐 = new SepiaTone(0.2);
        特效文[1][4].setEffect(调褐);
        // (1,5)阴影
        Shadow 阴影 = new Shadow(5, Color.GREY);
        特效文[1][5].setEffect(阴影);

        // "外影","内影","布光","远光","点光","光点"
        // (2,0)外影
        DropShadow 外影 = new DropShadow(5, 2, 2, Color.SEAGREEN);
        特效文[2][0].setEffect(外影);
        // (2,1)内影
        InnerShadow 内影 = new InnerShadow(5, 0, 0, Color.VIOLET);
        特效文[2][1].setEffect(内影);
        // (2,2)布光
        Lighting 布光 = new Lighting();
        特效文[2][2].setEffect(布光);
        // (2,3)远光
        Light.Distant 远光 = new Light.Distant(10, 10, Color.YELLOWGREEN);
        Lighting 布远光 = new Lighting();
        布远光.setLight(远光);
        特效文[2][3].setEffect(布远光);
        // (2,4)点光
        Light.Spot 点光 = new Light.Spot(10, 0, 10, 1, Color.RED);
        Lighting 布点光 = new Lighting();
        布点光.setLight(点光);
        特效文[2][4].setEffect(布点光);        
        // (2,5)光点        
        Light.Point 光点 = new Light.Point(30, 0, 10, Color.ORANGERED);
        Lighting 布光点 = new Lighting();
        布光点.setLight(光点);
        特效文[2][5].setEffect(布光点);        
        // 布景
        Scene 布景 = new Scene(网布);

        // 舞台
        舞台.setTitle("爪语图架图形特效");
        舞台.setScene(布景);
        舞台.show();
    }

    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}
