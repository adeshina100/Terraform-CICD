package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kislerdm/terraform-provider-awsamplify/client"
)

var (
	c      *client.Client
	region string
	appId  string
	branch string
	path   string
	ctx    = context.Background()
	app    = "awsamplifydeployzip"
	logs   = log.New(os.Stdout, "", log.LUTC|log.LstdFlags)
)

func stdInFlag() {
	f := flag.NewFlagSet("deploy", flag.ExitOnError)
	f.StringVar(&region, "region", "", "AWS Region, required")
	f.StringVar(&appId, "app-id", "", "AWS Amplify application ID, required")
	f.StringVar(&branch, "branch", "", "AWS Amplify frontend environment branch, required")
	f.StringVar(&path, "path", "", "Path to zip archive with the app, required")
	f.Usage = func() {
		fmt.Println("Tool to deploy web client to AWS Amplify from zip archive")
		fmt.Printf("\nUsage: %s [Args]\n\nArguments:\n", app)
		f.PrintDefaults()
	}
	f.Parse(os.Args[1:])

	if region == "" || appId == "" || branch == "" || path == "" {
		logs.Printf("[ERROR] Missing required arguments\n\n")
		f.Usage()
		os.Exit(1)
	}
}

func init() {
	stdInFlag()
	c = client.New(region)
}

func main() {
	if err := c.DeployZip(ctx, appId, branch, path); err != nil {
		log.Fatalln(err)
	}
}
