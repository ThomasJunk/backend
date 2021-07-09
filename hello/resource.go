// Copyright (c) 2021 Thomas Junk
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package hello

import (
	"github.com/ThomasJunk/backend/db"
	"go.uber.org/zap"
)

type HelloWorld struct {
	Logger   *zap.Logger
	Database *db.Database
}
