package orm

import (
	"fmt"
	"inn/internal/user/model"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var db *gorm.DB

//Init 根据配置初始化gorm 打开数据库连接
func Init() {
	var err error
	var URL string

	var (
		driver     = viper.GetString("DB.DRIVER")
		DBUser     = viper.GetString("DB.USER")
		DBPassword = viper.GetString("DB.PASSWORD")
		DBName     = viper.GetString("DB.NAME")
		DBHost     = viper.GetString("DB.HOST")
		DBPort     = viper.GetString("DB.PORT")
	)

	fmt.Println("DBHost")

	if driver == "mysql" {
		URL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	}
	if driver == "postgres" {
		URL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	}

	fmt.Println(URL)

	db, err = gorm.Open(driver, URL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", driver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", driver)
	}
	db.SingularTable(true)

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "" + defaultTableName
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.AutoMigrate(&model.User{})
}

// GetDB 获取gorm对象
func GetDB() *gorm.DB {
	return db
}

// Close 关闭数据库连接
func Close() {
	if db != nil {
		db.Close()
	}
}
