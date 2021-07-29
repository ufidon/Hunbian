import java.util.regex.Pattern

data class 点(var 横:Double = 0.0, var 纵:Double = 0.0){
    fun 显(){
        println("座标 =（${横}，${纵}）")
    }

    override fun toString(): String {
        return "（${横}，${纵}）"
    }
}

interface 帝{
    var 号:String
}
object  始皇:帝{
    override var 号: String= "秦始皇"
}

enum class 色(val 值: Int) {
    红(0xFF0000), 绿(0x00FF00), 蓝(0x0000FF);
}
enum class 向(val 角度: Double){
    东(0.0), 南(-90.0), 西(180.0), 北(90.0);
}

sealed class 哺乳动物类{
    override fun toString(): String {
        return "我是哺乳动物。"
    }
}
class 灵长动物类:哺乳动物类(){
    override fun toString(): String {
        return "我是灵长动物"
    }
}
class 有蹄类:哺乳动物类(){
    override fun toString(): String {
        return "我是有蹄动物"
    }
}

// 恒量只能定义为全局变量，在编译时赋值，故不能赋值以函式
const val 圆周率 = 3.1415926
// 在运行时赋值，可赋值以函式
val 地球质量 = 5.9736e24
fun 太阳系行星质量(行星:String):Double{
    return when(行星) {
        "地球" -> 5.9736e24
        "天王星" -> 8.6832e25
        else -> 0.0
    }
}
val 天王星质量 = 太阳系行星质量("天王星")
object 行星质量{
    const val 地球质量 = 5.9736e24
    const val 天王星质量 = 8.6832e25
}
class 圆{
    // 相当Java的static
    companion object{
        const val 圆周率 = 3.1415926
        fun 周长(直径:Double):Double{
            return 圆周率*直径
        }
        fun 面积(直径:Double):Double{
            return 圆周率*直径*直径/4.0
        }
    }
}

// 扩函
fun String.含数():Boolean{
    return this.contains("""\d""".toRegex())
}
fun String.含汉字():Boolean = find { it >= 0x4e00.toChar() && it <= 0x9fff.toChar() } != null
// 扩性
val String.有数: Boolean
    get() = 含数()

class 算术<型>(var 数: 型){
    companion object{
        fun <型> 加(甲: 型, 乙:型):型{
            return (甲 + 乙)
        }
    }
}

fun main(args: Array<String>) {
    val 甲 = 3
    val 乙 = 4
    val 算 = 算术<Int>(30)
    println(算.加<Int>(甲, 乙))
    /*
    println("邮编110108".含数())
    println("邮编110108".含汉字())

     */
    /*
    println("圆周率=${圆周率}")
    println("地球质量=${地球质量}")
    println("天王星质量=${天王星质量}")
    println("圆周率=${圆.圆周率}")
    val 直径 = 6.0
    println("直径为${直径}的周长为${圆.周长(直径)}")
    println("直径为${直径}的面积为${圆.面积(直径)}")
    println("地球质量=${行星质量.地球质量}")

     */
    
    /*
    // 列
    val 果园产量 = listOf<Double>(200.0, 123.3, 66.98)
    println("果园总产量：${果园产量.sum()}公斤")

    val 水果 = listOf<String>("苹果","葡萄","橙")
    println("水果名字长度和：${水果.sumOf { s: String -> s.length }}")

    println(水果)
    for (果 in 水果.listIterator()){
        println(果)
    }
    // 键值对
    val 众考生 = hashMapOf<String, Int>("张三" to 123456, "李四" to 22345, "王五" to 32345)
    println("众考生：${众考生}")
    println("张三的考号是${众考生["张三"]}")
    println("小二的考号是${众考生["小二"]}")
    println("小二的考号是${众考生.getOrDefault("小二", 55555)}")
    println("小二的考号是${众考生.getOrElse("小二", {55555})}")

    var 果产集 = mutableMapOf<String, Double>("苹果" to 100.0, "葡萄" to 212.2, "香蕉" to 232.5)
    println(果产集)
    果产集.put("李", 333.3)
    println(果产集)
    果产集.replace("苹果", 500.0)
    println(果产集)
    果产集.remove("葡萄")
    println(果产集)

     */

    /*
    // 藕
    var 国都 = "中国" to "北京"
    println("${国都.first}的首都是${国都.second}")

    // 叁
    val 高铁 = Triple("上海", "北京", 6.0)
    println("${高铁.first}到${高铁.second}的G106高铁只要${高铁.third}小时。")
    println(高铁.toList())
    println(高铁.toString())

    val 悬铁 = Triple("F106" to 6.0, "上海" to "北京", 106)
    println("悬浮列车${悬铁.first.first}共${悬铁.third}节车厢，从${悬铁.second.first}到${悬铁.second.second}只要${悬铁.first.second}小时。")
    println(悬铁.toList())
    println(悬铁.toString())

    // 取部件
    val (发,到,_) = 高铁
    println("发于${发}，到达${到}")

     */

    /*
    // 1. 独物
    println("${始皇}的值：${始皇.号}")
    // 2. 枚举
    println("红：名——${色.红.name}，号——${色.红.ordinal}，色——${色.红.值}")
    println("北：名——${向.北.name}，号——${向.北.ordinal}，向——${向.北.角度}")
    // 3. 闭类
    val 动物:哺乳动物类 = 有蹄类()
    when(动物){
        is 有蹄类 -> {
            println("有蹄类")
            println(动物)
        }
        is 灵长动物类 -> println("灵长动物类")
        else -> println("哺乳动物类")
    }

     */
    /*
    val 原点 = 点()
    原点.显()
    println("原点座标 = ${原点}")

    val 横点 = 点(5.0)
    val 叠原点 = 原点
    println("横点 == 原点：${if (横点 == 原点) "是" else "不"}")
    println("叠原点 == 原点：${if (叠原点 == 原点) "是" else "不"}")
    println("叠原点原点指同物：${if (叠原点 === 原点) "是" else "不"}")

    // 复制
    val 子点 = 横点.copy()
    println("横点 == 子点：${if (横点 == 子点) "是" else "不"}")
    println("横点子点指同物：${if (横点 === 子点) "是" else "不"}")

    // 解构
    val (横, _) = 横点
    println("横标 = ${横}")

     */
}