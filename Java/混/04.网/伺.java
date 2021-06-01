/* 服务器
*/
import java.net.*;
import java.io.*;

/**
 * 伺
 */
public class 伺 extends Thread {
    private ServerSocket 伺槽;

    public 伺(int 端口号) throws IOException {
        伺槽 = new ServerSocket(端口号);
        伺槽.setSoTimeout(10000); // 10秒
    }

    public void run() { // 线程函数
        while (true) {
            try {
                System.out.printf("\n\n服务器在端口%d上等待客户连接……\n", 伺槽.getLocalPort());

                Socket 会客槽 = 伺槽.accept();
                System.out.printf("客户槽%s连到\n", 会客槽.getRemoteSocketAddress());
                System.out.printf("本地槽%s\n", 会客槽.getLocalSocketAddress());

                DataInputStream 入 = new DataInputStream(会客槽.getInputStream());
                System.out.printf("客户%s说：%s", 会客槽.getRemoteSocketAddress(), 入.readUTF());
                DataOutputStream 出 = new DataOutputStream(会客槽.getOutputStream());
                出.writeUTF(String.format("谢谢您连到服务器%s\n再会!\n" , 会客槽.getLocalSocketAddress()));
                会客槽.close();
            } catch (SocketTimeoutException s) {
                System.out.println(s.getMessage());
                break;
            }catch(IOException io){
                System.out.println(io.getMessage());
                break;
            }
        }
    }

    public static void main(String[] 行参集) {
        int 端口号 = Integer.parseInt(行参集[0]);
        try {
            Thread 线1 = new 伺(端口号);
            线1.start();
        } catch (IOException io) {
            System.out.println(io.getMessage());
            io.printStackTrace();
        }
    }
    
}