/* 数据图
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
import javafx.scene.chart.*;

import java.util.ArrayList;

import javafx.animation.*;
import javafx.stage.*;
import javafx.util.*;
import javafx.css.*;
import java.math.*;

/**
 * 图表类
 */
public class 图表类 extends Application{
    @Override 
   public void start(Stage 舞台) { 
      // 景物
      String[] 图集 = new String[]{"饼图", "线图", "面图","柱图","泡图","点图","叠面图","叠柱图"};
      String[] 标题集 = new String[]{
         "4季度收入","4季度收入","4季度收入","4季度收入",
         "4季度收入","4季度收入","4季度收入","4季度收入"
      };
      String[] 季度 = new String[]{"春季","夏季","秋季","冬季"};
      int[] s季度 = {1,2,3,4};
      int[] 收入 = {100, 150, 200, 10};

      // 1. 饼图
      PieChart.Data[] _饼图数据 = new PieChart.Data[季度.length];
      for (int i = 0; i < _饼图数据.length; i++) {
         _饼图数据[i] = new PieChart.Data(季度[i], 收入[i]);
      }
      ObservableList<PieChart.Data> 饼图数据 = FXCollections.observableArrayList(
         /*
         new PieChart.Data("春季", 100),
         new PieChart.Data("夏季", 150),
         new PieChart.Data("秋季", 200),
         new PieChart.Data("冬季", 10)
         */
        _饼图数据);
      PieChart 饼图 = new PieChart(饼图数据);
      饼图.setTitle(图集[0]+": "+标题集[0]);
      饼图.setClockwise(true); 饼图.setLabelLineLength(10);
      饼图.setLabelsVisible(true); 饼图.setStartAngle(0);

      // 2. 线图
      NumberAxis 线图横轴 = new NumberAxis("季度", 1.0, 4.0, 1.0);
      NumberAxis 线图纵轴 = new NumberAxis("收入", 0.0, 200.0, 50.0);

      LineChart<Number,Number> 线图 = new LineChart<Number,Number>(线图横轴, 线图纵轴);
      线图.setTitle(图集[1]+": "+标题集[1]);

      XYChart.Series<Number,Number> 线图数据 = new XYChart.Series<Number,Number>();
      线图数据.setName(标题集[1]);
      for (int i = 0; i < 季度.length; i++) {
         线图数据.getData().add(new XYChart.Data<Number,Number>(s季度[i], 收入[i]));
      }
      线图.getData().add(线图数据);

      // 3. 面图
      CategoryAxis 面图横轴 = new CategoryAxis();
      面图横轴.setCategories(FXCollections.observableArrayList(季度));
      面图横轴.setLabel("季度");
      NumberAxis 面图纵轴 = new NumberAxis("收入", 0.0, 200.0, 50.0);
      AreaChart<String,Number> 面图 = new AreaChart<String,Number>(面图横轴, 面图纵轴);
      XYChart.Series<String,Number> 面图数据 = new XYChart.Series<String,Number>();
      面图数据.setName(标题集[2]);
      for (int i = 0; i < 季度.length; i++) {
         面图数据.getData().add(new XYChart.Data<String,Number>(季度[i],收入[i]));
      }
      面图.getData().add(面图数据);
      面图.setTitle(图集[2]+": "+标题集[2]);

      // 4. 柱图
      CategoryAxis 柱图横轴 = 面图横轴;
      NumberAxis 柱图纵轴 = 面图纵轴;

      BarChart<String,Number> 柱图 = new BarChart<String,Number>(柱图横轴,柱图纵轴);
      柱图.setTitle(图集[3]+": "+标题集[3]);
      XYChart.Series<String,Number> 柱图数据 = 面图数据;
      柱图.getData().add(柱图数据);

      // 5.泡图
      NumberAxis 泡图横轴 = 线图横轴;
      NumberAxis 泡图纵轴  =线图纵轴;

      BubbleChart<Number, Number> 泡图 = new BubbleChart<Number, Number>(泡图横轴, 泡图纵轴);
      XYChart.Series<Number,Number> 泡图数据 = new XYChart.Series<Number,Number>();
      for (int i = 0; i < s季度.length; i++) {
         泡图数据.getData().add(new XYChart.Data<Number,Number>(s季度[i], 收入[i], (int)Math.log10(收入[i])));
      }
      泡图.getData().add(泡图数据);
      泡图.setTitle(图集[4]+": "+标题集[4]);

      // 6. 点图
      CategoryAxis 点图横轴 = 面图横轴; 
      NumberAxis 点图纵轴 = 线图纵轴;
      ScatterChart<String, Number> 点图 = new ScatterChart<String, Number>(点图横轴, 点图纵轴);
      点图.getData().add(面图数据);
      点图.setTitle(图集[5]+": "+标题集[5]);

      // 7. 叠面图
      CategoryAxis 叠面图横轴 = 面图横轴; 
      NumberAxis 叠面图纵轴 = 线图纵轴;
      StackedAreaChart<String, Number> 叠面图 = new StackedAreaChart<String, Number>(叠面图横轴, 叠面图纵轴);
      XYChart.Series<String,Number> 叠面图数据 = new XYChart.Series<String,Number>();
      叠面图数据 = 面图数据;
      叠面图.getData().add(叠面图数据);
      叠面图.setTitle(图集[6]+": "+标题集[6]);      

      // 8. 叠柱图
      CategoryAxis 叠柱图横轴 = 面图横轴; 
      NumberAxis 叠柱图纵轴 = 线图纵轴;
      ScatterChart<String, Number> 叠柱图 = new ScatterChart<String, Number>(叠柱图横轴, 叠柱图纵轴);
      叠柱图.getData().add(面图数据);
      叠柱图.setTitle(图集[7]+": "+标题集[7]);

      // 排布设置
      GridPane 网布 = new GridPane();
      网布.setMinSize(1000, 600);
      网布.setPadding(new Insets(10, 10, 10, 10));
      网布.setVgap(5); 网布.setHgap(5);
      网布.setAlignment(Pos.CENTER);

      网布.add(饼图, 0, 1); 网布.add(线图,1, 1); 网布.add(面图,2,1); 网布.add(柱图,3,1);
      网布.add(泡图, 0, 2); 网布.add(点图,1, 2);网布.add(叠面图,2,2); 网布.add(叠柱图,3,2); 
      // 布景
      Scene 布景 = new Scene(网布); 

      
      // 舞台
      舞台.setTitle("数据图表"); 
      舞台.setScene(布景); 
      舞台.show(); 
   }      
   public static void main(String 行参集[]){ 
      launch(行参集); 
   }    
}