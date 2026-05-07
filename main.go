package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/utilimarc/terraform-provider-claude/internal/provider"
)

var version = "dev"

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/utilimarc/claude",
	}

	err := providerserver.Serve(context.Background(), provider.New(version), opts)
	if err != nil {
		log.Fatal(err)
	}
}
