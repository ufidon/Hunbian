/* 物体存取
*/

import java.io.*;

public class 存 {

    public static void main(String[] 行参集) {
        String 狗档 = "./狗.txt";
        // 1. 存狗档
        // 狗 李家狗 = new 狗("老李", "旺财", 12345);        
        // 存档(李家狗, 狗档);

        // 2. 取狗档
        狗 某狗 = 取档(狗档);
        System.out.printf("狗档《%s》里的信息：\n",  狗档);
        System.out.println(某狗);        
    }

    public static void 存档(狗 某狗, String 狗档) {
        try {
            FileOutputStream 档 = new FileOutputStream(狗档);
            ObjectOutputStream 信 = new ObjectOutputStream(档);
            信.writeObject(某狗);
            信.close();
            档.close();
            System.out.printf("%s狗%s档案存于档案%s\n", 某狗.主, 某狗.名, 狗档);
        } catch (IOException 障) {
            System.out.println(障.getMessage());
        } 
    }

    public static 狗 取档(String 狗档) {
        狗 某狗 = null;

        try {
            FileInputStream 档 = new FileInputStream(狗档);
            ObjectInputStream 信 = new ObjectInputStream(档);
            某狗 = (狗)信.readObject();
            信.close(); 档.close();

        } catch (IOException io) {
            System.out.println(io.getMessage());
        }catch(ClassNotFoundException c){
            System.out.println(c.getMessage());
        }

        return 某狗;
    }
}