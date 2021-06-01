/* 图形控件重复铺排
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
 * 铺排类
 */
public class 铺排类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      // 甲、乙、丙、丁、戊、己、庚、辛、壬、癸
     Text[] 天干= new Text[]{
      new Text("甲"),
      new Text("乙"),
      new Text("丙"),
      new Text("丁"),
      new Text("戊"),
      new Text("己"),
      new Text("庚"),
      new Text("辛"),
      new Text("壬"),
      new Text("癸")
     };
     for (Text 干 : 天干) {
        干.setFill(Color.GREEN); 干.setStroke(Color.YELLOW);
        干.setFont(Font.font(30));
     }
      
      // 排布设置
      TilePane 铺盒 = new TilePane(Orientation.VERTICAL, 10, 10, 天干);
      铺盒.setPrefRows(3);   


      // 布景
      Scene 布景 = new Scene(铺盒);  
      
      // 舞台
      舞台.setTitle("控件规则铺排"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}