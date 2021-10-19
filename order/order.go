package order

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Order struct {
	ID    uint64         // 订单ID
	Value *big.Int       // 货币数量
	Token common.Address // 货币类型
	Fee   uint64         // 应付金额
	From  common.Address // user地址
	To    common.Address // provider地址
	Time  time.Time      // 订单提交时间
	Name  string         // 购买人姓名
	Tel   string         // 购买人联系方式
	Email string         // 接收支票的邮件地址
	Paid  bool           // 标记是否已付款

	Sig string // 运营商的签名
}

type OrderPool struct {
	// user -> \[\]*Order
	Data map[common.Address][]*Order
}

func (pool *OrderPool) Store(o *Order) error {
	//将一张订单存储到每个user各自的队列下面，以订单ID为排列顺序。
	return nil
}

func (pool *OrderPool) Get(user common.Address, ID uint64) (*Order, error) {
	//根据user地址和订单ID来从订单池获取一张订单
	return nil, nil
}

func (pool *OrderPool) Pay(user common.Address, ID uint64) (*Order, error) {
	//在用户支付完成后，使用订单的user和ID信息来调用，以修改此订单的支付状态paid为true。
	return nil, nil
}
