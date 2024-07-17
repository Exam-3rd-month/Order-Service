package test

import (
	"context"
	"log"
	"log/slog"
	"order-service/genprotos/auth_pb"
	"order-service/genprotos/order_pb"
	"order-service/internal/config"
	"order-service/internal/storage"
	"os"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type testConfig struct {
	auth_pb.AuthServiceClient
}

func connect(port string) *grpc.ClientConn {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

func NewTestConfig() *testConfig {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	return &testConfig{
		auth_pb.NewAuthServiceClient(connect(config.Server.Auth_Port)),
	}
}

func setupAuthSt() *storage.OrderSt {
	logFile, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	configs, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	storage, err := storage.New(configs, logger)
	if err != nil {
		log.Fatal(err)
	}

	return storage
}

func TestAddDish(t *testing.T) {
	OrderSt := setupAuthSt()
	ctx := context.Background()
	test := auth_pb.RegisterRequest{
		Username: "testusername",
		Email:    "testemail",
		Password: "testpassword",
		FullName: "testfullname",
		UserType: "user",
	}

	auth_client := NewTestConfig()

	resp, err := auth_client.Register(ctx, &test)
	if err != nil {
		t.Error("faliled to register user", resp)
	}

	test2 := auth_pb.CreateKitchenRequest{
		OwnerId:     resp.UserId,
		Name:        "testname",
		Description: "testdescription",
		CuisineType: "testcuisinetype",
		Address:     "testaddress",
		PhoneNumber: "testphonenumber",
	}

	resp2, err := auth_client.CreateKitchen(ctx, &test2)
	if err != nil {
		t.Error("faliled to create kitchen", resp2)
	}

	test3 := order_pb.AddDishRequest{
		KitchenId:   resp2.KitchenId,
		Name:        "testname",
		Description: "testdescription",
		Price:       100,
		Category:    "testcategory",
		Ingredients: []string{"testingredients"},
		Available:   true,
	}

	resp3, err := OrderSt.AddDish(ctx, &test3)
	if err != nil {
		t.Error("faliled to add dish", resp3)
	}

	_, err = auth_client.DeleteKitchen(ctx, &auth_pb.DeleteKitchenRequest{KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = auth_client.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = OrderSt.DeleteDish(ctx, &order_pb.DeleteDishRequest{DishId: resp3.DishId, KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete dish", resp3)
	}
}

func TestUpdateDish(t *testing.T) {
	OrderSt := setupAuthSt()
	ctx := context.Background()
	test := auth_pb.RegisterRequest{
		Username: "testusername",
		Email:    "testemail",
		Password: "testpassword",
		FullName: "testfullname",
		UserType: "user",
	}

	auth_client := NewTestConfig()

	resp, err := auth_client.Register(ctx, &test)
	if err != nil {
		t.Error("faliled to register user", resp)
	}

	test2 := auth_pb.CreateKitchenRequest{
		OwnerId:     resp.UserId,
		Name:        "testname",
		Description: "testdescription",
		CuisineType: "testcuisinetype",
		Address:     "testaddress",
		PhoneNumber: "testphonenumber",
	}

	resp2, err := auth_client.CreateKitchen(ctx, &test2)
	if err != nil {
		t.Error("faliled to create kitchen", resp2)
	}

	test3 := order_pb.AddDishRequest{
		KitchenId:   resp2.KitchenId,
		Name:        "testname",
		Description: "testdescription",
		Price:       100,
		Category:    "testcategory",
		Ingredients: []string{"testingredients"},
		Available:   true,
	}

	resp3, err := OrderSt.AddDish(ctx, &test3)
	if err != nil {
		t.Error("faliled to add dish", resp3)
	}

	test4 := order_pb.UpdateDishRequest{
		DishId:      resp3.DishId,
		KitchenId:   resp2.KitchenId,
		Name:        "testname",
		Description: "testdescription",
		Price:       100,
		Category:    "testcategory",
		Ingredients: []string{"testingredients"},
		Available:   true,
	}

	resp4, err := OrderSt.UpdateDish(ctx, &test4)
	if err != nil {
		t.Error("faliled to update dish", resp4)
	}

	_, err = auth_client.DeleteKitchen(ctx, &auth_pb.DeleteKitchenRequest{KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = auth_client.DeleteUser(ctx, &auth_pb.DeleteUserRequest{UserId: resp.UserId})
	if err != nil {
		t.Error("failed to delete user", resp)
	}

	_, err = OrderSt.DeleteDish(ctx, &order_pb.DeleteDishRequest{DishId: resp3.DishId, KitchenId: resp2.KitchenId})
	if err != nil {
		t.Error("failed to delete dish", resp3)
	}
}
