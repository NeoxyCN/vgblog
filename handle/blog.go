package handle

const (
	DB_TABLE_BLOG = "blog_basic"
)

type blog_basic struct {
	Id       int    `db:"id"`
	Title    string `db:"title"`
	Username string `db:"username"`
	Date     string `db:"date"`
	Content  string `db:"content"`
	Comment  string `db:"comment"`
	Tag      string `db:"tag"`
	Category string `db:"category"`
}
