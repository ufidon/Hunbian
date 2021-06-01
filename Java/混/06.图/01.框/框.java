/* JavaFX显示一个图形框
*/
import javafx.application.*;
import javafx.scene.*;
import javafx.scene.paint.*;
import javafx.stage.*;

public class 框 extends Application{
    @Override
    public void start(Stage 舞台) throws Exception {
        // 造组
        Group 组 = new Group();
        // 造景
        Scene 景 = new Scene(组, 800,600);
        景.setFill(Color.BISQUE);

        // 舞台设置
        舞台.setTitle("爪语图形框架JavaFX");
        舞台.setScene(景);

        // 开幕
        舞台.show();
    }

    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}