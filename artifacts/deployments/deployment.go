package deployments

import (
	"github.com/kun-lun/kunlun/artifacts"
	yaml "gopkg.in/yaml.v2"
)

type Deployment struct {
	HostGroupName string
	Vars          yaml.MapSlice
	Roles         []apis.Role
}
