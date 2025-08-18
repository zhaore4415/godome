// database/db.go
package database

import (
	"context"
	"time"
	"web03/config"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局变量改为 *gorm.DB

// InitDB 初始化数据库连接 (使用 GORM)
func InitDB() error {
	cfg := config.GetDBConfig()

	// 构建 DSN (Data Source Name)
	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + cfg.PortStr() + ")/" + cfg.Name + "?charset=utf8mb4&parseTime=true&loc=Local"

	var err error
	// 使用 GORM Open
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 获取底层的 *sql.DB 以进行连接池设置
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// 设置连接池 (与原生 sql.DB 设置方式相同)
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	// GORM 的 DB 对象也有 Ping 方法，但为了明确，也可以通过 sqlDB Ping
	// 这里使用 sqlDB.PingContext 与您之前的代码保持一致
	if err = sqlDB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if DB != nil {
		sqlDB, _ := DB.DB() // 获取底层 *sql.DB
		sqlDB.Close()       // 关闭底层连接
	}
}

// var DB *sql.DB   不用gorm

// // InitDB 初始化数据库连接
// func InitDB() error {
// 	cfg := config.GetDBConfig()

// 	// 构建 DSN (Data Source Name)
// 	dsn := cfg.User + ":" + cfg.Password + "@tcp(" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + ")/" + cfg.Name + "?parseTime=true&loc=Local"

// 	var err error
// 	DB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return err
// 	}

// 	// 设置连接池
// 	DB.SetMaxOpenConns(25)
// 	DB.SetMaxIdleConns(25)
// 	DB.SetConnMaxLifetime(5 * time.Minute)

// 	// 测试连接
// 	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
// 	defer cancel()

// 	if err = DB.PingContext(ctx); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // CloseDB 关闭数据库连接
// func CloseDB() {
// 	if DB != nil {
// 		DB.Close()
// 	}
// }
