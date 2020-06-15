package strcode

//AllProductData for result.html
type AllProductData struct {
	ListProduct []ProductData
}

// ProductData for result.html
type ProductData struct {
	Pid       string
	Tittle    string
	Pname     string
	Pprice    int
	Pamount   int
	Pquality  string
	Pcategory string
	Ptime     string
}
