// package main

// // https://cli.urfave.org/v1/examples/greet/
// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/urfave/cli"
// )

// func main() {
// 	app := cli.NewApp()
// 	app.Name = "greet"
// 	app.Usage = "fight the loneliness!"
// 	app.Action = func(c *cli.Context) error {
// 		resp, err := http.Get(c.Args().Get(0))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		body, err := io.ReadAll(resp.Body)
// 		if resp.StatusCode > 299 {
// 			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
// 		}
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer resp.Body.Close()
// 		fmt.Printf("%s", body)
// 		return nil
// 	}

// 	err := app.Run(os.Args)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "env",
			Value: "dev",
			Usage: "language for the greeting",
		},
	}
	app.Action = func(c *cli.Context) error {
		host := ""
		if c.String("env") == "prod" {
			host = "http://brainiakv2-api-prod.http-apps.tsuru.gcp.i.globo/healthcheck"
		} else {
			host = "http://brainiakv2-api-qa.http-apps.tsuru.dev.gcp.i.globo/healthcheck"
		}
		printone(host)
		printtwo(host)
		printthree()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func printone(h string) {
	res, err := http.Get(h)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", body)
	// os.WriteFile("test.txt", body, 0777)
	data := []byte(body)
	fmt.Printf("%x", md5.Sum(data))
}

func printtwo(n1 string) {
	fmt.Println("Hello friend2!", n1)
}

func printthree() {
	fmt.Println("Hello friend3!")
}
