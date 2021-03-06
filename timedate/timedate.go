/*
 * Copyright (C) 2013 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package timedate

import (
	"pkg.deepin.io/dde/daemon/loader"
	"pkg.deepin.io/lib/log"
)

var (
	_manager *Manager

	logger = log.NewLogger("daemon/timedate")
)

type Daemon struct {
	*loader.ModuleBase
}

func NewDaemon(logger *log.Logger) *Daemon {
	daemon := new(Daemon)
	daemon.ModuleBase = loader.NewModuleBase("timedate", daemon, logger)
	return daemon
}

func (d *Daemon) GetDependencies() []string {
	return []string{}
}

// Start to run time date manager
func (d *Daemon) Start() error {
	if _manager != nil {
		return nil
	}
	service := loader.GetService()

	var err error
	_manager, err = NewManager(service)
	if err != nil {
		return err
	}

	err = service.Export(dbusPath, _manager)
	if err != nil {
		_manager.destroy()
		_manager = nil
		return err
	}

	err = service.RequestName(dbusServiceName)
	if err != nil {
		_manager.destroy()
		_manager = nil
		return err
	}

	go func() {
		_manager.init()
		_manager.handlePropChanged()
	}()
	return nil
}

// Stop the time date manager
func (d *Daemon) Stop() error {
	if _manager == nil {
		return nil
	}

	_manager.destroy()
	_manager = nil
	return nil
}
