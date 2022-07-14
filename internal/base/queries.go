package base

const (
	insertLog     = "insertLog"
	insertReceipt = "insertReceipt"
)

var queries = map[string]string{
	insertLog:     "insert into public.logs (date, op_type, message) values ($1,$2,$3)",
	insertReceipt: "insert into public.receipts (date, op_type, receipt) values ($1,$2,$3)",
}
