package dal

import (
	"ByteDance/dal/query"
	"ByteDance/pkg/common"
	"ByteDance/utils"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var ConnQuery *query.Query
var once sync.Once

// 初始化，将ConnQuery与数据库绑定
func init() {
	once.Do(func() {
		ConnQuery = getQueryConnection()
		//启动redis
		err := getQueryConnectionRedis()
		if err != nil {
			//redis连接错误
			panic(err)
		}
	})
}

func getQueryConnection() *query.Query {
	var err error
	var db *gorm.DB
	db, err = gorm.Open(mysql.Open(common.MySqlDSN))
	if err != nil {
		utils.CatchErr("数据库连接错误", err)
	}
	ConnQuery = query.Use(db)
	return ConnQuery
}

// 连接到redis
func getQueryConnectionRedis() (err error) {
	common.RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = common.RedisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
