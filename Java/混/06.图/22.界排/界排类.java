/* 图形控件排在容器边界：上下左右中
*/
import javafx.application.*;
import javafx.event.EventHandler;
import javafx.geometry.*;
import javafx.scene.*;
import javafx.scene.control.*;
import javafx.scene.input.MouseEvent;
import javafx.scene.layout.*;
import javafx.scene.text.*;
import javafx.stage.*;

/**
 * 界排类
 */
public class 界排类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Text[] 界签 = {new Text("上"), new Text("下"), 
      new Text("左"), new Text("右"),new Text("中")};
      
      // 排布设置
      BorderPane 界盒 = new BorderPane();
      界盒.setTop(界签[0]); 界盒.setBottom(界签[1]); 界盒.setLeft(界签[2]);
      界盒.setRight(界签[3]); 界盒.setCenter(界签[4]);
      for (Node 子: 界盒.getChildren()) {
          BorderPane.setMargin(子, new Insets(20, 20, 20, 20));
      }
      BorderPane.setAlignment(界签[0], Pos.TOP_CENTER);
      BorderPane.setAlignment(界签[1], Pos.BOTTOM_CENTER);
      BorderPane.setAlignment(界签[2], Pos.CENTER_LEFT);
      BorderPane.setAlignment(界签[3], Pos.CENTER_RIGHT);


      // 布景
      Scene 布景 = new Scene(界盒);  
      
      // 舞台
      舞台.setTitle("控件排至边界"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}