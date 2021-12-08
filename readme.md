## 一 项目组成

### 主体

operator，user，provider三个角色的功能包

### internal

check包：支票结构体的实现

store：存储模块

cash：合约模块

common包：测试用公共包

utils包：工具函数包

### tests

测试代码

## 二 测试运行方法

先启动测试链：

我用的测试链，可以根据实际情况，将utils包中的HOST常量改为自己想用的节点。同时需要将utils中的常量chainid改成自己的测试链的实际chainid（不是networkID），否则交易出错。

测试代码位于tests目录下，一共5个测试程序，前4个测试支票支付的各个场景，最后一个测试聚合支票功能

注：测试链需要开启挖矿，确保交易执行成功