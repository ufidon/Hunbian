/* 爪语图架文本显示
*/

import javafx.application.Application;
import javafx.geometry.VPos;
import javafx.scene.Group; 
import javafx.scene.Scene;
import javafx.scene.effect.Effect;
import javafx.scene.paint.Color;
import javafx.stage.Stage;
import javafx.scene.text.Font;
import javafx.scene.text.FontPosture;
import javafx.scene.text.FontWeight;
import javafx.scene.text.Text; 
import javafx.scene.shape.*;

/**
 * 文
 */
public class 文 extends Application{
    private double 半径 = 200.0;
    private int 舞台宽 = 500;
    private int 舞台高 = 500;

    @Override
    public void start(Stage 舞台) {
        // 布景
        String[] 方位值 = {"东","东南","南","西南","西","西北","北","东北"};
        Text[] 方位 = new Text[方位值.length];
        Circle 圆 = new Circle(舞台宽/2, 舞台高/2, 半径);
        圆.setFill(Color.RED);
        圆.setStroke(Color.YELLOW);
        圆.setStrokeWidth(6.0);

        double 角, x, y;
        for (int i = 0; i < 方位.length; i++) {
            角 = i*Math.PI/4.0;            
            x = 舞台宽/2+半径*Math.cos(角);
            y = 舞台高/2+半径*Math.sin(角);

            方位[i] = new Text(方位值[i]);

            if (Math.cos(角)>0) {
                方位[i].setTextOrigin(VPos.BASELINE);
            } else if (Math.cos(角) ==0){
                方位[i].setTextOrigin(VPos.BOTTOM);
            }
            方位[i].setX(x);
            方位[i].setY(y);

            方位[i].setFill(Color.WHITESMOKE);
            方位[i].setStrokeWidth(2);
            方位[i].setStroke(Color.VIOLET);
            方位[i].setFont(Font.font("Noto Sans CJK SC", FontWeight.BOLD, FontPosture.REGULAR,36));

            if (i%2 == 0) {
                方位[i].setUnderline(true);
            } else {
                方位[i].setStrikethrough(true);
            }

        }
        
        Group 组 = new Group(圆);
        组.getChildren().addAll(方位);

        Scene 布景 = new Scene(组, 舞台宽,舞台高);
        布景.setFill(Color.SKYBLUE);

        舞台.setTitle("爪语图架文本显示");
        舞台.setScene(布景);


        舞台.show();

    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }    
}