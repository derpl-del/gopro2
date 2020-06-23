package strcode

//AllProductData for result.html
type AllProductData struct {
	ListProduct  []ProductData
	UsernameInfo UsernameInfo
}

// ProductData for result.html
type ProductData struct {
	Pid          string
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
}

//UsernameInfo Data
type UsernameInfo struct {
	Username string
}
