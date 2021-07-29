import java.util.*

fun 哪天():String{
  var 一周 = arrayOf<String>("周一","周二","周三","周四","周五","周六","周日")
  return (一周[Random().nextInt(6)])
}
fun 做啥(天 :String) :String{
  return when (天){
    "周一" -> "打篮球"
    "周二" -> "登山"
    "周三" -> "游泳"
    "周四" -> "打鼓"
    "周五" -> "唱歌"
    else -> "发呆"
  }
}
fun 周几做啥(){
  val 天 = 哪天()
  val 活动 = 做啥(天)
  println("我${天}${活动}")
}

fun 答对么(解: Double, 答案: Double) :Boolean{
  return if (解 == 答案) true else false
}
fun 答对(解: Double, 答案: Double) = if (解 == 答案) true else false

fun 矩形面积(长: Double, 宽:Double) :Double{
  return 长 * 宽
}
fun 矩面(长: Double, 宽:Double) :Double = 长 * 宽

fun 默参(甲: Int, 乙: Boolean = true, 丙: Double = 20.0){
  println("传入参数：甲=${甲}，乙=${乙}, 丙=${丙}。")
}

fun main(程参: Array<String>) {
  // 匿名函式（lambda），类型由kotlin隐式得出
  val 面积 = {长: Double, 宽:Double -> 长 * 宽}
  val 长 = 4.0
  val 宽 = 5.0
  println("长${长}宽${宽}的矩形面积为${面积(长, 宽)}")

  // 显式指定匿函类型
  val 矩积: (Double, Double) -> Double = {长, 宽 -> 长 * 宽}
  println("长${长}宽${宽}的矩形面积为${矩积(长, 宽)}")

  // 高阶函式（带函参的函式）
  val 求面积: (Double, Double, (Double,Double)->Double)->Unit = {
    长,宽,面积函式 -> val 面积 = 面积函式(长,宽)
    println("长${长}宽${宽}的矩形面积为${面积}")
  }
  求面积(5.0, 8.0, 矩积)
  求面积(5.0, 8.0, ::矩形面积) // 有名函式当函参

  // 若高阶函式仅最后的参数为函参，也可这样传，但会造成阅读困难
  求面积(5.0, 8.0){ 长,宽 -> 长*宽 }

  /*
  val 产品质量 = intArrayOf(53, 67, 13, 34, 88, 64)
  // 滤式默认为迫滤
  println("合格产品：" + 产品质量.filter { it >= 60 })

  // 迫滤：滤则创建完整的滤出队列，完全滤出
  val 迫滤 = 产品质量.filter { it < 60 }
  println("迫滤不合格产品：" + 迫滤)

  // 惰滤：要用到满足条件的元素时再滤出之
  val 惰滤 = 产品质量.asSequence().filter { it < 60 }
  println("惰滤不合格产品：" + 惰滤)
  val 全要 = 惰滤.toList() // 要时再滤
  println("惰滤全部结果：" + 全要)

  // 逐个要惰滤结果
  val 懒藕 = 惰滤.asSequence().map { it * -1 }
  println("懒藕：${懒藕}")

  println("懒藕第一个元素：${懒藕.first()}")
  println("懒藕出了几个：${懒藕}")
  println("懒藕第一个元素：${懒藕.elementAt(0)}")
  println("懒藕第二个元素：${懒藕.elementAt(1)}")
  println("懒藕第三个元素：${懒藕.elementAt(2)}")

  // 全要
  val 惰藕 = 惰滤.asSequence().map { it - 60 }
  println("惰藕：${惰藕.toList()}")

   */

  /*
  val 长 = 4.0
  val 宽 = 5.0
  val 面积 = 矩形面积(长, 宽)
  println("长${长}宽${宽}的矩形面积为${面积}")
  println("长${长}宽${宽}的矩形面积为${矩面(长, 宽)}")

  val 答案 = 20.0
  val 结果 = 答对(19.0, 答案)
  println("你的结果${if (结果) "对了" else "错了"}。")
  println("你的结果${if (答对么(19.0, 答案)) "对了" else "错了"}。")

  默参(1)
  默参(2,false)
  默参(3,false, 11.3)
  默参(4, 丙=11.4)
  默参(5, 丙=11.5, 乙 = false)
  默参(丙=11.6, 甲=6)
  默参(丙=11.7, 甲=7, 乙 = false)

   */


  // 周几做啥()
  /*
  val 单元 = println("欢迎, ${程参[0]}")
  println(单元)

  val 得分 = 98
  val 通过 = if (得分 > 70) true else false
  println(通过)
  println("我${if (得分>70) "通过" else "挂科"}了。")

   */
}