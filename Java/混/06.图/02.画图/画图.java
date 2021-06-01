/* 用爪语图架画画
*/

import javax.swing.BoxLayout;

import javafx.application.*;
import javafx.scene.*;
import javafx.scene.layout.HBox;
import javafx.scene.paint.*;
import javafx.scene.shape.*;
import javafx.scene.text.Font;
import javafx.scene.text.Text;
import javafx.stage.*;

public class 画图 extends Application{
    @Override
    public void start(Stage 舞台) throws Exception {
        // 画图
        Line 线 = new Line();
        线.setStartX(20);
        线.setStartY(30);
        线.setEndX(500);
        线.setEndY(400);
        线.setStrokeWidth(5.0);
        线.setStrokeType(StrokeType.CENTERED);

        Text 文 = new Text(100, 100, "爪语画图框架");
        文.setFont(new Font("Noto Sans CJK SC", 36));

        // 路径
        Path 径 = new Path(
            new MoveTo(300,180),
            new LineTo(330,330),
            new LineTo(430,150),
            new LineTo(470,350),
            new LineTo(300,180)
        );
        径.setStrokeWidth(8);
        径.setFill(Color.BROWN);

        // 造组
        // HBox 横盒 = new HBox(线,文);
        // Group 组 = new Group(横盒);
        Group 组 = new Group(线, 文, 径);

        

        // 布景
        Scene 景 = new Scene(组, 800,600);
        景.setFill(Color.BISQUE);

        // 舞台设置
        舞台.setTitle("用爪语图形框架JavaFX画画");
        舞台.setScene(景);

        // 开幕
        舞台.show();
    }

    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}