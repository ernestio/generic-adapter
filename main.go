/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"log"
	"os"
	"runtime"

	ecc "github.com/ernestio/ernest-config-client"
	"github.com/nats-io/nats"
	o "github.com/r3labs/otomo"
)

var nc *nats.Conn
var natsErr error

func main() {
	// TODO : Types probably need to be get from the config with getConnectorTypes
	//types := []string{"elb", "ebs", "s3", "rds", "route53", "elasticsearch"}
	components := []string{"vpc", "elb"}
	types := []string{"aws-fake", "vcloud-fake", "aws", "vcloud", "fake"}
	nc = ecc.NewConfig(os.Getenv("NATS_URI")).Nats()

	c := o.Config{
		Client:     nc,
		ValidTypes: types,
	}

	for _, v := range components {
		log.Println("Setting up " + v)
		o.StandardSubscription(&c, v+".create", "_type")
		o.StandardSubscription(&c, v+".delete", "_type")
	}

	runtime.Goexit()
}
