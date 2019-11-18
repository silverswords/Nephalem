package main

import (
	"fmt"
	"os"
	"os/exec"
)

var initCmd []string = []string{
	"version: '2'\n",
	"services:\n",
	"  mysql:\n",
	"    environment:\n",
	"      MYSQL_ROOT_PASSWORD: '123456'\n",
	"      MYSQL_USER: 'root'\n",
	"      MYSQL_PASS: '123456'\n",
	"    image: 'docker.io/mysql:latest'\n",
	"    volumes:\n",
	"      - './db:/var/lib/mysql'\n",
	"      - './conf/my.cnf:/etc/my.cnf'\n",
	"      - './init:/docker-entrypoint-initdb.d/'\n",
	"    ports:\n",
	"      - '3306:33060'",
}

func createFile() {
	err := os.Mkdir("./mysql_space", 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("succeed to create mysql_space")
}

func createCompose() {
	file, err := os.Create("./mysql_space/docker-compose.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("succeed to create ./mysql_space/docker-compose.yml")
}

func initCompose(cmds []string) {
	file, err := os.OpenFile("./mysql_space/docker-compose.yml", os.O_RDWR|os.O_CREATE, 0777)
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

	fmt.Println("write file successful")
}

func startCompose() {
	_, err := exec.Command("bash", "-c", "cd ./mysql_space; docker-compose up -d").Output()
	if err != nil {
		fmt.Printf("fail to execute command, error: %s\n", err)
		return
	}

	fmt.Printf("succeed to execute command, mysql:latest user: root, password: 123456")
}

func main() {
	createFile()
	createCompose()
	initCompose(initCmd)
	startCompose()
}
