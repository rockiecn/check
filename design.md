## 一、Order

### 1.1 Order

```go
type Order struct {
    ID uint64				// 订单ID
    Value *big.Int			// 货币数量
    Token common.Address	// 货币类型
    Fee uint64				// 应付金额
    From common.Address		// user地址
    To common.Address		// provider地址
    Time date				// 订单提交时间
    Name string				// 购买人姓名
    Tel string				// 购买人联系方式
    Email string            // 接收支票的邮件地址
    Paid bool               // 标记是否已付款
    
    Nonce uint64			// 支票nonce，申请的时候无法填入，由运营商在生成支票后补填，并存储到订单池中
    
    Sig string				// 运营商的签名
}
```

### 1.2 OrderPool

```go
type OrderPool struct {
    // user -> []*Order
    Data map[common.Address][]*Order
}

// 将一张订单存储到每个user各自的队列下面，以订单ID为排列顺序。
func (pool *OrderPool) Store(o *Order) error {
    
}

// 根据user地址和订单ID来从订单池获取一张订单
func (pool *OrderPool) Get(user common.Address, ID uint64) (* Order, error) {
    
}

// 在用户支付费用并为其生成check以后，向订单补填生成的支票nonce
func (pool *OrderPool) SetNonce(od *Order, nonce uint64) error {
    
}

// 在用户支付完成后，使用订单的user和ID信息来调用，以修改此订单的支付状态paid为true。
func (pool *OrderPool) SetPaid(od *Order, paid bool) error {
    
}
```

## 二、Operator

### 2.1 Contract

```go
// 部署合约
func (op *Operator) DeployContract(value *big.Int) (*tx types.Transaction, ctrAddr common.Address, err error) {

}

// 合约余额查询
func (op *Operator) QueryBalance() (*big.Int, error) {

}

// 查询指定节点的当前nonce。
func (op *Operator) QueryNonce(to common.Address) (uint64, error) {

}

// 向银行合约充钱
func (op *Operator) Deposit(value *big.Int) error {

}
```

### 2.2 GenCheck *

```go
// 用户订单支付完成后，使用user提交的订单来生成一张check
func (op *Operator) GenCheck(o *Order) (*check.Check, error) {

}
实现逻辑：
先在支票池中根据订单中的节点地址定位到对应的支票队列，然后查看末尾支票的nonce值maxNonce（队列中最大），然后使用maxNonce+1做为nonce来生成新的check。
如果支票队列为空，表示这是给节点支付的第一张支票，nonce设为1。
另外支票还需加入operator自身的地址等相关信息。
```

### 2.3 Mail？

```go
// 根据用户的订单，将对应生成的支票（已付费）通过邮件发送给user
func (op *Operator) Mail(od *Order) error {

}
实现逻辑：
先使用订单在支票池调用GetCheck方法找到指定支票，然后将支票发送到订单中的邮箱地址。
```

### 2.4 Aggregate

```go
// 聚合支票，运营商对节点提供的将多张小额支票聚合成一张大额支票返回给节点
func (op *Operator) Aggregate(data []byte) (batch *check.BatchCheck, error) {

}
实现逻辑：
先将收到的数据反序列化（数据结构成员包含paycheck数组）
验证数组中每一张paycheck的签名（operator和user）
确认每张支票的payvalue值不能超过其value值。
计算这批paycheck的payvalue累加值（总支付金额）
找到这批paycheck的minNonce和maxNonce
验证所有paycheck的节点地址to是否一样
计算operator对聚合支票的签名sig，签名内容包含所有数据成员
然后使用节点地址，累计总金额，minNonce，maxNonce，sig，生成聚合支票。
返回聚合支票batch
```

### 2.5 Refund

```go
当user的一张check或者paycheck需要退钱的时候，如何实现？

退款分为常规退款和异常退款:

1.常规退款
某张check，用于生成paycheck支付以后还剩余了一点金额，这张paycheck需要退钱。
先根据支票的from，to，value，payvalue来在区块链的event中找到此支票的提现记录。然后根据value和ayvalue的差值来确定具体该退多少钱。
2.异常退款
User的某张check，由于某种原因还没使用就过期了，无法用于支付，需要退钱。比如这张支票迟到了。

运营商如何知道user的某张支票到底花掉了多少？
查询链上的交易历史，查到此支票被兑现了多少钱，就可以算出应该退多少钱了。如果链上不存在此支票的交易历史，说明此支票没有被使用过，需要全额退款。

需要在每次provider提现的时候，使用event记录一个提现的事件，并记录一个topic以便以后可以根据交易信息进行查询。单独使用一套监控机制来监控每一次提现事件。

如何知道某张支票是否已经退过钱了（这个可以在check池中想办法解决）？
运营商线下保存每次退钱记录，先考虑在check池中记录每次退钱动作。
建立一个退钱池，记录跟每个节点相对应的所有已退款的支票。

//退款池，记录每张已退款的支票。
type RefundPool struct {
    // to -> []*check
    Data map[common.Address][]*check.Check
}

// 将退款支票记录到池中
func (p *CheckPool) Record(c *check.Check) error {

}
实现逻辑：
先验证支票check的签名合法性，然后在退款池中查询此支票是否已经存在，如果已经存在，则返回错误。最后将要退款的支票，以nonce为顺序，放入到对应的支票切片中。

// 两个预退款方法，正常/异常退款，在运营商将钱退还给user之前调用，此调用成功了再将钱退还给user。如果执行失败了则不能退钱。
// 正常退款预处理，退还paycheck余额。这里需要支票和退款人user的地址，用于验证user的签名。Paycheck由需要退款服务的user提供。
func (p *CheckPool) NomalRefund(pc *check.Paycheck) error {

}
实现逻辑：
先验证paycheck的签名合法性（user签名和operator签名）。
然后判断此支票是否已经退过款了（看池中是否已经存在）。
然后在链上查询此paycheck是否有提现记录，根据关键字to，nonce来查询。
最后调用退款池的Record方法，将要退款的支票记录到池中。
记录成功则返回成功，否则返回失败。

// 异常退款预处理，退还迟到支票的所有金额。
func (p *CheckPool) ExceptionRefund(c *check.Check, user common.Address) error {

}
实现逻辑：
先验证check的签名合法性。
然后判断此支票是否已经退过款了（看池中是否已经存在）。
然后在链上查询此支票是否被提现过。使用关键字to，nonce来查询。
最后调用Record方法将它记录到退款池中。
记录成功后返回成功，否则返回失败。
```

### 2.6 CheckPool

```go
type CheckPool struct {
    // to -> []check
    Data map[common.Address][]*check.Check
}

// 以nonce大小为序，将支票存储到支票池合适的位置。
func (p *CheckPool) Store(chk *check.Check) error {

}
先查看当前nonce是否越界
如果nonce越界，则先使用nil填充池，直到nonce前的位置，然后把nocne添加到pool中
如果nonce没有越界，并且check已经存在于池中，则报错返回
如果nonce没越界，并且check不存在于池中，则将check直接放到nonce指定的位置

// 根据订单来从支票池中获取对应支票。
func (p *CheckPool) GetCheck(od *Order) (*check.Check, error) {

}
```

### 2.7 Questions

```go
订单系统如何跟运营商沟通？
用户下了订单以后，如何通过订单系统调用operator的方法来存储支票？
以及支票付款以后如何调用operator的生成支票的GenCheck方法？
```

## 三、User

### 3.1 购买支票流程

```go
用户先向运营商提交订单Order，运营商审核通过后，等待user付款。
用户付款完成后，运营商根据订单生成支票check存储到本地支票池，然后把支票文件通过邮件发送给user。
用户在邮箱中收取支票文件。

疑问：支票的金额如何定？
购买支票的金额是根据用户的要求来定，还是根据本次数据传输大小来定？
如果根据数据大小来定支票金额，那么如何付款？
```

### 3.2 ImportCheck

```go
// 用户支付成功后会收到一个支票文件，通过支票文件导入支票
func (user *User) ImportCheck(path string) (*check.Check, error) {

}
```

### 3.3 Pay*

```go
顺序支付方案：
收到的数据块放入队列，然后依次为它们发送paycheck并获得user确认后，再发下一个paycheck
聚合支付方案：
同时接收多个文件的多个数据块，然后将它们生成一张大的支票发送给user以完成支付

开始传输之前的协商过程：
开始传输数据和支付支票前，user和provider需要有一个协商过程:
用户先把支票发送给节点，节点验证通过以后，向用户发送第一个数据块。
用户收到数据块后，生成paycheck发送给节点，节点验证paycheck通过以后，再发后续数据块。

user收到数据块后的支付流程：
1.用户可能会同时收到provider发来的多个数据块（同时传输多个文件的情况），那么需要把这些数据块放置到一个队列里面，每次取出一个块，然后针对这个块生成paycheck发送给provider。
2.对每个数据块：
调用paycheck池的GetCurrent取出current支票，查看其余额是否足够支付blockvalue。
如果current支票余额足够，则使用累加blockvalue后的payvalue值，调用GenPaycheck生成一张新的paycheck，用其替换存储池中的current支票。
如果current支票的余额不够支付，或者paycheck池为空，需要调用check池的GetNew方法，获取一张已购买但尚未使用过的check

疑问：考虑获取到的支票无法支付怎么办？如果获取到的支票nonce没有大于当前支付的支票nonce，说明此支票还没用就已经过期了（比如它迟到了），导致被user放到支票池的时候就已经有nonce更大的支票在支付了。这种情况需要将此支票放入到退钱池中，以便以后向运营商退钱。同时，运营商需要有一个将paycheck的剩余金额退钱的机制。

可以重点考虑provider节点在收到paycheck时的验证细节，因为只要provider拒收了user的paycheck，那么user就会重新选择合法的paycheck进行支付（需要根据provider告知的拒收原因做重发依据）。

user方发出的paycheck有很多原因会在provider那里验证失败，重点依靠重发机制来做容错处理。

然后使用此check，以blockvalue作为参数调用GenPaycheck来生成一张新的paycheck替换掉paycheck池里的current支票。

如果GetNew返回空，说明user的支票池中已经没有可支付支票，那么提示用户需要购买check。

注意：
在GetNew获取到一个新支票用于支付以后，万一还有一张迟到的具有较小nonce的新购买支票存在，则会导致这张迟到的支票过期（nonce值小于当前正在支付的paycheck），不能用于支付。

疑问1：user如何知道provider是否成功接收到了paycheck?
答：这个是网络层要实现的逻辑，支付层只需要定义好支付（send）接口，然后在返回值中来判断本次发送的结果状态。

疑问2：user如何知道provider已经传输完毕，何时结束支付？
答：数据是否传输完成由user来决定，不是由provider来决定。比如user想要的文件大小为100M，当user成功接收到100M数据了，就知道本次传输结束了。如果收到的数据超过了100M，则表示接收的数据不正确，需要报错。
```

### 3.4 GetNew *

```go
// 从check池中取出一张指定节点的新check（未使用过的check）。
func (user *User) GetNew(to common.Address) (*check.Check, error){
    
}
实现逻辑：
// 如果paycheck队列为空，表示没有支票被用过，所有的check都能支付，则直接取出check池的第一张支票返回。
// 如果paycheck队列不为空，则在paycheck队列中取出末尾项的nonce（队列中的最大nonce）
// 然后在check池中从nonce+1开始向后找，一直找到存在支票的数据项返回。
// 如果一直找到切片末尾都是空值，表示当前已无可用支票，返回空。
```

### 3.5 GenPaycheck

```go
// 使用check和payvalue值来生成一张paycheck
func (user *User) GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error) {

}
实现逻辑：
使用check，以及payvalue来生成一张paycheck，并存储到paycheck池。
其参数为通过GetNew方法来返回的没使用过的check。
```

### 3.6 BlockValue？

```go
// 数据块价值计算方法，考虑放到common包中
func (user *User) BlockValue(size *big.Int, factor int64) *big.Int{
	return size.Mul(size, factor)
}
实现逻辑：
根据数据块大小，以及价格系数来确定数据块的实际价值。
```

### 3.7 SendPaycheck？

```go
// 将paycheck发送给目标节点
func (user *User) SendPaycheck(to common.Address, pc *check.Paycheck) error{

}
实现逻辑：
将paycheck发送给provider节点，以支付本次数据块的费用。
provider在收到paycheck以后，如果paycheck验证不通过，必须要通知user知道，user好重新选择一张check继续支付数据费用。否则传输会挂起。

疑问：provider如何把验证结果反馈给user？
```

### 3.8 PreStore*

```go
// 存储支票进池之前，验证接收到的check
func (user *User) PreStore(chk *check.Check) (bool, error) {

}
实现逻辑：
验证支票签名(operator)。
支票的from字段是否等于user地址。
支票的nonce必须大于合约中节点地址对应的当前nonce。
支票在本地池中不能已存在（不能有相同nonce）。
验证通过返回true，否则返回false
```

### 3.9 CheckPool

```go
// 每个目标节点to对应一个支票队列，以nonce大小为序
type CheckPool struct {
    // to -> []*check.Check
    Data map[common.Address][]*check.Check // nonce有序
}

// 存储一张新收到的check到check池：
func (p *CheckPool) Store(c *check.Check) error {

}
实现逻辑：
先查看当前nonce是否越界
如果nonce越界，则先使用nil填充池，直到nonce前的位置，然后把nocne添加到pool中
如果nonce没有越界，并且check已经存在于池中，则报错返回
如果nonce没越界，并且check不存在于池中，则将check直接放到nonce指定的位置
```

### 3.10 PaycheckPool*

```go
type PaycheckPool struct {
    Data map[common.Address][]*check.Paycheck // nonce有序 	
}

// 将paycheck发送给节点用于支付，并在收到节点的验证确认后，将paycheck存储到池中，用于下一次支付计算支付金额时用。
func (p *PaycheckPool) Store(pc *check.Paycheck) error {

}
实现逻辑：
同上

// 获取当前正在支付的paycheck（队列中nonce最大的那个）
func (p *PaycheckPool) GetCurrent(to common.Address) (*check.Paycheck, error) {
    
}
实现逻辑：
直接返回len(slice)-1位置的paycheck，就是nonce最大的那个。

疑问：

1.是否允许user同时向provider请求多个不同的文件？
允许。

2.如果同时向provider请求多个文件，那么在支付这些文件的时候是使用不同的支票进行支付，还是使用同一个支票支付所有的文件？

注：如果使用相同的支票支付所有请求，那么一旦这张支票余额不够了，会导致所有支付动作挂起。而且在paycheck的角度来看，也无法区分不同的文件请求，它只知道接收到了一个数据块，但不知道这个数据块属于哪个文件。

所以这里还是使用同一张支票来支付所有文件请求。如果支票余额不足，则会挂起用户的所有支付及传输动作，直到用户购买并使用了新支票。
```

## 四、Provider

### 4.1 提现流程

```go
流程：
调用paycheck池的GetNextPayable()方法，找到下一个能提现的paycheck
如果没找到，则返回错误，表示当前没有可用于提现的paycheck。
如果找到了，先使用此paycheck的nonce更新Provider结构的txNonce值。
然后使用此paycheck作为参数，调用SendTx向链发送提现交易，以获取收益，并更新合约中的nonce值。

疑问：合约交易发送成功了是否能保证此交易一定能上链？
答：不能保证。但是不影响提现流程，因为如果一个paycheck没有上链成功，那么合约的nonce值就没有被改变，在下次提现的时候，还是会把这个paycheck再次找出来（依据合约的nonce来查找），并再次使用它发出提现交易。
```

### 4.2 SendTx

```go
// 向区块链发送提现交易
func (pro *Provider) SendTx(pc *check.Paycheck) (tx *types.Transaction, err error) {

}
实现逻辑：
使用paycheck为参数，向链发送提现交易，跟合约交互后向provider付款。
```

### 4.3 CalcPay

```go
// 计算收到的paycheck实际支付金额
func (pro *Provider) CalcPay(pchk *check.Paycheck) (*big.Int, error) {

}
实现逻辑：
先查看当前nonce是否超过了切片的长度，如果超过了，说明这是一张新用于支付的支票，需要将切片扩展到当前nonce的位置，并且将它存放到nonce所在位置，然后返回它的payvalue作为支付金额。
如果nonce在切片当前长度范围内，则先看此nonce位置知否已经存在paycheck。
如果不存在，则将它存放到当前nonce位置，并返回其payvalue值。
如果已存在，则计算当前支票和nonce所在位置的paycheck的payvalue差值并返回。

疑问：
万一出现支票池数据丢失的情况怎么办？支票池没数据就无法正确计算支付金额了。
```

### 4.4 PreStore*

```go
// 存储paycheck之前的验证步骤
func (pro *Provider) PreStore(pc *check.Paycheck, size uint64) (bool, error) {

}
实现逻辑：
验证一张paycheck的合法性。
首先是两个签名是否正确。
然后是value值是否大于payvalue值。
然后是to地址跟provider地址是否相同。
nonce值是否大于合约中to地址的当前nonce（决定了它是否能够提现）。
nonce值是否大于txNonce的值（决定了它是否能够提现）。
计算实际支付金额（CalcPay）是否等于数据块的自身价值。
```

### 4.5 PaycheckPool

```go
type PaycheckPool struct {

}
Data []*check.Paycheck 	//按照nonce有序

// 存储一张paycheck到池中
(p *PaycheckPool) Store(pc *check.Paycheck) error {

}
实现逻辑：
同上。

// 从支票池中取出能提现的支票中nonce最小的非current支票的paycheck（因为current支票是当前正在用于支付的paycheck）进行提现。
(p *PaycheckPool ) GetNextPayable() (*check.Paycheck, error) {

}
实现逻辑：
// 先查看合约中节点对应的nonce，然后在本地paycheck池中找出第一个比它大的支票
// 如果此支票不是current支票就正常返回它。
// 如果没有合适的可提现paycheck，就返回空
```

