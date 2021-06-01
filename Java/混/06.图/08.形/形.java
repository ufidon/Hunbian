/* 爪语图架基本图形
*/
import javafx.application.*;
import javafx.geometry.*;
import javafx.scene.Group;
import javafx.scene.Scene;
import javafx.scene.effect.BlendMode;
import javafx.scene.paint.Color;
import javafx.scene.shape.*;
import javafx.stage.*;


/**
 * 形
 */
public class 形 extends Application{
    private int 台宽 = 800;
    private int 台高 = 600;

    @Override
    public void start(Stage 舞台) throws Exception {
        // 基本图形
        // 平面点, 画10*10的网格，横标数=台宽/10, 纵标数=台高/10
        int 横标数 = 台宽/10, 纵标数 = 台高/10;

        Point2D[][] 点 = new Point2D[横标数][纵标数];
        for (int x = 0; x < 横标数; x++) { // 点号按列排序，从上到下，从左到右
            for(int y=0; y<纵标数; y++){
                点[x][y] = new Point2D(x*10, y*10);
            }
        }
        // 直线
        Line[] 直线 = new Line[4];
        for (int i=0; i<直线.length; i++) {
            直线[i] = new Line();
        }
        直线[0].setStartX(点[1][1].getX()); 直线[0].setStartY(点[1][1].getY());
        直线[0].setEndX(点[1][纵标数-1].getX()); 直线[0].setEndY(点[1][纵标数-1].getY());
        直线[0].setBlendMode(BlendMode.OVERLAY); // 怎样融入背景

        // 矩形
        Rectangle 矩形 = new Rectangle();
        矩形.setX(5); 矩形.setY(5);
        矩形.setWidth(台宽-10); 矩形.setHeight(台高-10);
        矩形.setFill(Color.SKYBLUE);
        矩形.setArcHeight(18); 矩形.setArcWidth(24); // 圆角

        // 圆
        Circle 日 = new Circle(600,150,60);
        日.setFill(Color.RED);
        日.setStroke(Color.ORANGERED);
        日.setStrokeWidth(5);

        // 椭圆
        Ellipse 菇头 = new Ellipse(200, 400, 80, 40);
        菇头.setFill(Color.PINK);
        Rectangle 菇干 = new Rectangle(180, 410, 40, 120);
        菇干.setArcWidth(16); 菇干.setArcHeight(10);
        菇干.setFill(Color.ANTIQUEWHITE);

        // 多边形
        Polygon 树冠 = new Polygon(600,300,500, 450,700,450);
        树冠.setFill(Color.LIGHTGREEN);
        Polygon 树干 = new Polygon(575,450,625,450,625,550,575,550);
        树干.setFill(Color.DARKSALMON);

        // 曲线
        Path 云 = new Path(
            new MoveTo(120, 100),
            new CubicCurveTo(140, 70, 200,70, 240, 100),
            new ArcTo(20, 40, 40, 240, 160, true, false),
            new QuadCurveTo(170, 200, 120, 160),
            new ArcTo(20, 40, 30, 120, 100, false, true)
        );
        云.setStroke(Color.WHITESMOKE);
        云.setFill(Color.WHITESMOKE);

        // svg路径
        SVGPath 残月 = new SVGPath();
        String 月廓 = "M300,200 h-150 a150,150 0 1,0 150,-150 z";
        残月.setContent(月廓);
        残月.setFill(Color.SILVER);
        残月.setStroke(Color.ROSYBROWN); 残月.setStrokeWidth(5);

        // 布景
        Group 组 = new Group(直线); // 直线被矩形遮住了，看不见
        组.getChildren().addAll(矩形);
        组.getChildren().addAll(日);
        组.getChildren().addAll(菇干);
        组.getChildren().addAll(菇头);
        组.getChildren().addAll(树冠);
        组.getChildren().addAll(树干);
        组.getChildren().addAll(云);
        组.getChildren().addAll(残月);

        Scene 布景 = new Scene(组, 台宽, 台高);
        布景.setFill(Color.YELLOWGREEN);
        
        // 舞台
        舞台.setTitle("爪语图架基本图形");
        舞台.setScene(布景);

        舞台.show();
    }
    
    public static void main(String[] 行参集) {
        Application.launch(行参集);
    }
}