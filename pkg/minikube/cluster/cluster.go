/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"flag"
	"fmt"

	"github.com/docker/machine/libmachine"
	"github.com/docker/machine/libmachine/ssh"
	"github.com/pkg/errors"

	"k8s.io/minikube/pkg/minikube/bootstrapper"
	"k8s.io/minikube/pkg/minikube/bootstrapper/kubeadm"
	"k8s.io/minikube/pkg/minikube/exit"
)

// This init function is used to set the logtostderr variable to false so that INFO level log info does not clutter the CLI
// INFO lvl logging is displayed due to the kubernetes api calling flag.Set("logtostderr", "true") in its init()
// see: https://github.com/kubernetes/kubernetes/blob/master/pkg/kubectl/util/logs/logs.go#L32-L34
func init() {
	if err := flag.Set("logtostderr", "false"); err != nil {
		exit.WithError("unable to set logtostderr", err)
	}

	// Setting the default client to native gives much better performance.
	ssh.SetDefaultClient(ssh.Native)
}

// Bootstrapper returns a new bootstrapper for the cluster
func Bootstrapper(api libmachine.API, bootstrapperName string) (bootstrapper.Bootstrapper, error) {
	var b bootstrapper.Bootstrapper
	var err error
	switch bootstrapperName {
	case bootstrapper.Kubeadm:
		b, err = kubeadm.NewBootstrapper(api)
		if err != nil {
			return nil, errors.Wrap(err, "getting a new kubeadm bootstrapper")
		}
	default:
		return nil, fmt.Errorf("unknown bootstrapper: %s", bootstrapperName)
	}
	return b, nil
}
