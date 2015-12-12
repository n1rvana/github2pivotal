package main

import (
	"flag"
	"github.com/google/go-github/github"
	"log"
	"oauth2"
)

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Printf("USAGE: g2p <Github Personal API Token> ")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: args[0]},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List("", nil)
	log.Printf("Authorized users repos: %+v\n\n", repos)
	orgs, _, err := client.Organizations.List("", nil)
	log.Printf("Authorized users orgs: %+v\n\n", orgs)

}

//  fmt.Printf("USAGE: %+v\n\n", se)
