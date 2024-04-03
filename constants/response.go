package constants

const (
	ErrorTypeSystem    string = "SYS"
	ErrorTypeCommon    string = "COM"
	ErrorTypeKyc       string = "KYC"
	ErrorTypeCif       string = "CIF"
	ErrorTypeSme       string = "SME"
	ErrorTypeAccount   string = "ACC"
	ErrorTypePayment   string = "PAY"
	ErrorTypeReconcile string = "REC"
	ErrorTypeLedger    string = "GLE"
	ErrorTypeBank      string = "BNK"
)

const (
	SuccessGetData = iota + 1
	SuccessCreated
	SuccessUpdated
	SuccessDeleted
)
