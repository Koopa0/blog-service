package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "取得標籤列表失敗")
	ErrorCreateTagFail  = NewError(20010002, "建立標籤失敗")
	ErrorUpdateTagFail  = NewError(20010003, "更新標籤失敗")
	ErrorDeleteTagFail  = NewError(20010004, "刪除標籤失敗")
	ErrorCountTagFail   = NewError(20010005, "統計標籤失敗")
)
