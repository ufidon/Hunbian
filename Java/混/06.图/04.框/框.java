/* JavaFX显示一个图形框
*/
import javafx.application.*;
import javafx.scene.*;
import javafx.scene.input.KeyCode;
import javafx.scene.input.KeyEvent;
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


        // 舞台布景
        舞台.setTitle("爪语图形框架JavaFX");
        舞台.setScene(景);

        // 事件处理

        舞台.setOnHiding((事)->{
            System.out.println("主窗口将要隐藏"+事.toString());
        });
        舞台.setOnHidden((事)->{
            System.out.println("主窗口隐藏了"+事.toString());
        });
        舞台.setOnShowing((事)->{
            System.err.println("主舞台将要显示"+事.toString());
        });
        舞台.setOnShown((事)->{
            System.out.println("主窗口显示了"+事.toString());
        });
        舞台.setOnCloseRequest((事)->{
            System.out.println("主窗口被点击关闭"+事.toString());
        }); 
        
        舞台.addEventHandler(KeyEvent.KEY_PRESSED, (事)->{
            System.out.println("有键被按下："+事.toString());
            // System.out.println("事.getCode()："+事.getCode());
            // System.out.println("事.getCode().getCode()："+事.getCode().getCode());
            // System.out.println("KeyCode.UP:"+KeyCode.UP);
            
            switch (事.getCode()) {
                case ESCAPE:
                    舞台.close();
                    break;
                case LEFT:
                    舞台.setX(舞台.getX()-10.0);
                    break;
                case RIGHT:
                    舞台.setX(舞台.getX()+10.0);
                    break;
                case UP:
                    舞台.setY(舞台.getY()-10.0);
                    break;
                case DOWN: // ☹枚举值必须是单独的，不能是KeyCode.DOWN, 
                    舞台.setY(舞台.getY()+10.0);
                    break;
                case H:
                    景.setCursor(Cursor.CLOSED_HAND);
                    break;
                case C:
                    景.setCursor(Cursor.CROSSHAIR);
                    break;
                default:
                    System.out.println("按下的键是："+事.getCharacter());
                    break;
            }
            
        });
        
        // 开幕
        舞台.show();
    }

    public static void main(String[] 行参集) {
        Application.launch(行参集);
        System.out.println("行参集："+行参集);
    }
}