package model

type Todo struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`    
	
	/** Để convert instance của struct User sang JSON chúng ta sử dụng 
	go package encoding/json 
	hàm json.Marshal() return về mảng kiểu byte và lỗi ( []byte , err)
	*/
}

