package config

type Config struct {
	DelPeriod           uint64   `json:"del_period"`
	DelMsgOnTime        uint64   `json:"del_msg_on_time"`
	MinIdleConn         int      `json:"min_idle_conn"`
	MaxConn             int      `json:"max_conn"`
	SendMsgTimeout      uint64   `json:"send_msg_timeout"`
	SchedScanTimePeriod uint64   `json:"sched_scan_time_period"`
	MaxWaitTime         uint64   `json:"max_wait_time"`
	CloseWaitTime       uint64   `json:"close_wait_time"`
	StatsTimePeriod     uint64   `json:"stats_time_period"`
	HistoryMsgStoreTime uint64   `json:"history_msg_store_time"`
	MsgTableName        string   `json:"msg_table_name"`
	EtcdHosts           []string `json:"etcd_hosts"`
	ServiceName         string   `json:"service_name"`
}
