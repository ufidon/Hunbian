#include <iostream>
#include <memory>
#include <functional>
#include <future>

using namespace std;

class 能集
{
public:
    int operator()(int x, int y) const
    {
        cout << x + y << endl;
        return x + y;
    }
    int 加(int x, int y) const
    {
        cout << x + y << endl;
        return x + y;
    }
};

int main()
{
    能集 能;
    shared_ptr<能集> 享(new 能集);

    cout << bind(能集(), 3, 4)() << endl;
    cout << bind(&能集::加, 能, 3, 4)() << endl;
    cout << bind(&能集::加, 享, 3, 4)() << endl;

    async(能,5,7);
    async(&能集::加, &能, 5,7);
    async(&能集::加, 享, 5, 7);
}