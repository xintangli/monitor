package params

import (
	derrors "github.com/xintangli/monitor/error"
	"time"
)

type OuterBusiReq struct {
	MemId	string `form:"men_id"`
	OrderId string `form:"order_id"`
	StartTime string `form:"start_time"`
	EndTime string `form:"end_time"`
	Timestamp string `form:"timestamp"`

	Sign string `form:"sign"`
}

func (r *OuterBusiReq) ParamVerify() error {
	if r.MemId == "" {
		return derrors.New("1099", "MemId is null")
	}
	if r.OrderId == "" {
		return derrors.New("1099", "OrderId is null")
	}
	if r.StartTime == "" {
		return derrors.New("1099", "StartTime is null")
	}
	if r.EndTime == "" {
		return derrors.New("1099", "EndTime is null")
	}
	if r.Sign == "" {
		return derrors.New("1099", "Sign is null")
	}
	start, _ := time.Parse("2006-01-02 03:04",r.StartTime)
	end, _ := time.Parse("2006-01-02 03:04",r.EndTime)

	interval := end.Unix() - start.Unix()

	if interval > 10 * 60 {
		return derrors.New("1099", "interval too long")
	}

	return nil
}
