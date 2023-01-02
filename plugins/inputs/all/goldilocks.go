//go:build !custom || inputs || inputs.goldilocks

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/goldilocks" // register plugin
