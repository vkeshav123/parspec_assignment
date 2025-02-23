package providers

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

type YMLConfig struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"db"`
	Order    Order    `yaml:"order"`
}

type Order struct {
	LoadCreationBatchSize    int `yaml:"load_creation_batch_size"`
	OrderProcessingBatchSize int `yaml:"processing_batch_size"`
	OrderProcessingInterval  int `yaml:"processing_interval"`
	OrderCompletionBatchSize int `yaml:"completion_batch_size"`
	OrderCompletionInterval  int `yaml:"completion_interval"`
}

type App struct {
	Name string `yaml:"name"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func GetConfig() YMLConfig {
	if _, err := os.Stat("oms-service-configuration.yml"); os.IsNotExist(err) {
		panic("Configuration file not found!")
	}
	data, err := os.ReadFile("oms-service-configuration.yml")
	if err != nil {
		panic(err)
	}
	var yamlConfig YMLConfig
	err = yaml.Unmarshal(data, &yamlConfig)
	if err != nil {
		fmt.Println("error in unmarshalling yaml - ", err)
		panic(err)
	}
	return yamlConfig
}

func GetDbConnection() *gorm.DB {
	YMLConfig := GetConfig()
	connectionCommand := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", YMLConfig.Database.Host, YMLConfig.Database.Port, YMLConfig.Database.Name, YMLConfig.Database.User, YMLConfig.Database.Password)
	fmt.Print("Connection Command")
	fmt.Print(connectionCommand)
	dbInstance, err := gorm.Open("postgres", connectionCommand)
	if err != nil {
		fmt.Print("error in getting db connection - ", err)
		panic(err)
	}
	return dbInstance
}

func GetOrderConfig() Order {
	ymlConfig := GetConfig()
	return ymlConfig.Order
}
