/* 图形控件横排
* 排成一行
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
 * 横排类
 */
public class 横排类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Text 登签 = new Text("今日口令");
      TextField 号框 = new TextField("在这里输入今天进入口令");
      Button 登钮 = new Button("进入");
      Button 清钮 = new Button("清空");
      
      // 排布设置
      HBox 横盒 = new HBox(10, 登签, 号框, 登钮,清钮);
      for (Node 子: 横盒.getChildren()) {
          HBox.setMargin(子, new Insets(20, 20, 20, 20));
      }
      横盒.setAlignment(Pos.CENTER);
      /*
      Node[] 节点集 = new Node[]{登签, 号框, 登钮,清钮};
      for (Node 节点 : 节点集) {
          HBox.setMargin(节点, new Insets(20, 20, 20, 20));
      }
      */
      
      // 布景
      Scene 布景 = new Scene(横盒);  
      
      // 舞台
      舞台.setTitle("控件横排成一行"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}