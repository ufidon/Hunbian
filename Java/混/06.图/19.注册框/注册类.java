/* 基本图形控件之二
*/
import java.util.*;
import javafx.application.*;
import javafx.collections.*;
import javafx.event.*;
import javafx.geometry.*;
import javafx.scene.*;
import javafx.scene.control.*;
import javafx.scene.input.*;
import javafx.scene.layout.*;
import javafx.scene.text.*;
import javafx.stage.*;

/**
 * 注册类
 */
public class 注册类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      Text 名签 = new Text("姓名"); TextField 名框 = new TextField();
      Text 诞签 = new Text("生日"); DatePicker 诞择 = new DatePicker();
      Text 性签 = new Text("性别");      
      ToggleGroup 性组 = new ToggleGroup();
      RadioButton 女钮 = new RadioButton("女");
      RadioButton 男钮 = new RadioButton("男");
      性组.getToggles().addAll(女钮,男钮);
      Text 订签 = new Text("预定");
      ToggleGroup 订组 = new ToggleGroup();
      ToggleButton 是钮 = new ToggleButton("是");
      ToggleButton 否钮 = new ToggleButton("否");
      订组.getToggles().addAll(是钮,否钮);
      Text 技签 = new Text("技能");
      CheckBox 会爪语 = new CheckBox("爪语"); 会爪语.setIndeterminate(false);
      CheckBox 会狗语 = new CheckBox("狗语"); 会狗语.setIndeterminate(false);

      Text 学位签 = new Text("学位");
      /*
      List<String> 学位名 = Arrays.asList(new String[]{
        "工科学士","理学学士","商科学士","工科硕士","理学硕士","商科硕士","工科博士","理学博士","商科博士"
      });
      ObservableList<String> 学位列表 = FXCollections.observableArrayList(学位名);
      */
      ObservableList<String> 学位列表 = FXCollections.observableArrayList(
        "工科学士","理学学士","商科学士","工科硕士","理学硕士","商科硕士","工科博士","理学博士","商科博士"
      );
      ListView<String> 学位列视 = new ListView<String>(学位列表);
      Text 址签 = new Text("地址");
      ChoiceBox<String> 地址列表 = new ChoiceBox<String>(FXCollections.observableArrayList(
        "北京","上海","深圳","广州","杭州","天津","南京","重庆","西安"
      ));
      Button 注钮 = new Button("注册");     

      GridPane 网布 = new GridPane();
      网布.setMinSize(600, 600);
      网布.setPadding(new Insets(10, 10, 10, 10));
      网布.setVgap(5); 网布.setHgap(5);
      网布.setAlignment(Pos.CENTER);

      网布.add(名签, 0, 0); 网布.add(名框, 1, 0);
      网布.add(诞签, 0, 1); 网布.add(诞择, 1, 1);
      网布.add(性签, 0, 2); 网布.add(女钮, 1, 2); 网布.add(男钮, 2, 2);
      网布.add(订签, 0, 3); 网布.add(是钮, 1, 3); 网布.add(否钮, 2, 3);
      网布.add(技签, 0, 4); 网布.add(会爪语, 1, 4); 网布.add(会狗语, 2, 4);
      网布.add(学位签, 0, 5); 网布.add(学位列视, 1, 5);
      网布.add(址签, 0, 6); 网布.add(地址列表, 1, 6);
      网布.add(注钮, 2, 8);

      // 风格设置
      注钮.setStyle("-fx-background-color: lightblue; -fx-textfill: white;");
      名签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      订签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      诞签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      性签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      技签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      学位签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      址签.setStyle("-fx-font: normal bold 15px 'Noto Sans CJK SC';");
      网布.setStyle("-fx-background-color: lightgray;"); 

      // 事件处理
    //   清钮.setOnMouseClicked(new EventHandler<MouseEvent>(){
    //       public void handle(MouseEvent 事) {
    //           邮框.clear(); 密框.clear();
    //       }
    //   });
      // 布景
      Scene 布景 = new Scene(网布);  
      
      // 舞台
      舞台.setTitle("基本控件之二: 注册框"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}