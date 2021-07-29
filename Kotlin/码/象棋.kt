package 象棋

class 象棋 {
    // 布棋：双方各5兵2炮9将帅

    // 起手：黑方(0)还是红方（1）
    // 走棋%2 ==0起手，==1对手
    // 举棋时间、举棋限时


}


abstract class 棋(棋阶:Int = 0, 黑红: Int = 0, 横位: Int = 0, 纵位: Int = 0){
    val 棋名: CharArray = charArrayOf('帅','士','相','马','车','兵','炮')
    val 棋方: CharArray = charArrayOf('红','黑')
    // 棋阶号棋名
    var 棋阶: Int = 棋阶
    // 黑红号棋方
    var 黑红:Int = 黑红

    var 横位:Int = 横位
    var 纵位:Int = 纵位

    fun 报位(){
        val 位 = "${棋方[黑红]}${棋名[棋阶]}横${横位}纵${纵位}"
        println(位)
    }
}

// 面（接口）与委托
interface 棋性{
    var 棋阶:Int
    var 黑红:Int
    var 横位:Int
    var 纵位:Int
}
interface 棋行{
    fun 走(目标横位:Int, 目标纵位:Int):Int
}
object 棋核:棋性{
    override var 棋阶:Int = 0
    override var 黑红:Int = 0
    override var 横位:Int = 0
    override var 纵位:Int = 0
}
class 棋行委:棋行{
    override fun 走(目标横位: Int, 目标纵位: Int): Int {
        return 0
    }
}

class 甩:棋性 by 棋核,棋行 by 棋行委()

class 蟀(棋阶:Int = 6, 黑红:Int = 0,
        横位:Int = 4, 纵位:Int = 0):棋(棋阶, 黑红, 横位, 纵位){
    fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
// 标以open的class才能被继承，标以open的成员才能被重载
// 棋阶，大吃小：兵车马相士帅（从小到大1-6），官兵都可吃炮（0），炮也可灭任何官兵
open class 棋子(open var 棋阶:Int = 6, open var 黑红:Int = 0,
              open var 横位:Int = 4, open var 纵位:Int = 0){
    /*
    init {
        // 布棋，下黑上红，取正常数学座标9x10：横轴0-8，纵轴0-9
    }
     */
    // 棋子在棋盘上的位置
    open fun 走(目标横位:Int = 0, 目标纵位:Int = 0):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}

// 创建子类
class 帅(override var 棋阶:Int = 6, override var 黑红:Int = 0,
        override var 横位:Int = 4, override var 纵位:Int = 0):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
class 士(override var 棋阶:Int = 5, override var 黑红:Int = 0,
        override var 横位:Int = 3, override var 纵位:Int = 0):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
class 相(override var 棋阶:Int = 4, override var 黑红:Int = 0,
        override var 横位:Int = 2, override var 纵位:Int = 0):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
class 马(override var 棋阶:Int = 3, override var 黑红:Int = 0,
        override var 横位:Int = 1, override var 纵位:Int = 0):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
class 车(override var 棋阶:Int = 2, override var 黑红:Int = 0,
        override var 横位:Int = 0, override var 纵位:Int = 0):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
class 兵(override var 棋阶:Int = 1, override var 黑红:Int = 0,
        override var 横位:Int = 0, override var 纵位:Int = 3):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}
class 炮(override var 棋阶:Int = 0, override var 黑红:Int = 0,
        override var 横位:Int = 1, override var 纵位:Int = 2):棋子(棋阶){
    override fun 走(目标横位:Int, 目标纵位:Int):Int{
        return 0 // 0-成功到空位，1-吃掉目标位棋子，-1-不能走到目标位
    }
}