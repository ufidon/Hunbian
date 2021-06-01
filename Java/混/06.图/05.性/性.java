/* 爪语图架视件属性变化响应
*/
import javafx.application.*;
import javafx.beans.property.DoubleProperty;
import javafx.beans.property.ReadOnlyDoubleProperty;
import javafx.beans.value.ChangeListener;
import javafx.beans.value.ObservableValue;
import javafx.scene.*;
import javafx.scene.layout.Pane;
import javafx.scene.paint.*;
import javafx.stage.*;

public class 性 extends Application{
    @Override
    public void start(Stage 舞台) throws Exception {
        // 造格
        Pane 格 = new Pane();
        ReadOnlyDoubleProperty 格宽 = 格.widthProperty();
        格宽.addListener(new ChangeListener<Number>(){
            @Override
            public void changed(
                ObservableValue<? extends Number> 变化值,
                Number 旧值, Number 新值) {
                System.out.printf("格宽从%.2f变成%.2f\n", 旧值,新值);
            }
        });

        DoubleProperty 喜格宽 = 格.prefWidthProperty();
        喜格宽.addListener((ObservableValue<? extends Number>喜,
        Number 旧值, Number 新值)->{
            System.out.printf("喜好的格宽从%.2f变成%.2f\n", 旧值,新值);
        });

        喜格宽.set(200); // 改变喜格宽

        // 造景
        Scene 景 = new Scene(格, 800,600, true);
        景.setFill(Color.BISQUE);

        // 舞台设置
        舞台.setTitle("爪语图形框架JavaFX属性变化响应");
        舞台.setScene(景);

        // 开幕
        舞台.show();
    }

    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}