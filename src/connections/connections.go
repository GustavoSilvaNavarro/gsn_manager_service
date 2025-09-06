package connections

import (
	"github.com/gsn_manager_service/src/adapters"
	"github.com/gsn_manager_service/src/adapters/db"
	"github.com/gsn_manager_service/src/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Connections struct {
	Db *mongo.Client
}

func StartConnections() (*Connections, error) {
	client, err := adapters.ConnectToMongoDb(config.Cfg.MONGO_URI)
	if err != nil {
		return nil, err
	}

	return &Connections{
		Db: client,
	}, nil
}

func CreateAllFactories(client *mongo.Client) {
	db.NewTaskRepository(client, config.Cfg.DB_NAME, config.Cfg.TASK_COLLECTION_NAME)
}
