package init

import (
	"context"
	"log"
	"os"
	"strings"

	"com.fernando/internal/minio"
	mongoClient "com.fernando/internal/mongo"
	"com.fernando/internal/mongo/config"

	//pgClient "com.fernando/internal/postgresql"
	authMongo "com.fernando/pkg/app/auth/repository/mongo"
	authTx "com.fernando/pkg/app/auth/transaction/mongo"
	"com.fernando/pkg/app/ecomm/address/repository/mongo"
	categoryRepository "com.fernando/pkg/app/ecomm/category/repository/mongo"
	productConfigRepository "com.fernando/pkg/app/ecomm/product-config/repository/mongo"
	productItemRepository "com.fernando/pkg/app/ecomm/product-item/repository/mongo"
	productRepository "com.fernando/pkg/app/ecomm/products/repository/mongo"
	productService "com.fernando/pkg/app/ecomm/products/service"
	productTransaction "com.fernando/pkg/app/ecomm/products/transaction/mongo"
	storeRepository "com.fernando/pkg/app/ecomm/stores/repository/mongo"
	storeService "com.fernando/pkg/app/ecomm/stores/service"
	storeTransaction "com.fernando/pkg/app/ecomm/stores/transaction/mongo"
	varOptionRepository "com.fernando/pkg/app/ecomm/variation-option/repository/mongo"
	variationRepository "com.fernando/pkg/app/ecomm/variation/repository/mongo"
	onboardingRepository "com.fernando/pkg/app/onboarding/repository/mongo"
	onboardingService "com.fernando/pkg/app/onboarding/service"
	permissionRepository "com.fernando/pkg/app/permission/repository/mongo"
	"com.fernando/pkg/file-storage/minio/adapter"
	fileStorageService "com.fernando/pkg/file-storage/service"

	//initPG "com.fernando/pkg/services/app/init/repository/postgresql"
	initSrv "com.fernando/pkg/app/init/service"
	//rolePG "com.fernando/pkg/services/app/role/repository/postgresql"
	otpMongo "com.fernando/pkg/app/otp-codes/repository/mongo"
	roleMongo "com.fernando/pkg/app/role/repository/mongo"
	userMongo "com.fernando/pkg/app/user/repository/mongo"
)

type InitConf struct {
}

func NewInitConfig() *InitConf {
	return &InitConf{}
}

// ! Recuerda validar si el admin ya existe
// * no olvidarse cambiar de mongo a postgres y viseversa
// cualquier repo o service lo declaro dentro de RUN() otros desde el server
// debería volver a crear el cliente auth u otro  o mejor usar el ya creado en el server?
// declaro aqui sololo de mi repository y se necesita otro llamar desde el main
// ! tener un estandar para lo nombre de los roles datos q inser tar todos en mayúscula u otro
// ! vaidar que si ya se han creado no vover a crear
// ! analizar si Run() deve estar envuelto en un if env = "dev"
// ? que métodos del coneccion con mongo o postgresql son mas para desarrollo
// y cuales se harán manualmente
func (init *InitConf) Run() {

	// esto es propio de init
	//initRepo := postgresql.NewInitRepository(init.Conn)

	// se puede solo psar roles y en el service usar split o dejarlo asi para agregarlo a user o role service
	// * Envs
	mongoDBname := os.Getenv("MONGO_DB_NAME")
	if mongoDBname == "" {
		log.Fatal("environment variabe MONGO_DB_NAME is not set")
	}

	roles := os.Getenv("ROLES")
	roleNames := strings.Split(roles, ",")

	adminUser := os.Getenv("ADMIN_USER")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	adminRoles := []string{"ADMIN"}

	// * Database connections
	mongoClient := mongoClient.New()
	minioClient := minio.NewClient()
	mongoDB := mongoClient.DB.Database(mongoDBname)
	//pgConn := pgClient.New()

	// * Repos
	//initPGRepo := initPG.NewRepository(pgConn.DB)
	//rolePGRepo := rolePG.NewRepository(pgConn.DB)

	// ! eliminar user collection de authMongo
	authColl := mongoDB.Collection(config.CollAuth)
	userColl := mongoDB.Collection(config.CollUsers)
	twoFaSmsColl := mongoDB.Collection(config.CollTwoFaSms)
	userRoleColl := mongoDB.Collection(config.CollUserRole)
	roleColl := mongoDB.Collection(config.CollRoles)
	rolePermissionColl := mongoDB.Collection(config.CollRolePermission)
	permissionColl := mongoDB.Collection(config.CollPermissions)
	optColl := mongoDB.Collection(config.CollOtp)
	categoryColl := mongoDB.Collection(config.CollCategories)
	variationColl := mongoDB.Collection(config.CollVariations)
	varOptionColl := mongoDB.Collection(config.CollVarOptions)
	productColl := mongoDB.Collection(config.CollProducts) // revisar product collections debe ir en el indiado
	productItemColl := mongoDB.Collection(config.CollProductItems)
	productConfigColl := mongoDB.Collection(config.CollProductConfig)
	onboardingColl := mongoDB.Collection(config.CollOnboarding)
	storeColl := mongoDB.Collection(config.CollStores)
	addressColl :=  mongoDB.Collection(config.CollAddress)

	authMongoRepo := authMongo.NewRepository(mongoClient.DB, authColl)
	userMongoRepo := userMongo.NewRepository(mongoClient.DB, userColl, twoFaSmsColl, userRoleColl)
	roleMongoRepo := roleMongo.NewRepository(mongoClient.DB, roleColl, userRoleColl)
	permissionMongoRepo := permissionRepository.NewRepository(mongoClient.DB, permissionColl, rolePermissionColl)
	otpMongoRepo := otpMongo.NewRepository(mongoClient.DB, optColl)
	categoryRepo := categoryRepository.NewCategoryRepo(mongoClient.DB, categoryColl)
	onboardingRepo := onboardingRepository.NewRepository(mongoClient.DB, onboardingColl)

	productRepo := productRepository.NewRepository(mongoClient.DB, productColl)
	productItemRepo := productItemRepository.NewRepository(mongoClient.DB, productItemColl)
	productConfigRepo := productConfigRepository.NewRepository(mongoClient.DB, productConfigColl)
	variationRepo := variationRepository.NewRepository(mongoClient.DB, variationColl)
	varOptionRepo := varOptionRepository.NewRepository(mongoClient.DB, varOptionColl)
	storeRepo := storeRepository.NewRepository(mongoClient.DB, storeColl)
	addressRepo := mongo.NewAddressRepo(mongoClient.DB,addressColl)

	// * Transaction layer authMongoRepo
	authTx := authTx.NewTransaction(mongoClient.DB, authMongoRepo, userMongoRepo, roleMongoRepo, otpMongoRepo)
	productTx := productTransaction.NewTransaction(mongoClient.DB, productRepo, productItemRepo, productConfigRepo)
	storeTx := storeTransaction.NewStoreTx(mongoClient.DB,storeRepo,addressRepo)
	
	// * Adapter
	minioAdp := adapter.NewAdapter(minioClient.Client)

	// * Services
	fileStorageSrv := fileStorageService.NewFileStgSrv(minioAdp)
	productSrv := productService.NewService(productRepo, fileStorageSrv, productTx)
	onboardingSrv := onboardingService.NewService(onboardingRepo, fileStorageSrv)
	storeSrv := storeService.NewService(storeRepo,storeTx)

	initSrv := initSrv.NewService(authMongoRepo, userMongoRepo, roleMongoRepo, permissionMongoRepo, categoryRepo, variationRepo, varOptionRepo, productRepo, onboardingRepo, onboardingSrv, productSrv, storeSrv, authTx)

	err := initSrv.CreateRoles(context.TODO(), roleNames)
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateAdmin(context.Background(), adminUser, adminPassword, adminRoles)
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateCategories()
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateVariations()
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateVariationOptions()
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateProducts()
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreatePermissions()
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateOnboarding()
	if err != nil {
		log.Fatal(err)
	}

	err = initSrv.CreateStores()
	if err != nil {
		log.Fatal(err)
	}

}
