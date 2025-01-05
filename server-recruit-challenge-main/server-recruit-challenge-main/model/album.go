package model


type AlbumID int
//課題4
type Album struct {
    ID       AlbumID `json:"id"`
    Title    string  `json:"title"`
    SingerID int `json:"singer_id"`  // 外部キー（SingerのID）
    Singer   *Singer `json:"singer"`     // 曲の歌手情報（Singerの詳細）
}
//課題3
//type Album struct {
//	ID       AlbumID  `json:"id"`
//	Title    string   `json:"title"`
//	Singer	 Singer	  `json:"singer"`
//	SingerID SingerID `json:"singer_id"` // モデル Singer の ID と紐づきます
//}



func (a *Album) Validate() error {
	if a.Title == "" {
		return ErrInvalidParam
	}
	if len(a.Title) > 255 {
		return ErrInvalidParam
	}
	return nil
}
