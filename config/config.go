package config

import (
	"context"
	"flag"

	"github.com/sirupsen/logrus"
	"go.deanishe.net/env"
)

var (
	// Erlang application config
	Version     = "0.0.1"
	AppName     = "go-space"
	Description = "SPACE node written in go to handle libp2p functionality."

	log *logrus.Logger

	// Package internal config
	VaultAddr  string = env.Get("GO_SPACE_VAULT_ADDR", "http://localhost:8200")
	VaultToken string = env.Get("GO_SPACE_VAULT_TOKEN", "space")

	// Global config
	LogLevel      string = env.Get("GO_SPACE_LOG_LEVEL", "error")
	SPACENodeName string = env.Get("GO_SPACE_SPACE_NODE_NAME", "space@localhost")
	NodeCookie    string = env.Get("GO_SPACE_NODE_COOKIE", "space")
	NodeName      string = env.Get("GO_SPACE_NODE_NAME", "pubsub@localhost")
	Rendezvous    string = env.Get("GO_SPACE_RENDEZVOUS", "space")
)

func Init(ctx context.Context) {

	// Flags - user configurations
	flag.StringVar(&LogLevel, "loglevel", LogLevel, "Loglevel to use for application")
	flag.StringVar(&SPACENodeName, "space_nodename", SPACENodeName, "Name of the node running the actual SPACE")
	flag.StringVar(&NodeCookie, "nodecookie", NodeCookie, "Secret shared by all erlang nodes in the cluster")
	flag.StringVar(&NodeName, "nodename", NodeName, "Name of the erlang node")
	flag.StringVar(&Rendezvous, "rendezvous", Rendezvous, "Unique string to identify group of nodes. Share this with your friends to let them connect with you")
	flag.StringVar(&VaultAddr, "vaultaddr", VaultAddr, "Address of the vault server")
	flag.StringVar(&VaultToken, "vaulttoken", VaultToken, "Token to use to authenticate with the vault server. This is required.")

	flag.Parse()

	// Init logger
	log = logrus.New()
	level, err := logrus.ParseLevel(LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(level)
	log.Info("Logger initialized")

}
