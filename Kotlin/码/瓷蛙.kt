package 礼品
import java.awt.Color
/*
class 瓷蛙 {
  // var定义可变属性，val定义只读属性
  var 颜色 = Color.RED
  var 重量 = 50.0
  var 价格 = 32.9

  // 方法
  override fun toString():String{
    return ("瓷蛙描绘：\n"+"\t颜色：${颜色.toString()}\n"+
        "\t重量：${重量}克\n"+"\t价格：\$${价格}\n")
  }

  fun 瓷蛙描绘(){
    println(toString())
  }
}

 */

/*
// 定义构造函式之法1
class 瓷蛙 (颜色:Color = Color.RED, 重量:Double = 50.0, 价格:Double = 32.9)
{
  // var定义可变属性，val定义只读属性
  var 颜色 = 颜色
  var 重量 = 重量
  var 价格 = 价格

  // 方法
  override fun toString():String{
    return ("瓷蛙描绘：\n"+"\t颜色：${颜色.toString()}\n"+
        "\t重量：${重量}克\n"+"\t价格：\$${价格}\n")
  }

  fun 瓷蛙描绘(){
    println(toString())
  }
}

 */

/*
// 定义构造函式之法2
class 瓷蛙 (var 颜色:Color = Color.RED, var 重量:Double = 50.0, var 价格:Double = 32.9){
  // var定义可变属性，val定义只读属性

  // 方法
  override fun toString():String{
    return ("瓷蛙描绘：\n"+"\t颜色：${颜色.toString()}\n"+
        "\t重量：${重量}克\n"+"\t价格：\$${价格}\n")
  }

  fun 瓷蛙描绘(){
    println(toString())
  }
}

 */

// 初始化模块
class 瓷蛙 (var 颜色:Color = Color.RED, var 重量:Double = 50.0, var 价格:Double = 32.9){

  // 可有一个或多个初始化模块
  init {
    println("初始化1: 主构函式调用一次")
  }
  init {
    println("初始化2：主构函式调用一次")
    println("主构函式的参数：\n"+"\t颜色：${颜色.toString()}\n"+
        "\t重量：${重量}克\n"+"\t价格：\$${价格}\n")
  }
  // 一个或多个次构函，每个次构函必须先调用主构函一次，直接地——this()，或间接地——调用别的次构函
  // 不建议使用次构函，应使用主构函集中初始化
  constructor(高度:Double):this(){
    this.高度 = 高度
  }

  // var定义可变属性，val定义只读属性
  var 高度:Double = 1.0
  var 保费:Double // Kotlin隐式为属性生成设法set()与取法get()，也可显式定义
    get() = 价格*保率
    set(保费) {保率 = 保费/价格}
  var 保率:Double = 1.0


  // 方法
  override fun toString():String{
    return ("瓷蛙描绘：\n"+"\t颜色：${颜色.toString()}\n"+
        "\t重量：${重量}克\n"+"\t价格：\$${价格}\n"+"\t高度：${高度}\n")
  }

  fun 瓷蛙描绘(){
    println(toString())
  }
  fun 快递保费(){
    println("保费：${保费}\t 保率:${保率}")
  }
}