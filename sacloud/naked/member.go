package naked

// Member 会員情報
type Member struct {
	Class string `json:",omitempty"` // クラス
	Code  string `json:",omitempty"` // 会員コード
}
