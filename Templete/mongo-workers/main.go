package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	// Name -
	Name = os.Getenv("NAME")
	// Password -
	Password = os.Getenv("PASSWORD")
	// Ports -
	Ports = os.Getenv("PORTS")
)

func main() {
	fmt.Println(Name, Password, Ports)
	message := []byte(generateDockerCompose(Name, Password, Ports))
	err := ioutil.WriteFile("docker-compose.yaml", message, 0644)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("/bin/bash", "-c", `echo "#! /bin/bash
docker-compose up -d " >docker.sh`)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "echo", err.Error())
		return
	}

	cmd = exec.Command("/bin/bash", "-c", "chmod 777 docker.sh")
	output, err = cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", "chmod", err.Error())
		return
	}

	command := `./docker.sh .`
	cmd = exec.Command("/bin/bash", "-c", command)
	output, err = cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func generateDockerCompose(name, password, ports string) string {

	data := fmt.Sprintf(
		`version: "3"

services:
  mongo:
    container_name: %s
    image: mongo:4.2.0-bionic
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: %s
    volumes:
      - ./single:/data/db
    ports:
      - "%s"
    restart: always`, name, password, ports)

	return data
}
