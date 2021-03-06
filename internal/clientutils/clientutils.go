package clientutils

import (
	"errors"
	"preq/internal/pkg/bitbucket"
	"preq/internal/pkg/client"
	"preq/internal/pkg/github"
)

type ClientFactory struct{}

func (cf ClientFactory) DefaultClient(provider client.RepositoryProvider) (client.Client, error) {
	switch provider {
	case client.RepositoryProviderEnum.BITBUCKET:
		return bitbucket.DefaultClient()
	case client.RepositoryProviderEnum.GITHUB:
		return github.DefaultClient()
	}

	return nil, errors.New("unknown provider")
}
