package hosts

import (
	"github.com/docker/docker/client"
	"github.com/Sirupsen/logrus"
)

type Hosts struct {
	Hosts []Host `yaml:"hosts"`
}

type Host struct {
	IP           string   `yaml:"ip"`
	Role         []string `yaml:"role"`
	Hostname     string   `yaml:"hostname"`
	User         string   `yaml:"user"`
	Sudo         bool     `yaml:"sudo"`
	DockerSocket string   `yaml:"docker_socket"`
	DClient      *client.Client
}

func DivideHosts(hosts []Host) ([]Host, []Host, []Host) {
	etcdHosts := []Host{}
	cpHosts := []Host{}
	workerHosts := []Host{}
	for _, host := range hosts {
		for _, role := range host.Role {
			logrus.Debugf("Host: " + host.Hostname + " has role: " + role)
			if role == "etcd" {
				etcdHosts = append(etcdHosts, host)
			}
			if role == "controlplane" {
				cpHosts = append(cpHosts, host)
			}
			if role == "worker" {
				workerHosts = append(workerHosts, host)
			}
		}
	}
	return etcdHosts, cpHosts, workerHosts
}
