/* 图形控件锚定在锚盒内
* 跑：java -Dprism.forceGPU=true 锚排类
*/
import javafx.application.*;
import javafx.collections.*;
import javafx.event.*;
import javafx.geometry.*;
import javafx.scene.*;
import javafx.scene.control.*;
import javafx.scene.input.*;
import javafx.scene.layout.*;
import javafx.scene.paint.*;
import javafx.scene.text.*;
import javafx.scene.transform.Rotate;
import javafx.scene.image.*;
import javafx.scene.shape.*;
import javafx.animation.*;
import javafx.stage.*;
import javafx.util.*;

/**
 * 锚排类
 */
public class 锚排类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Sphere 球 = new Sphere(100);

      PhongMaterial 球质 = new PhongMaterial();
      球质.setBumpMap(new Image("file:///home/x/workspace/java/06.图/00.图像/凤.png"));
      球.setMaterial(球质);

      RotateTransition 转动 = new RotateTransition(Duration.seconds(1));
      转动.setNode(球);  转动.setAxis(Rotate.Y_AXIS); 转动.setByAngle(360);
      转动.setCycleCount(RotateTransition.INDEFINITE);
      转动.setAutoReverse(false);
      转动.play();

      
      // 排布设置
      AnchorPane 锚盒 = new AnchorPane();
      AnchorPane.setTopAnchor(球, 50.0);
      AnchorPane.setLeftAnchor(球, 50.0);
      AnchorPane.setRightAnchor(球, 50.0);
      AnchorPane.setBottomAnchor(球, 50.0);
      锚盒.getChildren().addAll(球);


      // 布景
      Scene 布景 = new Scene(锚盒);  
      
      // 舞台
      舞台.setTitle("控件锚定在锚盒内"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}