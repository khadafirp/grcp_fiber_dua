package client

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/khadafirp/grpc_fiber_dua/github.com/khadafirp/grpc_fiber_dua/proto/greeter"
	"google.golang.org/grpc"
)

func StartClient() {

	time.Sleep(1 * time.Second) // tunggu gRPC ready (seharusnya pakai health check)

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()
	client := greeter.NewBarangServiceClient(conn)

	app := fiber.New()

	app.Get("/hello/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		res, err := client.AllBarang(context.Background(), &greeter.BarangRequest{Name: name})
		if err != nil {
			return c.Status(500).SendString("gRPC error: " + err.Error())
		}

		return c.JSON(fiber.Map{
			"message": res.GetMessage(),
		})
	})

	log.Println("Fiber HTTP server running on :3001")
	log.Fatal(app.Listen(":3001"))
}
