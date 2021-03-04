package main

import "fmt"
import "os/exec"
import "log"

func main() {
        cmd := exec.Command("sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	err = cmd.Wait()
	fmt.Println("Hi there!")
}
