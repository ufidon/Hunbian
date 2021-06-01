/* 图形变形
*/
import javafx.application.Application;
import javafx.geometry.VPos;
import javafx.scene.Group; 
import javafx.scene.Scene;
import javafx.scene.layout.HBox;
import javafx.scene.paint.Color; 
import javafx.scene.shape.Rectangle; 
import javafx.scene.transform.Rotate; 
import javafx.scene.transform.Scale;
import javafx.scene.transform.Shear;
import javafx.scene.transform.Translate; 
import javafx.stage.Stage;
import javafx.scene.text.*;


/**
 * 变形类
 */
public class 变形类 extends Application{
    @Override
    public void start(Stage 舞台) {
        // 景物
        String[] 变换名 = {"旋转","伸缩","歪斜","平移"};
        Text[] 变换 = new Text[变换名.length];
        for (int i = 0; i < 变换.length; i++) {
            变换[i] = new Text(变换名[i]);
            变换[i].setFont(Font.font("Noto Sans CJK SC", FontWeight.BOLD, 80));
            变换[i].setY(100);
            变换[i].setFill(Color.BLUE);
        }

        // 图形变换："旋转","伸缩","歪斜","平移"
        // 旋转：
        Rotate 转 = new Rotate(30, 200, 100);
        变换[0].setX(100);
        变换[0].setTextOrigin(VPos.BASELINE);
        变换[0].getTransforms().add(转);
        // 伸缩：
        Scale 伸 = new Scale(0.6, 2);
        变换[1].getTransforms().add(伸);
        // 歪斜：
        Shear 歪 = new Shear(0.2, -0.2);
        变换[2].getTransforms().add(歪);
        // 平移：
        Translate 移 = new Translate(0, 80);
        变换[3].getTransforms().add(移);
        
        HBox 横盒 = new HBox(40, 变换);
        // 布景
        Scene 布景 = new Scene(横盒,800,200);

        // 舞台
        舞台.setTitle("爪语图架图形变换");
        舞台.setScene(布景);
        舞台.show();        
    }
    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}