package routes

type lineModel struct {
	Day  string `json:"day"`  // 日付(yyyy-mm-dd形式)
	Memo string `json:"memo"` // 説明
}
