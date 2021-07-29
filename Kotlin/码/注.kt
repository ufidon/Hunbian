import kotlin.reflect.*

class 车(var 轮数:Int = 4){
    fun 自述(){
        println("我是${轮数}车。")
    }
    fun 驶(){
        println("启动-挂挡-松刹-走")
    }
}
fun main(args: Array<String>) {
    val 类述 = 车::class
    for (法 in 类述.members){
        println(法)
    }
}