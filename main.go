/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"time"

	ecc "github.com/ernestio/ernest-config-client"
	"github.com/nats-io/nats"
	o "github.com/r3labs/otomo"
)

var nc *nats.Conn
var natsErr error

func getConnectorTypes(ctype string) []string {
	var connectors map[string][]string

	resp, err := nc.Request("config.get.connectors", nil, time.Second)
	if err != nil {
		log.Println("could not get config for connectors")
		log.Fatal(err)
	}

	err = json.Unmarshal(resp.Data, &connectors)
	if err != nil {
		log.Println("could not read config response")
		log.Fatal(err)
	}

	if connectors[ctype] == nil {
		log.Fatal("connector type not found")
	}

	return connectors[ctype]
}

func main() {
	// TODO : Types probably need to be get from the config with getConnectorTypes
	//types := []string{"elb", "ebs", "s3", "rds", "route53", "elasticsearch"}
	types := []string{"vpc"}
	nc = ecc.NewConfig(os.Getenv("NATS_URI")).Nats()

	c := o.Config{
		Client:     nc,
		ValidTypes: types,
	}

	for _, v := range types {
		log.Println("Setting up " + v)
		o.StandardSubscription(&c, v+".create", "_type")
		o.StandardSubscription(&c, v+".delete", "_type")
	}

	runtime.Goexit()
}
