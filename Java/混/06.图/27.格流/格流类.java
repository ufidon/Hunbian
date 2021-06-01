/* 图形控件流水般排列
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
 * 格流类
 */
public class 格流类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      // 10天干：甲、乙、丙、丁、戊、己、庚、辛、壬、癸
      // 12地支：子、丑、寅、卯、辰、巳、午、未、申、酉、戌、亥
      String[] 天干 = new String[]{"甲","乙","丙","丁","戊","己","庚","辛","壬","癸"};
      String[] 地支 = new String[]{"子","丑","寅","卯","辰","巳","午","未","申","酉","戌","亥"};
      String[] 干支 = new String[60]; Text[] 干支签 = new Text[60];
      for (int i = 0; i < 干支.length; i++) {
         干支[i] = 天干[i%10]+地支[i%12];
         /*
         System.out.printf("%s ", 干支[i]);
         if (i%10 == 0) {
            System.out.printf("\n");
         }
         */        
         干支签[i] = new Text(干支[i]);
         干支签[i].setFill(Color.GREEN); 干支签[i].setStroke(Color.YELLOW);
         干支签[i].setFont(Font.font(30));         
      }

      
      // 排布设置
      FlowPane 流盒 = new FlowPane(Orientation.HORIZONTAL, 10, 20, 干支签);
      流盒.setPrefSize(700, 300);


      // 布景
      Scene 布景 = new Scene(流盒);  
      
      // 舞台
      舞台.setTitle("控件流水般排列"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}