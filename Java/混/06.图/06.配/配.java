/** 用fxml配置图形界面
*/
import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Scene;
import javafx.scene.input.KeyEvent;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;
import javafx.scene.Group;


import java.net.URL;


public class 配 extends Application{

    private int 场景数 = 5;

    public static void main(String[] 行参集) {
        launch(行参集);
    }

    @Override
    public void start(Stage 舞台) throws Exception {
        FXMLLoader 载器 = new FXMLLoader();
        载器.setLocation(new URL("file:///home/x/workspace/java/06.图/06.配/配.fxml"));
        //载器.setLocation(new URL("file://./配.fxml"));
        VBox 竖盒 = 载器.<VBox>load();

        Scene[] 布景 = new Scene[场景数];
        布景[0] = new Scene(竖盒);

        Group 组 = new Group();
        布景[1] = new Scene(组,500,500);

        舞台.setScene(布景[0]);
        舞台.setTitle("用FXML配置图形用户界面");
        舞台.show();
    }
}
