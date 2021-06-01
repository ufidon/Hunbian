/* 图形控件叠在一起
* 跑：java -Dprism.forceGPU=true 叠排类
*/
import javafx.application.*;
import javafx.collections.*;
import javafx.event.*;
import javafx.geometry.*;
import javafx.scene.*;
import javafx.scene.control.*;
import javafx.scene.input.*;
import javafx.scene.layout.*;
import javafx.scene.paint.Color;
import javafx.scene.paint.PhongMaterial;
import javafx.scene.text.*;
import javafx.scene.shape.*;
import javafx.stage.*;

/**
 * 叠排类
 */
public class 叠排类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Text 签 = new Text("从前至后：签⇒球⇒圆");
      签.setFont(Font.font(20)); 签.setFill(Color.BLUE); 
      签.setX(15); 签.setY(20);

      Circle 圆= new Circle(300, 120, 100);
      圆.setFill(Color.RED); 圆.setStroke(Color.YELLOW); 圆.setStrokeWidth(10);

      Sphere 球 = new Sphere(50);
      球.setMaterial(new PhongMaterial(Color.GREEN));
      
      // 排布设置
      StackPane 叠盒 = new StackPane();
      StackPane.setMargin(圆, new Insets(50, 50, 50, 50));
      叠盒.getChildren().addAll(圆,球,签);


      // 布景
      Scene 布景 = new Scene(叠盒);  
      
      // 舞台
      舞台.setTitle("控件叠放在一起"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}