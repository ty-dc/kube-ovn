package util

const (
	DefaultHostVhostuserBaseDir = "/run/openvswitch/vhost_sockets"

	ChassisLoc = "/etc/openvswitch/system-id.conf"

	VfioSysDir = "/sys/bus/pci/drivers/vfio-pci"
	NetSysDir  = "/sys/class/net"

	HtbQos   = "linux-htb"
	NetemQos = "linux-netem"

	KoDir  = "/tmp/"
	KoENV  = "MODULES"
	RpmENV = "RPMS"

	DefaultCniConfigFile = "/kube-ovn/01-kube-ovn.conflist"
	DefaultCniConfigDir  = "/etc/cni/net.d"
)
