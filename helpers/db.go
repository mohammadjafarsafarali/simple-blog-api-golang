package helpers

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "root"
	dbname   = "blog_go"
)

/*func InitDB() {

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}


}*/
