/* 网络客户
*/

import java.net.*;
import java.io.*;

/**
 * 客
 */
public class 客 {

    public static void main(String[] 行参集) {
        String 伺名 = 行参集[0];
        int 伺号 = Integer.parseInt(行参集[1]);

        try {
            System.out.println("连到服务器"+伺名+"的端口"+伺号);
            Socket 客槽 = new Socket(伺名,伺号);

            System.out.println("从本地槽"+客槽.getLocalSocketAddress());
            System.out.println("连到服务器槽"+客槽.getRemoteSocketAddress());

            DataOutputStream 出 = new DataOutputStream(客槽.getOutputStream());
            出.writeUTF("服务器你好，来自"+客槽.getLocalSocketAddress()+"的问候\n");

            DataInputStream 入 = new DataInputStream(客槽.getInputStream());
            System.out.println("服务器说："+入.readUTF());

            客槽.close();

        } catch (IOException io) {
            System.out.println(io.getMessage());
        }
    }
}