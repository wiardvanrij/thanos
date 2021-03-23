// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

package azure

import (
	"fmt"
	"testing"
)

type test struct {
	name    string
	config  []byte
	wantErr bool
}

var validConfig = []byte(`storage_account: "myStorageAccount"
storage_account_key: "abc123"
container: "MyContainer"
endpoint: "blob.core.windows.net" 
ready_retry_config:
  "max_retry_requests": 100
pipeline_retry_config:
  "try_timeout": 0
  `)

func TestConfig_validate(t *testing.T) {

	conf, err := parseConfig(validConfig)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(conf)
}
