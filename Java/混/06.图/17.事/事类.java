/* 事件处理
* 跑：java -Dprism.forceGPU=true 材类
*/
import javafx.animation.RotateTransition;
import javafx.application.*;
import javafx.event.*;
import javafx.scene.*;
import javafx.scene.input.*;
import javafx.scene.paint.*;
import javafx.scene.shape.*;
import javafx.scene.text.*;
import javafx.scene.image.*;
import javafx.stage.*;
import javafx.scene.transform.*;
import javafx.util.*;

/**
 * 事类
 */
public class 事类 extends Application{
    private int 几次 = 0, 几次1 = 0, 次=0;
    private boolean 启停 = false;
    private double 角x = 0.0, 角y = 0.0, 角z = 0.0, 增角=1.0;

    @Override
    public void start(Stage 舞台) throws Exception{
      
        // 景物
        Circle 圆 = new Circle(200, 200, 100, Color.SEAGREEN);
        圆.setStroke(Color.YELLOW); 圆.setStrokeWidth(10);

        Text 说明 = new Text(100, 500, "请点击圆或文本");
        说明.setStroke(Color.MAGENTA); 说明.setFill(Color.CADETBLUE);
        说明.setFont(Font.font(30));

        Box 盒 = new Box(200, 200, 200);
        盒.setTranslateX(600);
        盒.setTranslateY(250);
        盒.setTranslateZ(50);
        Text 盒阐 = new Text(400, 500, "请用方向键及QH转动盒子");
        盒阐.setStroke(Color.MAGENTA); 盒阐.setFill(Color.CADETBLUE);
        盒阐.setFont(Font.font(30));    
        
        PhongMaterial 盒质 = new PhongMaterial(Color.ORANGE);
        盒质.setDiffuseMap(new Image("file:///home/x/workspace/java/06.图/00.图像/大象.png"));
        盒.setMaterial(盒质);

        RotateTransition 转 = new RotateTransition(Duration.seconds(0.8));
        转.setNode(盒);
        转.setAxis(Rotate.X_AXIS);
        转.setByAngle(360); 转.setCycleCount(5);
        转.setAutoReverse(false);

        Rotate 绕x转 = new Rotate(角x, 0, 0, 0, Rotate.X_AXIS);
        Rotate 绕y转 = new Rotate(角y, 0, 0, 0, Rotate.Y_AXIS);
        Rotate 绕z转 = new Rotate(角z, 0, 0, 0, Rotate.Z_AXIS);
                

        EventHandler<KeyEvent> 理事0 = new EventHandler<KeyEvent>(){
            @Override
            public void handle(KeyEvent 事) {
                switch (事.getCode()) {
                    case SPACE:
                        启停 = !启停;
                        if (启停) {                        
                            转.play();
                        }else{
                            转.stop();
                        }
                        break;
                    case Q:
                        角x += 增角;
                        绕x转.setAngle(角x);
                        盒.getTransforms().add(绕x转);
                        break;
                    case H:
                        角x -= 增角;
                        绕x转.setAngle(角x);
                        盒.getTransforms().add(绕x转);
                        break;
                    case UP:
                        角z += 增角;
                        绕z转.setAngle(角z);
                        盒.getTransforms().add(绕z转);
                        break;
                    case DOWN:
                        角z -= 增角;
                        绕z转.setAngle(角z);
                        盒.getTransforms().add(绕z转);
                        break;
                    case LEFT:
                        角y += 增角;
                        绕y转.setAngle(角y);
                        盒.getTransforms().add(绕y转);
                        break;
                    case RIGHT:
                        角y += 增角;
                        绕y转.setAngle(角y);
                        盒.getTransforms().add(绕y转);
                        break;
                    default:
                        break;
                }
                //绕x转.setAngle(角x);绕y转.setAngle(角y);绕z转.setAngle(角z);
                //盒.getTransforms().addAll(绕z转, 绕y转, 绕x转);
            }
        };


        // 事件处理

        // 便法
        盒阐.setOnMouseClicked(new EventHandler<MouseEvent>(){            
            public void handle(MouseEvent 事) {
                次++;
                if (次%2 == 1) {
                    盒阐.setFill(Color.YELLOW);
                }else{
                    盒阐.setFill(Color.CADETBLUE);
                }
                
            }
        });

        // 理事
        EventHandler<MouseEvent> 理事 = new EventHandler<MouseEvent>(){
            @Override
            public void handle(MouseEvent 事) {
                几次++;
                System.out.println("理事："+几次+"次=>"+事);
                
                if (几次%2 == 1) {
                    圆.setFill(Color.VIOLET);
                } else {
                    圆.setFill(Color.SEAGREEN);
                }                
            }
        };
        EventHandler<MouseEvent> 理事1 = new EventHandler<MouseEvent>(){
            @Override
            public void handle(MouseEvent 事) {
                几次1++;
                System.out.println("理事1："+几次1+"次=>"+事);
                
                if (几次1%2 == 1) {
                    圆.setStroke(Color.RED);
                } else {
                    圆.setStroke(Color.YELLOW);
                }                
            }
        };
        EventHandler<MouseEvent> 理事2 = new EventHandler<MouseEvent>(){
            @Override
            public void handle(MouseEvent 事) {
                if (事.getEventType() == MouseEvent.MOUSE_ENTERED_TARGET &&
                事.getTarget() == 说明) {
                    说明.setFill(Color.CYAN);
                } 
                if(事.getEventType() == MouseEvent.MOUSE_EXITED_TARGET &&
                事.getTarget() == 圆){
                    说明.setFill(Color.CADETBLUE);
                }
                
            }
        };        
        // 注册事件处理
        圆.addEventFilter(MouseEvent.MOUSE_CLICKED, 理事);
        

        // 布景
        Group 景物 = new Group(圆, 说明,盒, 盒阐);
        // 注意执者（handler）与滤者（filter）的执行顺序，或事件传递顺序
        // 事件由先舞台传至发事节点，一路由滤者处理；然后从事发节点传至舞台，一路由执者处理
        // 景物.addEventFilter(MouseEvent.MOUSE_CLICKED, 理事1);
        景物.addEventHandler(MouseEvent.MOUSE_CLICKED, 理事1);

        Scene 布景 = new Scene(景物,800,600);
        布景.addEventHandler(MouseEvent.MOUSE_ENTERED_TARGET, 理事2);
        舞台.addEventFilter(MouseEvent.MOUSE_EXITED_TARGET, 理事2);
        舞台.addEventFilter(KeyEvent.KEY_PRESSED, 理事0);

        PerspectiveCamera 相机 = new PerspectiveCamera(false);
        相机.setTranslateX(0); 相机.setTranslateY(0); 相机.setTranslateZ(10);
        布景.setCamera(相机);

        // 舞台
        舞台.setTitle("事件处理");
        舞台.setScene(布景);
        舞台.show();        
    }
    public static void main(String[] 行参集) {
        launch(行参集);
    }  
    
}