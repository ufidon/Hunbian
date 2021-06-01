/* 多个变形演示
* 跑：java -Dprism.forceGPU=true 变形类
* 三维问题：
* https://stackoverflow.com/questions/30288837/warning-system-cant-support-conditionalfeature-scene3d-vmware-ubuntu
*/
import javafx.application.Application; 
import javafx.scene.Group; 
import javafx.scene.Scene; 
import javafx.scene.paint.Color;
import javafx.scene.shape.Box;
import javafx.scene.shape.Rectangle; 
import javafx.scene.transform.Rotate; 
import javafx.scene.transform.Scale; 
import javafx.scene.transform.Translate; 
import javafx.stage.Stage; 


public class 变形类 extends Application { 
   @Override 
   public void start(Stage 舞台) { 
      // 景物
      Rectangle 矩形 = new Rectangle(50, 50, 100, 75); 
      矩形.setFill(Color.GREEN); 
      矩形.setStroke(Color.RED); 

      Box 盒 = new Box(150, 160, 170);

      // 矩形依次变形 
      // 转动 
      Rotate 转动 = new Rotate(30, 160, 230);
      // 伸缩
      Scale 伸缩 = new Scale(1.6, 1.2, 300, 140);
      // 平移
      Translate 平移 = new Translate(300, 100); 
       
      // 对矩形进行上述变换
      矩形.getTransforms().addAll(转动, 伸缩, 平移); 

      // 盒依次变形
      Translate 平移2 = new Translate(400, 200, 30);
      Rotate 绕x转 = new Rotate(20, 0, 0, 0, Rotate.X_AXIS);
      Rotate 绕y转 = new Rotate(30, 10, 20, 30, Rotate.Y_AXIS);
      Rotate 绕z转 = new Rotate(40, 30, 20, 10, Rotate.Z_AXIS);
      盒.getTransforms().addAll(平移2, 绕z转, 绕y转, 绕x转);
        

      Group 组 = new Group(矩形, 盒); 
      
      // 布景
      Scene 布景 = new Scene(组, 800, 600);  
      
      // 舞台
      舞台.setTitle("依次变形"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   } 
}