package base

const (
	insertLog     = "insertLog"
	insertReceipt = "insertReceipt"
)

var queries = map[string]string{
	insertLog:     "insert into public.logs (date, message_type, message) values ($1,$2,$3)",
	insertReceipt: "insert into public.receipts (date, message_type, message) values ($1,$2,$3)",
}
