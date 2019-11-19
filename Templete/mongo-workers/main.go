package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	initCmd = []string{
		"version: '3'\n",
		"\n",
		"services:\n",
		"  mongo:\n",
		"    container_name: mongo-test\n",
		"    image: mongo:4.2.0-bionic\n",
		"    environment:\n",
		"      MONGO_INITDB_ROOT_USERNAME: root\n",
		"      MONGO_INITDB_ROOT_PASSWORD: 123456\n",
		"    volumes:\n",
		"      - ./single:/data/db\n",
		"    ports:\n",
		"      - '127.0.0.1:27047:27017'\n",
		"    restart: always\n",
	}
)

func createFile() {
	err := os.Mkdir("./mongo_env", 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Succeed to create mongo_env")

	file, err := os.Create("./mongo_env/docker-compose.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("Succeed to create ./mongo_env/docker-compose.yml")
}

func initCompose(cmds []string) {
	file, err := os.OpenFile("./mongo_env/docker-compose.yml", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < len(cmds); i++ {
		content := []byte(cmds[i])
		_, err := file.Write(content)

		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Write file successful")
}

func startCompose() {
	_, err := exec.Command("/bin/bash", "-c", "cd ./mongo_env; docker-compose up -d").Output()
	if err != nil {
		fmt.Printf("Fail to execute command, error: %s\n", err)
		return
	}

	fmt.Println("Succeed to execute command, mongo:4.2.0-bionic, user: root, password: 123456")
}

func main() {
	createFile()
	initCompose(initCmd)
	startCompose()
}
