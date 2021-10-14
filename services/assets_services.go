// handles the entire  business logic of the application
package services

import (
	"github.com/nytro04/asset_management/domain/assets"
	"github.com/nytro04/asset_management/utils/errors"
)

func CreateAsset(asset assets.Asset) (*assets.Asset, *errors.RestErr)  {

	return &asset, nil
}