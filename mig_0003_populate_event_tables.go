// Copyright 2020, 2021, 2022 Red Hat, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"database/sql"

	mig "github.com/RedHatInsights/insights-operator-utils/migrations"
	types "github.com/RedHatInsights/insights-results-types"
)

var mig0003PopulateEventTables = mig.Migration{
	StepUp: func(tx *sql.Tx, _ types.DBDriver) error {
		_, err := tx.Exec(`
			INSERT INTO event_targets (id, name, metainfo) 
				VALUES (1, 'notifications backend', 'the target of the report is the ccx notification service back end'),
				       (2, 'service log', 'the target of the report is the ServiceLog')

		`)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`
			UPDATE reported
			   SET event_type_id = 1
			 WHERE event_type_id IS NULL
		`)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`
			 ALTER TABLE reported
			ALTER COLUMN event_type_id SET NOT NULL
		`)
		return err
	},
	StepDown: func(tx *sql.Tx, _ types.DBDriver) error {
		_, err := tx.Exec(`
			 ALTER TABLE reported 
			ALTER COLUMN event_type_id DROP NOT NULL;
		`)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`
			UPDATE reported
			   SET event_type_id = NULL
		`)
		if err != nil {
			return err
		}
		_, err = tx.Exec(`DELETE FROM event_targets`)
		return err
	},
}
