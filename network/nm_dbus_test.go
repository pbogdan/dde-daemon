package network

// import "testing"
// import "dlib/dbus"
// import nm "dbus/com/deepin/daemon/network"

// func init() {
// 	manager = NewManager()
// 	dbus.InstallOnSession(manager)
// 	manager.initManager()
// }

// // TODO test case need update
// func TestInvalid(t *testing.T) {
// 	// manager.ActiveAccessPoint("/", "/")

// 	manager.DisconnectDevice("/")

// 	// manager.ActiveWiredDevice("/")

// 	manager.GetAccessPoints("/")

// 	// manager.GetActiveConnection("/")

// 	// manager.GetConnectionByAccessPoint("/")

// 	manager.FeedSecret("xxoo", "sd", "ss")
// 	manager.GetDBusInfo()

// 	// dumy := make(map[string]map[string]string)
// 	// manager.UpdateConnection(dumy)
// }

// func TestDBusFailed(t *testing.T) {
// 	if _, err := nm.NewNetworkManager("com.deepin.daemon.Network", "/"); err == nil {
// 		t.Fail()
// 	}

// 	m, err := nm.NewNetworkManager("com.deepin.daemon.Network", "/com/deepin/daemon/Network")
// 	if err != nil {
// 		t.Fatal("NewNetworkManager1")
// 	}
// 	if _, err := nm.NewNetworkManager("com.deepin.daemon.Network", "/com/deepin/daemon/Networkxx"); err == nil {
// 		t.Fatal("NewNetworkManager2")
// 	}
// 	// if err = m.ActiveAccessPoint("/", "/"); err == nil {
// 	// 	t.Fatal("ActiveAccessPoint")
// 	// }

// 	if err = m.DisconnectDevice("/"); err == nil {
// 		t.Fatal("DisconnectDevice")
// 	}
// 	// if err = m.ActiveWiredDevice("/"); err == nil {
// 	// 	t.Fatal("ActiveWiredDevice")
// 	// }
// 	if _, err = m.GetAccessPoints("/"); err == nil {
// 		t.Fatal("GetAccessPoints")
// 	}
// 	// if _, err = m.GetActiveConnection("/"); err == nil {
// 	// 	t.Fatal("GetActiveConnection")
// 	// }
// 	// if _, err = m.GetConnectionByAccessPoint("/"); err == nil {
// 	// 	t.Fatal("GetConnectionByAccessPoint")
// 	// }
// 	//if err = m.FeedSecret("xxoo", "s", "sd"); err == nil {
// 	////This wouldn't failed
// 	//}
// }

// func TestDBusSuccess(t *testing.T) {
// 	m, err := nm.NewNetworkManager("com.deepin.daemon.Network", "/com/deepin/daemon/Network")
// 	if err != nil {
// 		t.Fatal("Create NetworkManager failed")
// 	}
// 	for _, d := range m.WirelessDevices.Get() {
// 		path := d[0].(dbus.ObjectPath)
// 		// aps, err := m.GetAccessPoints(path)
// 		_, err := m.GetAccessPoints(path)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		// for _, ap := range aps {
// 		// logger.Debug(ap)
// 		// }
// 	}
// }