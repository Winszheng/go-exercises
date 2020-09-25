package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/Winszheng/go-grpc-http-rest-microservice/pkg/protocol/grpc"
	v1 "github.com/Winszheng/go-grpc-http-rest-microservice/pkg/service/v1"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//服务器的基本配置
type Config struct {
	GrpcPort string
	DatabaseHost string
	DatabaseUser string
	DatabasePassword string
	DatabaseSchema string
}

//RunServer负责读取命令行输入的参数, 创建数据库连接, 创建ToDo服务实例以及调用之前grpc服务中的RunServer函数
func RunServer() error {
	ctx := context.Background()

	//1.命令行参数解析
	//(1)注册命令行参数并绑定到变量中
	var config Config
	flag.StringVar(&config.GrpcPort, "grpc-port","", "grpc port to bind")
	flag.StringVar(&config.DatabaseHost,"db-host","", "database host")
	flag.StringVar(&config.DatabaseUser,"db-user","", "database user")
	flag.StringVar(&config.DatabasePassword,"db-password","", "database password")
	flag.StringVar(&config.DatabaseSchema,"db-schema", "", "database schema")
	//(2)解析命令行参数到注册的flag当中, 至此，可以使用相应的指针或变量(值)
	flag.Parse()
	if len(config.GrpcPort) == 0 {
		return fmt.Errorf("invalid tcp port for grpc server: %s", config.GrpcPort)
	}

	//2.创建数据库连接
	// add MySQL driver specific parameter to parse date/time
	param := "parseTime=True"

	//连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseHost,
		config.DatabaseSchema,
		param)
	db, err := sql.Open("mysql", dsn)   //打开数据库连接
	if err!= nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	//3.创建todo服务实例并调用之前grpc服务中的Runerver函数
	//不太理解为什么需要两个server.go
	v1API := v1.NewToDoServiceServer(db)
	return grpc.RunServer(ctx, v1API, config.GrpcPort)

}
