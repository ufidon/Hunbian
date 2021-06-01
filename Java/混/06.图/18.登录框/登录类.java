/* 基本图形控件之一
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
 * 登录类
 */
public class 登录类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Text 邮签 = new Text("电子邮件");
      Text 密签 = new Text("登录密码");

      TextField 邮框 = new TextField();
      PasswordField 密框 = new PasswordField();

      Button 登钮 = new Button("登录");
      Button 清钮 = new Button("清除");

      GridPane 网布 = new GridPane();
      网布.setMinSize(400, 200);
      网布.setPadding(new Insets(10, 10, 10, 10));
      网布.setHgap(5); 网布.setVgap(5);
      网布.setAlignment(Pos.CENTER);

      网布.add(邮签, 0, 0); 网布.add(邮框, 1, 0);
      网布.add(密签, 0, 1); 网布.add(密框, 1, 1);
      网布.add(登钮, 0, 2); 网布.add(清钮, 1, 2);
      邮签.setStyle("-fx-font: normal bold 20px 'Noto Serif CJK SC';"); 
      密签.setStyle("-fx-font: normal bold 20px 'Noto Serif CJK SC';");     
      登钮.setStyle("-fx-background-color: blue; -fx-text-fill: white;-fx-font: normal bold 20px 'Noto Serif CJK SC';");
      清钮.setStyle("-fx-background-color: blue; -fx-text-fill: white;-fx-font: normal bold 20px 'Noto Serif CJK SC';");
      邮框.setStyle("-fx-font: normal bold 20px 'Noto Serif CJK SC';"); 
      密框.setStyle("-fx-font: normal bold 20px 'Noto Serif CJK SC';"); 
      网布.setStyle("-fx-background-color: lightgreen;"); 

      // 事件处理
      清钮.setOnMouseClicked(new EventHandler<MouseEvent>(){
          public void handle(MouseEvent 事) {
              邮框.clear(); 密框.clear();
          }
      });
      // 布景
      Scene 布景 = new Scene(网布);  
      
      // 舞台
      舞台.setTitle("基本控件之一: 登录框"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}