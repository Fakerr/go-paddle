// The simple command demonstrates a simple functionality which
// list all users subscribed to any of your subscription plans
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Fakerr/go-paddle/paddle"
)

func main() {

	var vendorID = os.Getenv("VENDOR_ID")
	var vendorAuthCode = os.Getenv("VENDOR_AUTH_CODE")
	var planID = os.Getenv("PLAN_ID")

	client := paddle.NewClient(vendorID, vendorAuthCode, nil)

	options := &paddle.UsersOptions{
		PlanID:      planID,
		ListOptions: paddle.ListOptions{ResultsPerPage: 10},
	}

	users, _, err := client.Users.List(context.Background(), options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for i, user := range users {
		fmt.Printf("%v. %v\n", i+1, *user.UserID)
	}
}
