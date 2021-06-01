/* CSS 级联式样表
* 参考：
* 1. https://docs.oracle.com/javafx/2/api/javafx/scene/doc-files/cssref.html
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

import java.net.URL;

import com.sun.tools.javac.Main;

import javafx.animation.*;
import javafx.stage.*;
import javafx.util.*;
import javafx.css.*;

/**
 * 式样类
 */
public class 式样类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
     Rectangle 矩形 = new Rectangle(100, 100);
     矩形.setLayoutX(50); 矩形.setLayoutY(50);
     //矩形.getStyleClass().add("矩形");
     矩形.getStyleClass().add("jx");

      
      // 排布设置
      Group 组 = new Group(矩形);


      // 布景
      Scene 布景 = new Scene(组,300,300); 
      //布景.getStylesheets().add(Main.class.getResource("式样.css").toExternalForm());
      //布景.getStylesheets().add("file:///home/x/workspace/java/06.图/28.式样/式样.css");
      布景.getStylesheets().add("式样.css");
      
      // 舞台
      舞台.setTitle("级联式样表"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}