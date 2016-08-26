package openstack

import (
	"github.com/emc-advanced-dev/pkg/errors"
	"github.com/emc-advanced-dev/unik/pkg/types"
)

func (p *OpenstackProvider) ListInstances() ([]*types.Instance, error) {
	return nil, errors.New("not yet supportded for openstack", nil)
}
