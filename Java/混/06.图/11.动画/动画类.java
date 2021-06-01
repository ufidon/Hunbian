/* 动画：连续变换
*/
import javafx.animation.RotateTransition;
import javafx.animation.ScaleTransition;
import javafx.animation.TranslateTransition;
import javafx.application.Application; 
import static javafx.application.Application.launch; 
import javafx.scene.*; 
import javafx.scene.Scene; 
import javafx.scene.paint.Color; 
import javafx.stage.Stage; 
import javafx.util.Duration; 
import javafx.scene.shape.*;
import javafx.scene.transform.Scale;
import javafx.scene.layout.*;

/**
 * 动画类
 */
public class 动画类 extends Application {
    @Override
    public void start(Stage 舞台) {
        // 景物
        Polygon 三角形 = new Polygon(150,100,250,280, 50,280);
        三角形.setFill(Color.YELLOW);
        三角形.setStroke(Color.YELLOWGREEN);
        三角形.setStrokeWidth(10);

        Ellipse 椭圆 = new Ellipse(300, 500, 100, 60);
        椭圆.setFill(Color.WHITESMOKE);
        椭圆.setStroke(Color.BLUE);
        椭圆.setStrokeWidth(10);

        Circle 太阳 = new Circle(600, 100, 100);
        太阳.setFill(Color.RED);
        太阳.setStroke(Color.ORANGERED);
        太阳.setStrokeWidth(10);

        // 转动三角形
        RotateTransition 转动 = new RotateTransition(Duration.seconds(1), 三角形);
        转动.setByAngle(360);
        转动.setCycleCount(5);
        转动.setAutoReverse(false);
        转动.play();

        // 伸缩椭圆
        ScaleTransition 伸缩 = new ScaleTransition(Duration.seconds(1), 椭圆);
        伸缩.setByX(0.6); 伸缩.setByY(1.6);
        伸缩.setCycleCount(5); 伸缩.setAutoReverse(false);
        伸缩.play();

        // 平移太阳
        TranslateTransition 降 = new TranslateTransition(Duration.seconds(1), 太阳);
        降.setByY(600); 降.setCycleCount(5);
        降.setAutoReverse(false);
        降.play();

        // 布景
        Group 景物 = new Group(三角形,椭圆,太阳);
        Scene 布景 = new Scene(景物,800,800);

        // 舞台
        舞台.setTitle("爪语图架图形动画");
        舞台.setScene(布景);
        舞台.show();   
    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }
    
}