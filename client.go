package goclient

import (
	"github.com/flowswiss/goclient/v2/common"
	"github.com/flowswiss/goclient/v2/compute"
	"github.com/flowswiss/goclient/v2/core"
	"github.com/flowswiss/goclient/v2/kubernetes"
	"github.com/flowswiss/goclient/v2/macbaremetal"
	"github.com/flowswiss/goclient/v2/objectstorage"
)

type Client struct {
	Common        *common.Service
	Compute       *compute.Service
	Kubernetes    *kubernetes.Service
	MacBareMetal  *macbaremetal.Service
	ObjectStorage *objectstorage.Service
}

func newClient(client *core.Client) *Client {
	return &Client{
		Common:        common.NewService(client),
		Compute:       compute.New(client),
		Kubernetes:    kubernetes.NewService(client),
		MacBareMetal:  macbaremetal.NewService(client),
		ObjectStorage: objectstorage.NewService(client),
	}
}

func WithToken(token string) *Client {
	return newClient(core.NewClient(core.ClientOpts{
		Token: token,
	}))
}

func WithClient(client *core.Client) *Client {
	return newClient(client)
}
