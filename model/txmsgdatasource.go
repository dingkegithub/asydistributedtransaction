package model

import (
	"strings"
)

type TxDataMsgSource struct {
	Url      string `json:"url"`
	UserName string `json:"user_name"`
	Password string `json:"user_name"`
}

func (tdms *TxDataMsgSource) String() string {
	builder := strings.Builder{}
	builder.WriteString("TxDataMsgSource [url=")
	builder.WriteString(tdms.Url)
	builder.WriteString(", username=")
	builder.WriteString(tdms.UserName)
	builder.WriteString(", password=")
	builder.WriteString(tdms.Password)
	return builder.String()
}
