// +build !exclude_graphdriver_aufsext,linux

package daemon

import (
	_ "github.com/docker/docker/daemon/graphdriver/aufsext"
)
