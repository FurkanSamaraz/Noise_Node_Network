package API

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/perlin-network/noise"
)

var username UserName
var nodename NodeName

type UserName struct {
	Names string `json:"name"`
}
type NodeName struct {
	NodeName *noise.Node `json:"nodename"`
	NodeAddr string      `json:"nodeaddress"`
}

func Node(c *fiber.Ctx) error {

	name, err := noise.NewNode()
	name2, err := noise.NewNode()

	check(err)

	defer name.Close()
	defer name2.Close()

	name.Handle(func(ctx noise.HandlerContext) error {
		if !ctx.IsRequest() {
			return nil
		}

		fmt.Printf("Got a message from Alice: '%s'\n", string(ctx.Data()))

		return ctx.Send([]byte("Hi Alice!"))
	})

	check(name.Listen())
	check(name2.Listen())
	res, err := name2.Request(context.TODO(), ":50912", []byte("Hi Bob!"))
	fmt.Println("Address", name.Addr())

	return c.JSON(res)
}
func User(c *fiber.Ctx) error {

	names := c.FormValue("name")
	names = username.Names

	json.MarshalIndent(username, "", "\t")
	fmt.Printf(names)
	return c.JSON(names)
}

func Send(c *fiber.Ctx) error {
	recipient := c.FormValue("address")
	res, err := nodename.NodeName.Request(context.TODO(), recipient, []byte("Hi Bob!"))
	check(err)
	byte, _ := json.MarshalIndent(res, "", "\t")
	return c.JSON(byte)
}
func Api() {
	app := fiber.New()
	app.Post("/user", User)

	app.Get("/node", Node)
	app.Post("/send", Send)

	app.Listen(":8000")
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

/*
bob

	bob, err := noise.NewNode()

	check(err)

	// Gracefully release resources for Alice and Bob at the end of the example.

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

	check(bob.Listen())

	// Have Alice send Bob a request with the message 'Hi Bob!'

	// Print out the response Bob got from Alice.

	fmt.Printf("Got a message from Bob: '%s'\n", string(bob.Addr()))

	return c.JSON(bob.Addr())


,*/

/*

alice
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

	return c.JSON(bob.Addr())*/
