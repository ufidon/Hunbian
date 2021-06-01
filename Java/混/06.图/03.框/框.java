/* JavaFX显示一个图形框
*/
import javafx.application.*;
import javafx.scene.*;
import javafx.scene.control.Button;
import javafx.scene.input.KeyCode;
import javafx.scene.input.KeyCodeCombination;
import javafx.scene.input.KeyCombination;
import javafx.scene.layout.VBox;
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

        VBox 竖盒 = new VBox(4);
        竖盒.getChildren().addAll(new Button("确认"), new Button("取消"));
        Scene 景2 = new Scene(竖盒,500,500);

        // 舞台设置
        舞台.setTitle("爪语图形框架JavaFX");
        舞台.setScene(景);

        // 舞台2
        Stage 舞台2 = new Stage();
        舞台2.setTitle("舞台2");
        舞台2.setScene(景2);
        //舞台2.initModality(Modality.APPLICATION_MODAL); // 关闭舞台2后才能操作别的舞台

        舞台2.initModality(Modality.WINDOW_MODAL);
        舞台2.initOwner(舞台); // 关闭舞台2，才能操作其拥舞台

        //舞台2.initModality(Modality.NONE); // 不阻拦任何别的舞台

        // 舞台2.initStyle(StageStyle.DECORATED);
        // 舞台2.initStyle(StageStyle.UNDECORATED); //需添加关闭方法
        // 舞台2.initStyle(StageStyle.TRANSPARENT); //需添加关闭方法
        // 舞台2.initStyle(StageStyle.UNIFIED);
        // 舞台2.initStyle(StageStyle.UTILITY);

        舞台2.setFullScreen(true);
        舞台2.setFullScreenExitKeyCombination(new KeyCodeCombination(KeyCode.KP_LEFT));

        // 开幕
        舞台.show();
        //舞台.showAndWait(); // 不能用于主舞台

        舞台2.showAndWait();
        //舞台2.show();

    }

    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}