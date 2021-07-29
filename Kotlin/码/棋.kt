package 象棋

fun main(行参: Array<String>){
    val 黑蟀 = 蟀(棋阶 = 0, 黑红 = 1)
    黑蟀.走(4,5)
    黑蟀.报位()
}