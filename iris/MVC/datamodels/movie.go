package datamodels

/*
	Movie 是我们的样本数据结构。请注意，公共用途的标签（对于我们的网络应用程序）应保存在其他文件中，例如“web / viewmodels / movie.go”，可以通过嵌入datamodels.Movie或声明新字段来换行但我们将使用此数据模型作为我们应用中唯一的 Movie model，为了简单起见(for the shake of simplicty)。
*/
type Movie struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:poster`
}
