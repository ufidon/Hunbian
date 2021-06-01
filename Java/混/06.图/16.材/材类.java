/* 形体材质
* 跑：java -Dprism.forceGPU=true 材类
*/
import java.io.*;
import javafx.application.*;
import javafx.scene.*;
import javafx.scene.transform.*;
import javafx.scene.shape.*;
import javafx.scene.image.*;
import javafx.scene.paint.*;
import javafx.scene.text.Font;
import javafx.scene.text.Text;
import javafx.stage.*;

/**
 * 材类
 */
public class 材类 extends Application{
    @Override
    public void start(Stage 舞台) throws Exception{
      
        // 景物
        Sphere 球 = new Sphere(100);
        球.setTranslateX(120);
        球.setTranslateY(150);
        球.setCullFace(CullFace.BACK);
        /*
        PhongMaterial 球质 = new PhongMaterial(
            Color.GREENYELLOW, 
            new Image("file:///home/x/workspace/java/06.图/00.图像/大象.png"), 
            new Image("file:///home/x/workspace/java/06.图/00.图像/燕.png"),
            new Image("file:///home/x/workspace/java/06.图/00.图像/雀.png"),
            new Image("file:///home/x/workspace/java/06.图/00.图像/鸽.png")
        );
        */
        PhongMaterial 球质 = new PhongMaterial();
        球质.setBumpMap(new Image("file:///home/x/workspace/java/06.图/00.图像/大象.png"));
        球.setMaterial(球质);

        Box 盒 = new Box(120, 130, 150);
        盒.setTranslateX(400);
        盒.setTranslateY(150);
        盒.setCullFace(CullFace.BACK);
        Rotate 绕x转 = new Rotate(20, 0, 0, 0, Rotate.X_AXIS);
        Rotate 绕y转 = new Rotate(30, 10, 20, 30, Rotate.Y_AXIS);
        Rotate 绕z转 = new Rotate(40, 30, 20, 10, Rotate.Z_AXIS);
        盒.getTransforms().addAll(绕z转, 绕y转, 绕x转);
        PhongMaterial 盒质 = new PhongMaterial();
        盒质.setDiffuseMap(new Image("file:///home/x/workspace/java/06.图/00.图像/燕.png"));
        盒.setMaterial(盒质);       

        Cylinder 柱 = new Cylinder(100, 200, 10);
        柱.setTranslateX(660);
        柱.setTranslateY(150);
        柱.setCullFace(CullFace.BACK);
        柱.setDrawMode(DrawMode.FILL);
        PhongMaterial 柱质 = new PhongMaterial();
        柱质.setDiffuseMap(new Image("file:///home/x/workspace/java/06.图/00.图像/凤.png"));
        柱.setMaterial(柱质);

        // 布景
        Group 景物 = new Group(球,盒,柱);
        Scene 布景 = new Scene(景物,800,400);

        // 观察相机
        PerspectiveCamera 相机 = new PerspectiveCamera(false);
        相机.setTranslateX(0);
        相机.setTranslateY(0);
        相机.setTranslateZ(-10);
        布景.setCamera(相机);


        // 舞台
        舞台.setTitle("形体材质");
        舞台.setScene(布景);
        舞台.show();        
    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }    
}