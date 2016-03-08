// +build !exclude_graphdriver_aufs,linux

package register

import (
	// register the aufsext graphdriver
	_ "github.com/docker/docker/daemon/graphdriver/aufsext"
)
