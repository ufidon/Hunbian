#include <iostream>
#include <string>
#include <vector>
#include <memory>
#include <algorithm>

using namespace std;

int main(int argc, char const *argv[])
{
    shared_ptr<string> 龙村(new string("龙村"));
    shared_ptr<string> 凤寨{new string("凤寨")};
    shared_ptr<string> 羊堡 = make_shared<string>("羊堡");
    shared_ptr<string> 狼集;
    狼集.reset(new string("狼集"));
    shared_ptr<string> 马庄(new string("马庄"),
                            [](string *村)
                            {
                                cout << *村 << "迁走了." << endl;
                                delete 村;
                            });


    shared_ptr<string> 杏花镇(new string[3],
                            [](string *镇){
                                delete[] 镇;
                            }
    );

    杏花镇->append("一朵杏花");
    for(auto 杏: *杏花镇)
        cout << string(&杏) << endl;

    

    vector<shared_ptr<string>> 梅花镇;
    梅花镇.push_back(龙村);
    梅花镇.push_back(龙村);
    梅花镇.push_back(凤寨);
    梅花镇.push_back(龙村);
    梅花镇.push_back(凤寨);
    梅花镇.push_back(羊堡);
    梅花镇.push_back(狼集);
    梅花镇.push_back(马庄);

    //copy_backward(杏花镇->begin(), 杏花镇->back(), back_inserter(梅花镇));    

    for (auto 村寨 : 梅花镇)
        cout << *村寨 << " ";
    cout << endl;

    *龙村 = "龙潭";
    for (auto 村寨 : 梅花镇)
        cout << *村寨 << " ";
    cout << endl;

    for (int 号 = 0; 号 < 梅花镇.size(); 号++)
        cout << 梅花镇[号] << "共存于" << 梅花镇[号].use_count() << "处." << endl;

    狼集 = nullptr;
    for (int 号 = 0; 号 < 梅花镇.size(); 号++)
        cout << 梅花镇[号] << "共存于" << 梅花镇[号].use_count() << "处." << endl;

    梅花镇.resize(梅花镇.size() - 1);
    for (int 号 = 0; 号 < 梅花镇.size(); 号++)
        cout << 梅花镇[号] << "共存于" << 梅花镇[号].use_count() << "处." << endl;

    return 0;
}
