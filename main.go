package main

import "ecommerce/cmd"

func main() {

	cmd.Serve()

	// jwt, err := util.CreateJwt("my-secret", util.Payload{
	// 	Sub:         45,
	// 	FirstName:   "Safayet",
	// 	LastName:    "Araf",
	// 	Email:       "araf@gmail.com",
	// 	IsShowOwner: false,
	// })

	// if err != nil {
	// 	fmt.Println("Error creating JWT:", err)
	// } else {
	// 	fmt.Println("JWT:", jwt)
	// }

}
