package driver1

import "github.com/mirno/goexplorer/pkg/config"

type Driver1 struct {
    Config config.DriverConfig
}

func (d *Driver1) PerformTask() error {
    // Implementation specific to Driver1
    return nil
}