package main

import (
	"context"
	"fmt"
	"main/API"

	"github.com/gofiber/fiber/v2"
	"github.com/perlin-network/noise"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func node(c *fiber.Ctx) error {

	alice, err := noise.NewNode()
	check(err)

	bob, err := noise.NewNode()

	check(err)

	// Gracefully release resources for Alice and Bob at the end of the example.

	defer alice.Close()
	defer bob.Close()

	// When Bob gets a message from Alice, print it out and respond to Alice with 'Hi Alice!'

	bob.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Got a message from Alice: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte("Hi Alice!"))
	})

	// Have Alice and Bob start listening for new peers.

	check(alice.Listen())
	check(bob.Listen())

	// Have Alice send Bob a request with the message 'Hi Bob!'

	res, err := alice.Request(context.TODO(), ":50802", []byte("Hi Bob!"))
	check(err)

	// Print out the response Bob got from Alice.

	fmt.Printf("Got a message from Bob: '%s'\n", string(res))

	return c.JSON(bob.Addr())
}
func main() {
	/*alice, _ := noise.NewNode()

	a := alice.Addr()
	fmt.Printf(a)
	defer alice.Close()
	app := fiber.New()

	app.Get("/", node)

	app.Listen(":8000")*/

	API.Api()
}
