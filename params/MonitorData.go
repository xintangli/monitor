package params


type MonitorParams struct {
	Msg struct {
		SERINFO struct {
			RedisUsability bool `json:"redisUsability"`
			QPSSucc string `json:"qpsSucc"`
			QPSFail string `json:"qpsFail"`
			RedisMem string `json:"redisMem"`
			ProcessNum string `json:"processNum"`
			RedisSize int `json:"redisSize"`
			} `json:"SERINFO"`
		SYSINFO struct {
			Load struct {
				Lavg5 string `json:"lavg_5"`
				LastPid string `json:"last_pid"`
				Lavg15 string `json:"lavg_15"`
				Lavg1 string `json:"lavg_1"`
				Nr string `json:"nr"`
			     } `json:"Load"`
			Mem struct {
				Cached float64 `json:"Cached"`
				MemFree float64 `json:"MemFree"`
				MemTotal float64 `json:"MemTotal"`
				MemUsed float64 `json:"MemUsed"`
				Buffers float64 `json:"Buffers"`
			    } `json:"Mem"`
			MSGWARN struct {} `json:"MSGWARN"`
				Disk []struct {
					Use string `json:"Use"`
					Used string `json:"Used"`
					Avail string `json:"Avail"`
					Filesystem string `json:"Filesystem"`
					Mounted string `json:"Mounted"`
					Size string `json:"Size"`
				    } `json:"Disk"`
			CPU []struct {
				CPU float64 `json:"cpu,omitempty"`
				CPU0 float64 `json:"cpu0,omitempty"`
				CPU1 float64 `json:"cpu1,omitempty"`
				CPU2 float64 `json:"cpu2,omitempty"`
				CPU3 float64 `json:"cpu3,omitempty"`
				CPU4 float64 `json:"cpu4,omitempty"`
				CPU5 float64 `json:"cpu5,omitempty"`
				CPU6 float64 `json:"cpu6,omitempty"`
				CPU7 float64 `json:"cpu7,omitempty"`
				CPU8 float64 `json:"cpu8,omitempty"`
				CPU9 float64 `json:"cpu9,omitempty"`
				CPU10 float64 `json:"cpu10,omitempty"`
				CPU11 float64 `json:"cpu11,omitempty"`
				CPU12 float64 `json:"cpu12,omitempty"`
				CPU13 float64 `json:"cpu13,omitempty"`
				CPU14 float64 `json:"cpu14,omitempty"`
				CPU15 float64 `json:"cpu15,omitempty"`
				CPU16 float64 `json:"cpu16,omitempty"`
				CPU17 float64 `json:"cpu17,omitempty"`
				CPU18 float64 `json:"cpu18,omitempty"`
				CPU19 float64 `json:"cpu19,omitempty"`
				CPU20 float64 `json:"cpu20,omitempty"`
				CPU21 float64 `json:"cpu21,omitempty"`
				CPU22 float64 `json:"cpu22,omitempty"`
				CPU23 float64 `json:"cpu23,omitempty"`
				CPU24 float64 `json:"cpu24,omitempty"`
				CPU25 float64 `json:"cpu25,omitempty"`
				CPU26 float64 `json:"cpu26,omitempty"`
				CPU27 float64 `json:"cpu27,omitempty"`
				CPU28 float64 `json:"cpu28,omitempty"`
				CPU29 float64 `json:"cpu29,omitempty"`
				CPU30 float64 `json:"cpu30,omitempty"`
				CPU31 float64 `json:"cpu31,omitempty"`
			} `json:"CPU"`
			MSGERR struct {} `json:"MSGERR"`
			} `json:"SYSINFO"`
		DEMINFO struct {
			RequestOk []int `json:"requestOk"`
			RequestFail []int `json:"requestFail"`
			OrderIds []string `json:"orderIds"`
			} `json:"DEMINFO"`
		SUPINFO struct {
			MSGWARN struct {} `json:"MSGWARN"`
			SupData struct {
				Code []string `json:"code"`
				CodeCount []int `json:"codeCount"`
				QueryTotal int `json:"queryTotal"`
				QuerySuccess struct{
					TaskID []string `json:"taskID"`
					Count []int `json:"count"`
					     } `json:"querySuccess"`
				} `json:"supData"`
			MSGERR struct {} `json:"MSGERR"`
			} `json:"SUPINFO"`
	    } `json:"msg"`
	MemID string `json:"memId"`
	NodeID string `json:"nodeId"`
	MoniTarget string `json:"moniTarget"`
	WarnLevel string `json:"warnLevel"`
	CurTime string `json:"curTime"`
}