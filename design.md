## 一、Operator功能

### 1.1 相关数据结构

// 用户提交的购买申请，线上或线下提交

type Apply struct {

Value *big.Int				// 购买的支票金额

Token common.Address		// 购买的货币类型

From common.Address		// 支票的支付方地址

To common.Address		// 支票的接收方地址

Date	 date				// 购买日期

Name string				// 购买人姓名

Tel string					// 购买人联系方式

...

Sig string					// 运营商的签名

}

 

// 收据，在user付款后生成并发送给user的收据

type Receipt struct {

Dt Date 				// 购买日期

Value *big.Int			// 购买金额

Token common.Address 		// 货币类型

Op common.Address	// 运营商地址

From common.Address	// 付款方地址

To common.Address	// 收款方地址

Nonce uint64			//  nonce

...

Sig					// 运营商的签名

}

### **1.2 合约维护**

1.部署合约

DeployContract(value *big.Int) (tx *types.Transaction, ctrAddr common.Address, err error) {

部署银行合约。

}

2.合约余额查询

QueryBalance() (*big.Int, error) {

查询合约当前余额。

}

3.节点nonce查询

QueryNonce(to common.Address) (uint64, error) {

查询指定provider的nonce。

}

4.向银行合约充钱

Deposit(value *big.Int) error {

向合约转账。

}

### **1.3 GetReceipt（运营商收款后提供收据）**

// user付款成功后调用

(op *Operator) GetReceipt(a *Apply) (*Receipt, error) {

使用user提供的购买申请Apply，调用GenCheck生成一张check并存储到check池，然后将生成的收据Receipt返回给用户。

}

 

方法说明：

当user填完申请表格，并向运营商付款成功后，运营商通过此方法调用GenCheck来生成一张check，并返回购买收据，然后通过邮件的方式来将收据receipt发送给user。Apply参数是根据用户提交的线上或者线下申请表格生成。

 

疑问：

1.运营商是否需要存储收据？如果不存储，如何确保用户能顺利收到收据？

2.用户如何验证，合适验证收据的合法性？

### **1.4 GenCheck（辅助方法）**

// 使用user提交的购买申请来生成一张check

GenCheck(a *Apply) (*check.Check, error) {

当用户购买支票时，根据用户提供的申请信息Apply，为用户生成一张支票：先根据to值定位到对应的支票数组，然后查询到数组中的最大nonce值，然后使用nonce+1来生成新的check。如果数组为空，则新nonce设为1。

支票还需加入operator自身地址等相关信息。

最后调用store方法将生成的支票存储到本地支票池。

}

### **1.5 Check池**

type CheckPool struct {

// to -> []check

Data map[common.Address][]*check.Check

}

 

// 存储支票

func (p *CheckPool) Store(c *check.Check) error {

以nonce大小为顺序，将一张支票插入到合适的位置。

注：考虑到有可能出现user购买了多张check，但由于某种原因导致nonce小的check先存储到池里，那么后到达的nonce更大的check需要有序插入到池中。

}

 

// user可以远程调用来通过收据获取购买过的支票。

(p *CheckPool) GetCheck(r *Receipt) (*check.Check, error) { 

r为用户购买check时得到的收据信息。从收据可以得到支票的to和nonce信息，这样能够从支票池中定位一张check。

注意，需要先验证收据的合法性（验证收据是否有合法的operator签名）。

}

### **1.6 check或者paycheck的退钱机制**

当user的一张check或者paycheck需要退钱的时候，如何实现？

退款分为常规退款和异常退款

1.常规退款

某张check，用于生成paycheck支付以后还剩余了一点金额，这张paycheck需要退钱。

先根据支票的from，to，value，payvalue来在区块链的event中找到此支票的提现记录。然后根据value和payvalue的差值来确定具体该退多少钱。

2.异常退款

User的某张check，由于某种原因还没使用就过期了，无法用于支付，需要退钱。比如这张支票迟到了。

 

运营商如何知道user的某张支票到底花掉了多少？

查询链上的交易历史，查到此支票被兑现了多少钱，就可以算出应该退多少钱了。如果链上不存在此支票的交易历史，说明此支票没有被使用过，需要全额退款。

需要在每次provider提现的时候，使用event记录一个提现的事件，并记录一个topic以便以后可以根据交易信息进行查询。单独使用一套监控机制来监控每一次提现事件。

如何知道某张支票是否已经退过钱了（这个可以在check池中想办法解决）？

运营商线下保存每次退钱记录，先考虑在check池中记录每次退钱动作。

建立一个退钱池，记录跟每个节点相对应的所有已退款的支票。

### **1.7 退款池**

记录每张已退款的支票。

type RefundPool struct {

// to -> []*check

Data map[common.Address][]*check.Check

}

 

// 将退款支票记录到池中

func (p *CheckPool) Record(c *check.Check) error {

先验证支票check的签名合法性，然后在退款池中查询此支票是否已经存在，如果已经存在，则返回错误。

最后将要退款的支票，以nonce为顺序，放入到对应的支票切片中。

}

 

// 两个预退款方法，正常/异常退款，在运营商将钱退还给user之前调用，此调用成功了再将钱退还给user。如果执行失败了则不能退钱。

 

// 正常退款预处理，退还paycheck余额。这里需要支票和退款人user的地址，用于验证user的签名。Paycheck由需要退款服务的user提供。

Func (p *CheckPool) NomalRefund(pc *check.Paycheck) error {

先验证paycheck的签名合法性（user签名和operator签名）。

然后判断此支票是否已经退过款了（看池中是否已经存在）。

然后在链上查询此paycheck是否有提现记录，根据关键字to，nonce来查询。

最后调用退款池的Record方法，将要退款的支票记录到池中。

记录成功则返回成功，否则返回失败。

}

 

// 异常退款预处理，退还迟到支票的所有金额。

Func (p *CheckPool) ExceptionRefund(c *check.Check, user common.Address) error {

先验证check的签名合法性。

然后判断此支票是否已经退过款了（看池中是否已经存在）。

然后在链上查询此支票是否被提现过。使用关键字to，nonce来查询。

最后调用Record方法将它记录到退款池中。

记录成功后返回成功，否则返回失败。

}

## 二、**User功能**

### **2.1 购买支票**

用户先向运营商提交购买申请Apply，运营商审核申请通过后，等待user付款，付款完成后，operator根据提交的申请生成支票check存储到operator本地支票池，最后把收据Receipt发送给user，用户收到此收据后，将其存储到本地收据池ReceiptPool。

### **2.2 使用收据索要支票**

使用手上的购买支票的收据receipt来远程调用operator的Get方法，可以获取到自己曾经购买过的支票。

### **2.3 验证check**

用户获取到check以后，有如下验证流程：

 

// 验证check

Verify(chk *check.Check) (uint64, error) {

验证内容及错误码，0代表接收成功:

1.验证支票签名失败。

2.支票的from字段不等于user地址

3.支票的nonce没有大于合约当前nonce

4.支票在本地池中已存在

 

验证成功后，存储check到本地支票池。

}

### **2.4 GenPaycheck（Pay的辅助方法）**

// 当需要向provider支付费用时，使用check和payvalue值来生成一张paycheck

GenPaycheck(chk *check.Check, payValue *big.Int) (*check.Paycheck, error) {

使用一张check，以及payvalue来生成一张paycheck，并存储到paycheck池。

}

### **2.5 GetVirgin（Pay的辅助方法）**

// 从check池中取出处女check（尚未使用过的check）。

GetVirgin(to common.Address) (*check.Check, error){

先调用GetCurrent从paycheck池中取出current支票的nonce值。

然后在check池中找出第一张大于此nonce的check，即为未使用过的处女check，返回。

没找到则返回空，报错。

}

### **2.6 Value（计算数据块价值）**

根据数据块大小，以及价格系数来确定数据块的实际价值。

Value(size *big.Int, factor uint64) *big.Int{

return size.Mul(factor)

}

### **2.7 Pay**

注意：开始传输数据和支付支票前，user和provider需要有一个协商过程。

支付原理就是每当user收到一个数据块时，把数据块的价值blockvalue（由Value方法计算得到）累加到current支票的payvalue上来生成新的current支票。

current支票是支票池中当前被用于支付的paycheck（其nonce值最大）。

 

注：这里需要考虑，万一user收到一个异常数据块（比如数据块大小或者内容错误），如何处理。是否需要让provider对发送的数据块做一个hash？使得user可以根据此hash来验证接收到的数据块是否正确。

 

// accumulate blockvalue to generate new current paycheck

Pay(to common.Address, blockvalue *big.Int) error {

1.调用paycheck池的GetCurrent取出current支票，查看其余额是否足够支付blockvalue。

2.如果current支票余额足够，则使用累加blockvalue后的payvalue值，调用GenPaycheck生成一张新的paycheck，用其替换存储池中的current支票。

3.如果current支票的余额不够支付，需要调用check池的GetVirgin方法，获取一张已购买但尚未使用过的处女check（需要考虑获取到的处女支票无法支付怎么办？如果获取到的处女支票nonce没有大于当前支付的支票nonce，说明此处女支票还没用就已经过期了，比如它迟到了，导致被user放到支票池的时候就已经有nonce更大的支票在支付了。这种情况需要将此处女支票放入到退钱池中，以便以后向运营商退钱。同时，运营商需要有一个将paycheck的剩余金额退钱的机制）。

然后使用此check，以blockvalue作为参数调用GenPaycheck来生成一张新的paycheck替换掉paycheck池里的current支票。

如果GetVirgin返回空，说明user的支票池中已经没有可支付支票，那么提示用户需要购买check。

注意：

在GetVirgin获取到一个处女支票以后，万一还有一张迟到的具有较小nonce的新购买支票存在，则会导致这张迟到的支票过期（nonce值小于当前正在支付的paycheck），无法用于支付。考虑在check池存储支票的时候，先判断此支票是否是一张迟到的支票，再决定是否将其存入到池中。

可以考虑退钱的方案，但是这有可能会导致恶意用户反复购买退钱，因为购买支票和接收支票是分开的两个动作，运营商在售出支票给user以后，无法控制用户什么时候将购买的支票存储到池里面。用户可能会恶意的先将大nonce的支票放到池中，并花掉，这样的话前面的小nonce支票都无法使用，因而需要退钱。

}

 

疑问1：user如何知道provider是否成功接收到了paycheck?

答：这个是网络层要实现的逻辑，支付层只需要定义好支付（send）接口，然后在返回值中来判断本次发送的结果状态。

疑问2：user如何知道provider已经传输完毕，何时结束支付？

答：数据是否传输完成由user来决定，不是由provider来决定。比如user想要的文件大小为100M，当user成功接收到100M数据了，就知道本次传输结束了。如果收到的数据超过了100M，则表示接收的数据不正确，需要报错。

### **2.8 Check池**

// check以nonce为序添加到数组里

type CheckPool struct {

// to -> check

Data map[common.Address][]*check.Check // nonce有序

}

 

// 存储一张新收到的check到check池：

func (p *CheckPool) Store(c *check.Check) error {

先验证支票的签名合法性，然后判断此支票是否已经存在于池中，如果已经存在于池中，则返回错误“支票已存在”。

如果此支票还没有收到过（支票池中不存在），则判断此支票是否是一张迟到的支票（其nonce小于paycheck池的当前正用于支付的支票nonce），正常情况一张新收到的支票nonce必须大于正在支付的支票nonce，除非是这张支票迟到了，导致后收到的支票都投入支付了，才收到它。如果是迟到支票，则返回错误“支票迟到了，这张支票需要退款”。

全部验证通过以后，以nonce大小为序，将新支票放入到支票池正确位置。

}

### **2.9 Paycheck池**

type PaycheckPool struct {

Pool map[common.Address] []*check.Paycheck 	//按照nonce有序

}

 

// 存储paycheck到池中，每个to一个数组

func (p *PaycheckPool) Store(pc *check.Paycheck) error {

如果paycheck池为空，则直接append到池的末尾并返回。

用新paycheck的nonce跟current支票的nonce相比较。

如果nonce相等，并且payvalue值更大，则直接替换current支票，否则忽略掉，然后返回。

如果nonce大于current支票，表示这是一张新check，则直接添加到池末尾，然后返回。

其余情况，报错返回。

}

 

疑问：

1.是否允许user同时向provider请求多个不同的文件？

允许。

2.如果同时向provider请求多个文件，那么在支付这些文件的时候是使用不同的支票进行支付，还是使用同一个支票支付所有的文件？

注：如果使用相同的支票支付所有请求，那么一旦这张支票余额不够了，会导致所有支票支付失败。而且在paycheck的角度来看，也无法区分不同的文件请求，它只知道接收到了一个数据块，但不知道这个数据块属于哪个文件。所以这里还是使用同一张支票来支付所有文件请求。

 

// 从paycheck池中取出current支票

func (p *PaycheckPool) GetCurrent(to common.Address) (*check.Paycheck, error){

如果池为空则返回空，说明没有任何paycheck。

返回current支票（nonce值最大的那个paycheck，也即切片末尾的那个paycheck）。

}

### **2.10 Receipt池**

type Receipt struct {

Dt Date 				// 购买日期

Value *big.Int			// 购买金额

Token common.Address 	// 货币类型

Op common.Address	// 运营商地址

From common.Address	// 付款方地址

To common.Address	// 收款方地址

Nonce uint64			//  nonce

...

Sig					// 运营商的签名

}

 

type ReceiptPool struct {

Pool []*Receipt

}

 

// 存储一张收据到收据池

func (p *ReceiptPool) Store(r *Receipt) error{

p.Pool.append(r)

}

注意：需要考虑先验证收据的合法性，然后再存储到池中。

 

收据池信息的查询，显示，筛选，选择，以及返回，需要另外写一套功能来实现。

## 三、**Provider功能**

### **3.****1** **CalcPay（Verify的辅助方法）**

// 计算收到的paycheck实际支付金额

CalcPay(pchk *check.Paycheck) (*big.Int, error) {

如果nonce值等于支票池尾项的nonce，则直接计算他们payvalue的差值作为实际金额返回（支付金额必须为正值，否则报错）。

如果nonce值大于尾项的nonce，说明这是使用了一张新支票进行支付，直接返回支票的payvalue值作为实际支付金额。

如果nonce小于尾项的nonce，说明这是使用了一张已过期的支票尝试进行支付，返回错误。

 

疑问：

万一出现支票池数据丢失的情况怎么办？支票池没数据就无法正确计算支付金额了。

### **3.2 Verify**

// 当provider接收到一张paycheck时调用，用于验证接收到的paycheck是否合法，blockvalue为当前数据块的价值金额，由blockvalue函数根据数据块大小计算出来

Verify(pchk *check.Paycheck, blockvalue *big.Int) (uint64, error) {

 

返回值为错误码，0代表接收成功。

错误码：

1.支票的value小于payvalue

2.支票的nonce值没有大于合约的当前nonce（因为过期的nonce无法兑换了）

3.支票的to不等于provider地址

4.收到的paycheck的nonce，没有大于最近提现支票的nonce（这里命名为rejectNonce，由提现方法每次提现的时候更新，应对支票提现尚未完成，合约nonce还没更新，但支票实际上已经过期的情况）。

5.支票的实际支付金额，不等于数据块的价值blockvalue

调用CalcPay方法计算支票的实际支付金额。

 

以上验证全部通过以后，调用paycheck池的store方法将支票存储到池中。

}

### **3.3 SendTx（Withdraw的辅助方法）**

(pro *Provider) SendTx(pc *check.Paycheck) (tx *types.Transaction, err error) {

使用paycheck为参数，向链发送提现交易。

### **3.4 Withdraw**

// 从支票池中取出能提现的支票中nonce最小的那个paycheck（注意先查看支票的当前使用比率）进行提现。

Withdraw() (retCode uint64, e error) {

 

调用paycheck池的GetNextPayable()方法，找到下一个能提现的paycheck（其nonce值刚好大于合约中的对应nonce）。

如果没找到，则返回错误1，表示当前没有可用于提现的paycheck。

如果找到了，先使用此paycheck的nonce更新rejectNonce。

然后使用此paycheck作为参数，调用SendTx向链发送提现交易，以获取收益，并更新合约中的nonce值。

 

疑问：合约交易发送成功了是否能保证此交易一定能上链？

答：不能保证。但是不影响提现流程，因为如果一个paycheck没有上链成功，那么合约的nonce值就没有被改变，在下次提现的时候，还是会把这个paycheck再次找出来（依据合约的nonce来查找），并再次使用它发出提现交易。

}

### **3.5 Paycheck池**

type PaycheckPool struct {

Pool []*check.Paycheck 	//按照nonce有序

}

 

// 存储一张paycheck到池中

(p *PaycheckPool ) Store(pc *check.Paycheck) error {

按照nonce的大小，有序插入到数组（理论上是最新的paycheck的nonce值必须为最大，并且直接append到切片的末尾）。

}

 

// 找出下一个能提现的paycheck，如果找到了则返回它，如果没找到则返回空

(p *PaycheckPool ) GetNextPayable() (*check.Paycheck, error) {

先查看合约中节点对应的nonce，然后在本地paycheck池中找出第一个比它大的支票，找到了就返回，否则就是池中已经没有可提现的paycheck了。

}

 