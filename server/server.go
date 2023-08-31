package main

import (
	"context"
	"fmt"
	"net"

	pro "github.com/SELVAKANNAN-P/Customer"
	services "github.com/SELVAKANNAN-P/Netxd_Dal/service"

	controllers "github.com/SELVAKANNAN-P/Netxd_Customer_server/controller"
	"github.com/SELVAKANNAN-P/Netxdconfig/config"
	"github.com/SELVAKANNAN-P/Netxdconfig/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, "Bank", "profile")
	controllers.CustomerService = services.InitCustomerService(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)

		initDatabase(mongoclient)
		lis, err := net.Listen("tcp", constants.Port)
		if err != nil {
			fmt.Printf("Failed to listen: %v", err)
			return
		}
		s := grpc.NewServer()
		//pro.RegisterCustomerServiceServer(s, &controllers.RPCserver{})
		pro.RegisterCustomerServiceServer(s, &controllers.RPCserver{})

		fmt.Println("Server listening on", constants.Port)
		if err := s.Serve(lis); err != nil {
			fmt.Printf("Failed to serve: %v", err)
		}
	}
}
