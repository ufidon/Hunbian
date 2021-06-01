/* 多线程，继承Thread
 */


class 线程 extends Thread{
    private Thread 线;
    private String 线名;

    public  线程(String 线名) {
        this.线名 = 线名;
        System.out.println(线名+"出场☑");
    }

    public void run() { // 线程函数
        System.out.println(线名+"开跑");
        try {
            int 次 = 5;
            while (次>0) {
                次--;
                System.out.printf("%d, %s加油\n", 次, 线名);
                Thread.sleep(1000); // 1秒
            }           
        } catch (InterruptedException i) {
            System.out.printf("%s被打断了！\n", 线名);
        }
        System.out.printf("%s跑完全程！☺\n", 线名);
    }
    
    public void start() {// 启动线程
        System.out.printf("%s进入跑道\n", 线名);
        if (线 == null) {
            线 = new Thread(this, 线名);
            线.start();
        }
    }
}

public class 多线 {
    public static int 线程数 = 3;

    public static void main(String[] 行参集) {
        String[] 线程名 = {"李世民", "李元霸", "李元吉"};
        线程[] 线程集 = new 线程[线程数];
        for (int i=0; i<线程数; i++) {
            线程集[i] = new 线程(线程名[i]);
            线程集[i].start();
        }
    }
}