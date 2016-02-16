// package deptesthuge pulls lots of unused dependencies
// for testing dependency resolution.

package deptesthuge

import (
	_ "github.com/coreos/coreos-cloudinit/datasource"
	_ "github.com/coreos/coreos-cloudinit/datasource/url"
	_ "github.com/docker/docker/api"
	_ "github.com/docker/docker/pkg/ioutils"
	_ "github.com/docker/docker/pkg/system"
)

type Unused struct{}
