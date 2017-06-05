package params

type DTPageParam struct {
	SEcho int			`json:"sEcho"`
	ITotalRecords int64		`json:"iTotalRecords"`
	ITotalDisplayRecords int64	`json:"iTotalDisplayRecords"`
	AAData	[]interface{}		`json:"aaData"`
}
