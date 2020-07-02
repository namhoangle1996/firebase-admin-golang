package main

import (
	"context"
	firebase "firebase.google.com/go"
	//"fmt"
	"google.golang.org/api/option"
)

func main()  {
	type Account struct {
		Name string `json:"name"`
		Balance int64 `json:"balance"`
	}

	config := &firebase.Config{
		DatabaseURL: "https://dicom-c6569.firebaseio.com",
	}

	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		panic(err)
	}

	client, err := app.Database(context.Background())
	if err != nil {
		panic(err)
	}
	_ = client

	// @todo : set
	//acc := Account{
	//	Name:    "Alice",
	//	Balance: 1000,
	//}
	//if err := client.NewRef("accounts/alice").Set(context.Background(), acc); err != nil {
	//	panic(err)
	//}

	// @todo : get
	//var getAcc Account
	//if err := client.NewRef("accounts/alice").Get(context.Background(), &getAcc); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("get account name %s \n",getAcc.Name)
	//fmt.Printf("get account balance %.2f",getAcc.Balance)


	// @todo : push
	//arr1 := 200
	//val, err := client.NewRef("push/arr").Push(context.Background(), arr1)
	//if err != nil {
	//	panic(err)
	//}
 	//_ = val

 	// @todo verfify token email
	//idToken := "AM5PThCtCl0XQHw_oLvS6_7hio0mm7rbS1oba4dhzd5qUwqJccjMURGatozp0ZS-TbW09GJ3sokAs30H3a8jRAaLxAib31v0KSf8mEa6EzMjcKHymyqprg9zXdy4yDUI4J5CXmL3USsgSxdf7-ZTldn4NLTPDByQ_A"
	//getVerify, err := client.VerifyIDToken(context.Background(), idToken)
	//if err != nil {
	//	panic(err)
	//}else {
	//	log.Log(getVerify)
	//}



}

