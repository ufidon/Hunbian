package 中国.广东.深圳.动物园;
/* 档名 : 哺乳动物.java */

public class 哺乳动物 implements 动物 {

   public void 吃() {
      System.out.println("哺乳动物喝奶");
   }

   public void 跑() {
      System.out.println("哺乳动物爱跑路");
   } 

   public int 腿数() {
      return 4;
   }

   public static void main(String args[]) {
      哺乳动物 我家小牛 = new 哺乳动物();
      我家小牛.吃();
      我家小牛.跑();
   }
} 