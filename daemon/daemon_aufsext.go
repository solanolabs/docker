// +build !exclude_graphdriver_aufsext

package daemon

import (
	_ "github.com/docker/docker/daemon/graphdriver/aufsext"
)
