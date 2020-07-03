package strcode

//AllProductData for result.html
type AllProductData struct {
	ListProduct  []ProductData
	UsernameInfo UsernameInfo
}

// ProductData for result.html
type ProductData struct {
	No           int
	Pid          string
	Powner       string
	Tittle       string
	Pname        string
	Pprice       int
	Pamount      int
	Pquality     string
	Pcategory    string
	Ptime        string
	PCreateDate  string
	PLastUpdate  string
	UsernameInfo UsernameInfo
}

//BuyProduct for buysomeproduct
type BuyProduct struct {
	InPid      string `json:"pid"`
	InTittle   string `json:"tittle"`
	InName     string `json:"name"`
	InAmount   string `json:"amount_buy"`
	InTotalPay string `json:"total_pay"`
	InCategory string `json:"category_in"`
	InQuality  string `json:"quality_in"`
	InOwner    string `json:"owner_in"`
}

//EditProduct for buysomeproduct
type EditProduct struct {
	InPid      string `json:"pid"`
	InTittle   string `json:"tittle"`
	InName     string `json:"name"`
	InAmount   string `json:"amount"`
	InPrice    string `json:"price"`
	InCategory string `json:"category"`
	InQuality  string `json:"quality"`
}

//UsernameInfo Data
type UsernameInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Response Data
type Response struct {
	ErrorCode string `json:"ErrorCode"`
}

//Content Data
type Content struct {
	Message string `json:"Message"`
}

//UserLoginInfo Data
type UserLoginInfo struct {
	Username string
	Password string
	Message  string
}

//HistTrx Data
type HistTrx struct {
	Data  BuyProduct
	Owner string
	Buyer string
}

//TrxProduct for buysomeproduct
type TrxProduct struct {
	No         int
	Owner      string
	Buyer      string
	Product    string
	Amount     string
	Pay        string
	CreateDate string
}

//HistPage page
type HistPage struct {
	TrxProduct   []TrxProduct
	UsernameInfo UsernameInfo
}

//ChatProduct for buysomeproduct
type ChatProduct struct {
	No         int    `json:"No"`
	ChatID     string `json:"ChatID"`
	User1      string `json:"User1"`
	User2      string `json:"User2"`
	LastUpdate string `json:"LastUpdate"`
	Payload    string `json:"Payload"`
	Peek       string `json:"Peek"`
}

//ChatIdle page
type ChatIdle struct {
	ChatProduct  []ChatProduct
	UsernameInfo UsernameInfo
}
