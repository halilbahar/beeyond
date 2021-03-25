package test

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"strings"
	"testing"
	"yaml-validation/conf"
	"yaml-validation/routers"
	"yaml-validation/services"
)

var Router *gin.Engine
var mongoDbContainer testcontainers.Container

func TestMain(m *testing.M) {
	conf.Init()

	setupMongoDbContainer()

	services.Init()

	Router = routers.GetRouter()

	code := m.Run()
	mongoDbContainer.Terminate(context.Background())
	os.Exit(code)
}

func setupMongoDbContainer() {
	//credential := options.Credential{
	//	Username: conf.Configuration.Database.User,
	//	Password: conf.Configuration.Database.Password,
	//}
	//
	//clientOpts := options.Client().
	//	ApplyURI(conf.Configuration.Database.Type + "://" + conf.Configuration.Database.Host + ":" + conf.Configuration.Database.Port).
	//	SetAuth(credential)
	//
	//client, err := mongo.Connect(context.TODO(), clientOpts)
	//
	//if err == nil{
	//	client.Disconnect(context.TODO())
	//	return
	//}
	// polleichtner, wos// 10 min nu donn kinan ma e doa, jo i check nd warum wia erm jz zuahochn miasn ergibt e kan sinn wei ma e nd ds von gestan wissn, voi oba i was warum di
	// tests nd gaunga san warum? jo wegen dem drüba do, ds woa ds wos er ma augsogt hod wos passt do leicht nd, i was a nd euso do schaut ma afoch noch falls di normal db nd rennt
	// dass er si glei mid port 27017 vabindet

	mongoDbContext := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mongo",
		ExposedPorts: []string{conf.Configuration.Database.Port + "/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections"),
		Env: map[string]string{
			"MONGO_INITDB_DATABASE":      "beeyond_validation_db",
			"MONGO_INITDB_ROOT_USERNAME": "beeyond",
			"MONGO_INITDB_ROOT_PASSWORD": "beeyond"},
	}

	mongoDbContainer, _ = testcontainers.GenericContainer(mongoDbContext, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	_, _ = mongoDbContainer.Host(mongoDbContext)
	port, _ := mongoDbContainer.MappedPort(mongoDbContext, nat.Port(conf.Configuration.Database.Port))

	conf.Configuration.Database.Port = strings.Split(string(port), "/")[0]

	fmt.Println("------------------------------------------test db port: " + conf.Configuration.Database.Port)
}
