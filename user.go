/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"database/sql"
)

type User struct {
	ID             string         `db:"id"`
	GithubUsername sql.NullString `db:"github_username"`
	Plan           string         `db:"plan"`
}
