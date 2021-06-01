/* 图像处理
*/
import java.io.*;
import javafx.application.*;
import javafx.scene.*;
import javafx.scene.image.*;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.Text;
import javafx.stage.*;

/**
 * 像类
 */
public class 像类 extends Application{
    @Override
    public void start(Stage 舞台) throws Exception{
        String[] 图标名 = {"处理前", "处理后"};
        Text[] 图标 = new Text[图标名.length];
        for (int i = 0; i < 图标.length; i++) {
            图标[i] = new Text(180+i*400, 350, 图标名[i]);
            图标[i].setFill(Color.GREEN);
            图标[i].setFont(Font.font(30));
        }
        Image 象 = new Image("file:///home/x/workspace/java/06.图/00.图像/大象.png");
        int 图宽 = (int)象.getWidth();
        int 图高 = (int)象.getHeight();
        PixelReader 读图 = 象.getPixelReader();
        WritableImage 象2 = new WritableImage(读图, 图宽, 图高);
        PixelWriter 写图 = 象2.getPixelWriter();

        // 像素处理
        for(int 行=0; 行<图宽; 行++){
            for(int 列=0; 列<图高; 列++){
                Color 像素色 = 读图.getColor(行, 列);
                Color 新色 = 像素色.invert();
                写图.setColor(行, 列, 新色);
            }
        }
        
        ImageView 象像1 = new ImageView(象);
        // ImageView 象像1 = new ImageView("file:///home/x/workspace/java/06.图/00.图像/大象.png");

        象像1.setX(0);象像1.setY(0);
        象像1.setFitWidth(400); 象像1.setFitHeight(300);
        象像1.setPreserveRatio(true);

        ImageView 象像2 = new ImageView(象2);
        象像2.setX(400);象像2.setY(0);
        象像2.setFitWidth(400); 象像2.setFitHeight(300);
        象像2.setPreserveRatio(true);

        // 布景
        Group 景物 = new Group(象像1,象像2);
        景物.getChildren().addAll(图标);
        Scene 布景 = new Scene(景物,800,400);

        // 舞台
        舞台.setTitle("图像处理");
        舞台.setScene(布景);
        舞台.show();        
    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }    
}