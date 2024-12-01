// wire.go
//go:build wireinject
// +build wireinject

package dependency_injector

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
	"koriebruh/management/cnf"
	"koriebruh/management/controller"
	"koriebruh/management/repository"
	"koriebruh/management/service"
)

func ProvideDB() *gorm.DB {
	return cnf.InitDB()
}

func ProvideValidator() *validator.Validate {
	return validator.New()
}

var AuthSet = wire.NewSet(
	repository.NewAuthRepository,
	service.NewAuthService,
	controller.NewAuthController,
	wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)),
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
	ProvideDB,
	ProvideValidator,
)

var CategorySet = wire.NewSet(
	repository.NewCategoryRepository,
	service.NewCategoryService,
	controller.NewCategoryController,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
	ProvideDB,
	ProvideValidator,
)

var ItemSet = wire.NewSet(
	repository.NewItemRepository,
	service.NewItemService,
	controller.NewItemController,
	wire.Bind(new(repository.ItemRepository), new(*repository.ItemRepositoryImpl)),
	wire.Bind(new(service.ItemService), new(*service.ItemServiceImpl)),
	wire.Bind(new(controller.ItemController), new(*controller.ItemControllerImpl)),
	ProvideDB,
	ProvideValidator,
)

var SupplierSet = wire.NewSet(
	repository.NewSupplierRepository,
	service.NewSupplierService,
	controller.NewSupplierController,
	wire.Bind(new(repository.SupplierRepository), new(*repository.SupplierRepositoryImpl)),
	wire.Bind(new(service.SupplierService), new(*service.SupplierServiceImpl)),
	wire.Bind(new(controller.SupplierController), new(*controller.SupplierControllerImpl)),
	ProvideDB,
	ProvideValidator,
)

func InitializeAuth() (controller.AuthController, error) {
	wire.Build(AuthSet)
	return nil, nil
}

func InitializeCategory() (controller.CategoryController, error) {
	wire.Build(CategorySet)
	return nil, nil
}

func InitializeItem() (controller.ItemController, error) {
	wire.Build(ItemSet)
	return nil, nil
}

func InitializeSupplier() (controller.SupplierController, error) {
	wire.Build(SupplierSet)
	return nil, nil
}
