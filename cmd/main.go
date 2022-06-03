package main

import (
	"fmt"
	"github.com/aalkan/plesk-go/pkg/client"
	"github.com/aalkan/plesk-go/pkg/models"
)

func main() {

	client := client.NewClient(&models.Config{Username: "test", Password: "test"})
	res, err := client.GetLicense(&models.LicenseGet{})
	fmt.Println(res, err)
}
