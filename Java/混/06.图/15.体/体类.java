/* 形体
* 跑：java -Dprism.forceGPU=true 体类
*/
import java.io.*;
import javafx.application.*;
import javafx.scene.*;
import javafx.scene.transform.*;
import javafx.scene.shape.*;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.Text;
import javafx.stage.*;

/**
 * 体类
 */
public class 体类 extends Application{
    @Override
    public void start(Stage 舞台) throws Exception{
      
        // 景物
        Sphere 球 = new Sphere(100);
        球.setTranslateX(110);
        球.setTranslateY(150);
        球.setCullFace(CullFace.BACK);

        Box 盒 = new Box(120, 130, 150);
        盒.setTranslateX(400);
        盒.setTranslateY(150);
        盒.setCullFace(CullFace.BACK);
        Rotate 绕x转 = new Rotate(20, 0, 0, 0, Rotate.X_AXIS);
        Rotate 绕y转 = new Rotate(30, 10, 20, 30, Rotate.Y_AXIS);
        Rotate 绕z转 = new Rotate(40, 30, 20, 10, Rotate.Z_AXIS);
        盒.getTransforms().addAll(绕z转, 绕y转, 绕x转);        

        Cylinder 柱 = new Cylinder(100, 200, 10);
        柱.setTranslateX(700);
        柱.setTranslateY(150);
        柱.setCullFace(CullFace.BACK);
        柱.setDrawMode(DrawMode.LINE);

        // 布景
        Group 景物 = new Group(球,盒,柱);
        Scene 布景 = new Scene(景物,800,400);

        // 舞台
        舞台.setTitle("形体");
        舞台.setScene(布景);
        舞台.show();        
    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }    
}