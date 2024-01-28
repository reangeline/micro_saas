//go:build wireinject
// +build wireinject

package dependency_injection

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/wire"
	repo_interface "github.com/reangeline/micro_saas/internal/domain/contract/repository"
	uc_interface "github.com/reangeline/micro_saas/internal/domain/contract/usecasse"

	"github.com/reangeline/micro_saas/internal/domain/usecase"
	"github.com/reangeline/micro_saas/internal/infra/database"

	"github.com/reangeline/micro_saas/internal/presentation/controller"
)

var setUserUseCaseDependency = wire.NewSet(
	usecase.NewUserUseCase,
	wire.Bind(new(uc_interface.UserUseCaseInterface), new(*usecase.UserUseCase)),
)

var setUserRepositoryDependency = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(repo_interface.UserRepositoryInterface), new(*database.UserRepository)),
)

func InitializeUser(vc *dynamodb.DynamoDB) (*controller.UserController, error) {
	wire.Build(
		setUserRepositoryDependency,
		setUserUseCaseDependency,
		controller.NewUserController,
	)
	return &controller.UserController{}, nil
}
