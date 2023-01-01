#include <iostream>
#include <cstddef>
using namespace std;

void 哪(int){
    cout << "整参哪" << endl;
}

void 哪(void *){
    cout << "空指针哪" << endl;
}

int main(){
    哪(0);
    // 哪(NULL); // 编译器不能确定调用哪个哪
    哪(nullptr);
    // 哪(nullptr_t); // 非期望类型
}