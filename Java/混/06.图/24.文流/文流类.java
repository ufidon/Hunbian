/* 文字流经节点
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
import javafx.scene.shape.*;
import javafx.stage.*;

/**
 * 文流类
 */
public class 文流类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Text[] 诗词 = {
         new Text("北斗七星高，哥舒夜带刀。至今窥牧马，不敢过临洮。"),
         new Text("葡萄美酒夜光杯，欲饮琵琶马上催。醉卧沙场君莫笑，古来征战几人回？"),
         new Text("驿外断桥边，寂寞开无主。已是黄昏独自愁，更着风和雨。无意苦争春，一任群芳妒。零落成泥碾作尘，只有香如故。"),
         new Text("大江东去，浪淘尽，千古风流人物。故垒西边，人道是：三国周郎赤壁。乱石崩云，惊涛裂岸，卷起千堆雪。江山如画，一时多少豪杰。遥想公瑾当年，小乔初嫁了，雄姿英发。羽扇纶巾，谈笑间、强虏灰飞烟灭。故国神游，多情应笑我，早生华发。人间如梦，一尊还酹江月。")
      };

      诗词[0].setFill(Color.RED); 诗词[0].setFont(Font.font(30));
      诗词[1].setFill(Color.GREEN); 诗词[1].setFont(Font.font(16));
      诗词[2].setFill(Color.ORANGE); 诗词[2].setFont(Font.font(20));
      诗词[3].setFill(Color.SEAGREEN); 诗词[3].setFont(Font.font(24));
      
      // 排布设置
      TextFlow 文流盒 = new TextFlow(诗词);
      文流盒.setTextAlignment(TextAlignment.LEFT);
      文流盒.setPrefSize(500, 300); 文流盒.setLineSpacing(5);


      // 布景
      Scene 布景 = new Scene(文流盒);  
      
      // 舞台
      舞台.setTitle("文字流经节点"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}