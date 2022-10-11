/*
SPDX-License-Identifier: Apache-2.0
Copyright (C) 2022 Intel Corporation
Copyright (c) 2022 Dell Inc, or its subsidiaries.
Copyright (C) 2022 Red Hat.
*/

package secureAgent

import (
	"errors"
	"log"
	"os"
)

const (
	DHCLIENT_LEASE_FILE = "/var/lib/dhclient/dhclient.leases"
	SZTP_REDIRECT_URL   = "sztp-redirect-urls"
)

func (a *Agent) RunCommandDaemon() error {
	return a.runDaemon()
}

func (a *Agent) runDaemon() error {
	log.Println("[INFO] Get the Bootstrap URL from DHCP client")
	var line string
	if _, err := os.Stat(DHCLIENT_LEASE_FILE); err == nil {
		for {
			line = linesInFileContains(DHCLIENT_LEASE_FILE, SZTP_REDIRECT_URL)
			if line != "" {
				break
			}
		}
		a.BootstrapURL = extractURLfromLine(line, `(?m)[^"]*`)
		log.Println(a)
	} else {
		log.Printf(" File " + DHCLIENT_LEASE_FILE + " does not exist\n")
		return errors.New(" File " + DHCLIENT_LEASE_FILE + " does not exist\n")
	}
	log.Println("[INFO] Bootstrap URL retrieved successfully")
	return nil
}
