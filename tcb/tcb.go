package tcb

import "github.com/yellbuy/wechat/context"

//Tcb Tencent Cloud Base
type Tcb struct{
	*context.Context
}

//NewTcb new Tencent Cloud Base
func NewTcb(context *context.Context)*Tcb{
	return &Tcb{
		context,
	}
}

