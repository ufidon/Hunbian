/* 图形刷与渐变色
*/
import javafx.application.Application; 
import javafx.scene.Group; 
import javafx.scene.Scene; 
import javafx.scene.image.*; 
import javafx.scene.paint.*; 
import javafx.stage.Stage; 
import javafx.scene.shape.*; 
import javafx.scene.text.*; 

/**
 * 漆墨类
 */
public class 漆墨类 extends Application{
    @Override
    public void start(Stage 舞台) {
        // 景物
        // 鹰徽
        Circle 鹰徽 = new Circle(300, 300, 280);
        String 鹰档 = "file:///home/x/workspace/java/06.图/00.图像/小鹰.png";
        Image 鹰 = new Image(鹰档);
        ImagePattern 图章 = new ImagePattern(鹰, 5, 5, 80, 62, false);
        鹰徽.setFill(图章);
        鹰徽.setStroke(Color.BLUE);
        鹰徽.setStrokeWidth(10);

        // 色带
        Rectangle 色带 = new Rectangle(600, 10, 150, 580);
        Stop[] 色界 = new Stop[]{
            new Stop(0, Color.YELLOW),
            new Stop(1, Color.GREEN)
        };
        LinearGradient 渐变 = new LinearGradient(0, 0, 0, 1, true, CycleMethod.NO_CYCLE, 色界);
        色带.setFill(渐变);

        // 布景
        Group 景物 = new Group(鹰徽, 色带);
        Scene 布景 = new Scene(景物,800,600);

        // 舞台
        舞台.setTitle("爪语图架之画刷");
        舞台.setScene(布景);
        舞台.show();          
    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }
}
