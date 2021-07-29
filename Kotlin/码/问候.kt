fun main(行参集: Array<String>) {
/*
    println("你叫什么名字？")
    val 名字 = readLine()
  var 年龄: Int = 30;
  var 姓: String = "王"

  var 数组 = Array<Int>(6,{i->i+3})

    println("${名字}，你好。你姓${姓}，今年${年龄}岁！")
  System.out.println(数组.toString())
  println("${数组}")

 */
  /*
  var 读者 = Scanner(System.`in`)
  val 年龄 = 读者.nextInt();
  var 名字 = 读者.next();
  println("${名字}今年${年龄}")

   */
  /*
  println("==猜数游戏==")
  var 猜 = -1
  var 读者 = Scanner(System.`in`)
  var 靶 = Random(10).nextInt(10)
  猜数游戏@ while (true){

    print("请输入一个数:")
    猜 = 读者.nextInt()
    when(猜){
      靶 -> {
        println("恭喜！你猜中了！")
        break@猜数游戏
      }
      else->{
        var 提示 = if (猜<靶) "比目标数小" else "比目标数大"
        println("没有猜中，${提示}， 请继续")
      }
    }
  }

   */

  /*
  println("二元计算器==\n")

  var 加 : (Int, Int)->Int = {a:Int, b:Int -> a+b}
  var 减 : (Int, Int)->Int = {a:Int, b:Int -> a-b}

  var 计算器: (Int, Int, (Int,Int)->Int)->Int = {a:Int,b:Int,f:(Int,Int)->Int -> f(a,b) }
  println(计算器(2,3,加))
  println(计算器(2,3,减))

   */

  /*
  println("字符==")
  val 名字 = "张三丰☞√嚻"
  println("${名字}共有${名字.length}字")

   */

  /*
  println("==空==")
  var 物:Any? = "任物"
  if (物 is String){
    println("${物}长${物.length}")
  }
  物 = null
  if (物 is String){
    println("${物}长${物.length}")
  }

   */

  /*
  println("${行参集[0]} ${行参集[1]}")
  println("==安全转换==")
  var 地址 :Any? = "长江以南"
  var 名字 :String? = 地址 as? String?
  var 年龄 :Int? = 地址 as? Int?
  println("名字：${名字} 年龄：${年龄}")

   */
  /*
  println("紧凑函数==")
  fun 太热了吗(温度: Double) = 温度>50.0
  println(太热了吗(60.0))

   */
  /*
  println("懒==")
  val 桃 = listOf("桃花", "桃树", "桃叶","桃子", "橙子")
  val 急 = 桃.filter { it[0] == '桃' }
  val 懒 = 桃.asSequence().filter { it[0] == '桃' }
  val 迫 = 桃.toList()
  println("急：${急}")
  println("懒：${懒}")
  println("迫：${迫}")
    */

}