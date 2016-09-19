// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package kmod

import (
	"bytes"

	"github.com/snapcore/snapd/osutil"
)

func writeModulesFile(modules [][]byte, modulesFilePath string) error {
	var buffer bytes.Buffer
	buffer.WriteString("# This file is automatically generated.\n")
	for _, module := range modules {
		buffer.Write(module)
		buffer.WriteByte('\n')
	}

	modulesFile := &osutil.FileState{
		Content: buffer.Bytes(),
		Mode:    0644,
	}

	if err := osutil.EnsureFileState(modulesFilePath, modulesFile); err == osutil.ErrSameState {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
