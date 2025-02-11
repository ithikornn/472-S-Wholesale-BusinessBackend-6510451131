package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ppwlsw/sa-project-backend/adapters/api"
	"github.com/ppwlsw/sa-project-backend/adapters/database"
	"github.com/ppwlsw/sa-project-backend/usecases"
	"gorm.io/gorm"
)

func SetUpRouters(app *fiber.App, db *gorm.DB) {

	userRepo := database.InitiateUserPostgresRepository(db)
	userService := usecases.InitiateUserService(userRepo)
	userHandler := api.InitiateUserHandler(userService)

	productRepo := database.InitiateProductPostGresRepository(db)
	productService := usecases.InitiateProductsService(productRepo)
	productHandler := api.InitiateProductHandler(productService)

	transactionRepo := database.InitiateTransactionPostGresRepository(db)
	transactionService := usecases.InitiateTransactionService(transactionRepo)
	transactionHandler := api.InitiateTransactionHandler(transactionService)

	authService := usecases.InitiateAuthService(userRepo)
	authHandler := api.InitiateAuthHandler(authService)

	adminService := usecases.InitiateAdminService(userRepo)
	adminHandler := api.InitiateAdminHandler(adminService)
	adminHandler.InitializeAdmin()

	orderRepo := database.InitiateOrderPostgresRepository(db)
	orderService := usecases.InitiateOrderService(orderRepo)
	orderHandler := api.InitiateOrderHandler(orderService)

	orderLineRepo := database.InitiateOrderLinePostgresRepository(db)
	orderLineService := usecases.InitiateOrderLineService(orderLineRepo)
	orderLineHandler := api.InitiateOrderLineHandler(orderLineService)

	tierListRepo := database.InitiateTierListPostgres(db)
	tierListService := usecases.InitiateTierListService(tierListRepo)
	tierListHandler := api.InitiateTierListHandler(tierListService)
	tierListHandler.TierListUsecase.InitialTierList()

	supplierRepo := database.InitiateSupplierPostgresRepository(db)
	supplierService := usecases.InitiateSupplierService(supplierRepo)
	supplierHandler := api.InitiateSupplierHandler(supplierService)

	supplierOrderListRepo := database.InitiateSupplierOrderListPostgresRepository(db)
	supplierOrderListService := usecases.InitiateSupplierOrderListService(supplierOrderListRepo)
	supplierOrderListHandler := api.InitiateSupplierOrderListHandler(supplierOrderListService)

	handlers := api.ProvideHandlers(
		userHandler, productHandler, transactionHandler,
		authHandler, orderHandler,
		orderLineHandler, supplierHandler,
		supplierOrderListHandler, tierListHandler, adminHandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//User
	app.Get("/users", handlers.UserHandler.GetAllUsers)
	app.Get("/users/:id", handlers.UserHandler.GetUserByID)
	app.Put("/users/update", handlers.UserHandler.UpdateTierByUserID)

	//Tier
	app.Get("/discount/:id", handlers.TierListHandler.GetDiscountPercentByUserID)
	app.Post("/tierlist", handlers.TierListHandler.CreateTierList)
	app.Get("/tierlist", handlers.TierListHandler.GetAllTierList)

	//Product
	app.Post("/product", handlers.ProductHandler.CreateProduct)
	app.Post("/products", handlers.ProductHandler.CreateProducts)
	app.Post("/products/filter", handlers.ProductHandler.GetProductByFilter)
	app.Get("/products", handlers.ProductHandler.GetAllProducts)
	app.Get("product/:id", handlers.ProductHandler.GetProductByID)
	app.Put("/product/buy", handlers.ProductHandler.BuyProduct)
	app.Put("/products/buy", handlers.ProductHandler.BuyProducts)
	app.Put("products/:id", handlers.ProductHandler.UpdateProduct)

	//Transaction
	app.Post("/transaction", handlers.TransactionHandler.CreateTransaction)
	app.Post("/transactions", handlers.TransactionHandler.CreateTransaction)
	app.Get("/transactions", handlers.TransactionHandler.GetAllTransactions)
	app.Get("/transaction/:id", handlers.TransactionHandler.GetTransactionById)
	app.Get("/transaction/order/:order_id", handlers.TransactionHandler.GetTransactionByOrderId)
	app.Put("/transaction/:id", handlers.TransactionHandler.UpdateTransaction)
	app.Delete("/transaction/:id", handlers.TransactionHandler.DeleteTransaction)

	//Auth
	app.Post("/register", handlers.AuthHandler.Register)
	app.Post("/login", handlers.AuthHandler.Login)

	//Order
	app.Post("/order", handlers.OrderHandler.CreateOrder)
	app.Post("/orders", handlers.OrderHandler.CreateOrder)
	app.Get("/orders", handlers.OrderHandler.GetAllOrders)
	app.Get("/order/:id", handlers.OrderHandler.GetOrderByID)
	app.Get("/order/user/:id", handlers.OrderHandler.GetOrderByUserID)
	app.Get("/order/user/detail/:id", handlers.OrderHandler.GetOrderAndUserByID)
	app.Put("/order/:id", handlers.OrderHandler.UpdateOrder)
	app.Put("/order/status/update", handlers.OrderHandler.UpdateOrderStatus)

	//OrderLine
	app.Post("/orderLine", orderLineHandler.CreateOrderLine)
	app.Post("/orderLines", orderLineHandler.CreateOrderLines)
	app.Get("/orderLines/:id", orderLineHandler.GetOrderLineByID)
	app.Get("/orders/:orderID/orderLines", orderLineHandler.GetOrderLinesByOrderID)
	app.Get("/orderLines/:orderID/:productID", orderLineHandler.GetOrderLineByOrderIDAndProductID)
	app.Get("orderLines", orderLineHandler.GetAllOrderLines)
	app.Put("/orderLines/:id", orderLineHandler.UpdateOrderLine)
	app.Delete("/orderLines/:id", orderLineHandler.DeleteOrderLine)

	//Supplier
	app.Post("/suppliers", supplierHandler.CreateSupplier)
	app.Put("/suppliers/:id", supplierHandler.UpdateSupplier)
	app.Get("/suppliers/:id", supplierHandler.GetSupplierByID)
	app.Get("/suppliers", supplierHandler.GetAllSuppliers)

	//SupplierOrderList
	app.Post("/supplierOrderLists", supplierOrderListHandler.CreateSupplierOrderList)
	app.Get("/supplierOrderLists/:id", supplierOrderListHandler.GetSupplierOrderListByID)
	app.Get("/suppliers/:supplierID/supplierOrderLists", supplierOrderListHandler.GetSupplierOrderListsBySupplierID)
	app.Get("supplierOrderLists", supplierOrderListHandler.GetAllSupplierOrderLists)
	app.Put("/supplierOrderLists/:id", supplierOrderListHandler.UpdateSupplierOrderList)

	

}
