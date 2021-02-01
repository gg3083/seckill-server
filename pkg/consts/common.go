package consts

const (
	_ = iota
	BusinessUserTable
	BusinessGoodsTable
	BusinessOrderTable
	BusinessUserAddressTable
	BusinessUserFundRecordTable
	BusinessUserFundTable
)

const PasswordSalt = "PasswordSalt"

//常量
const (
	//gin的上下文的token_key
	TokenKey     = "claims"
	LogResponse  = "log_response"
	OperationLog = "operation_log"
	HeaderToken  = "token"
)

//充值来源
const (
	_ = iota
	ChargeForSystem
	ChargeForAlipay
	ChargeForWechat
)

//充值来源
const (
	_ = iota
	AmountAdd
	AmountSub
)
