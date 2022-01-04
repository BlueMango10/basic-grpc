package main

import (
	"context"
	"fmt"

	"example.com/grpcdemo/shared"
	"google.golang.org/grpc"
)

func main() {
	// Get servers address
	fmt.Print("Server address: ")
	var address string
	fmt.Scanln(&address)

	startClient(address)
}

func startClient(address string) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close()

	client := shared.NewDemoServiceClient(conn)

	mainloop(client)
}

func mainloop(client shared.DemoServiceClient) {
	for {
		// Ask for input
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		// Make request
		response, err := client.Poke(context.Background(), &shared.Finger{
			Description: input,
		})
		if err != nil { // Check for errors
			fmt.Println(err.Error())
		} else { // If no errors, we can use the request
			for _, insult := range response.Insults {
				fmt.Println(insult)
			}
		}
	}
}
