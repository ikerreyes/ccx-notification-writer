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

// Generated documentation is available at:
// https://pkg.go.dev/github.com/RedHatInsights/ccx-notification-writer/
//
// Documentation in literate-programming-style is available at:
// https://redhatinsights.github.io/ccx-notification-writer/packages/mig_0004_update_reported_table_with_eventtype.html

import (
	"database/sql"

	mig "github.com/RedHatInsights/insights-operator-utils/migrations"
	types "github.com/RedHatInsights/insights-results-types"
	"github.com/rs/zerolog/log"
)

// mig0004UpdateEventTypeIDInReportedTable migration mopdifies the reported
// table, updating all records are updated to contain the value '1' in the
// event_type_id column if it is currently null.
var mig0004UpdateEventTypeIDInReportedTable = mig.Migration{
	StepUp: func(tx *sql.Tx, _ types.DBDriver) error {
		log.Debug().Msg("Executing mig0004UpdateEventTypeIDInReportedTable stepUp function")
		query := "UPDATE reported SET event_type_id = 1 WHERE event_type_id IS NULL"
		result, err := executeQuery(tx, query)
		if err != nil {
			return err
		}
		rows, _ := result.RowsAffected()
		log.Debug().Int64(rowsAffectedMessage, rows).
			Msg("event_type_id successfully set to 1 for all rows with current value = null")

		query = "ALTER TABLE reported ALTER COLUMN event_type_id SET NOT NULL"
		_, err = executeQuery(tx, query)
		if err == nil {
			log.Debug().Msg("Column event_type_id successfully set to not null")
		}
		return err
	},
	StepDown: func(tx *sql.Tx, _ types.DBDriver) error {
		log.Debug().Msg("Executing mig0004UpdateEventTypeIDInReportedTable stepDown function")
		query := "ALTER TABLE reported ALTER COLUMN event_type_id DROP NOT NULL"
		_, err := executeQuery(tx, query)
		if err != nil {
			return err
		}
		log.Debug().Msg("'not null' attribute' successfully removed from event_type_id column")
		query = "UPDATE reported SET event_type_id = NULL WHERE event_type_id = 1"
		result, err := executeQuery(tx, query)
		if err == nil {
			rows, _ := result.RowsAffected()
			log.Debug().Int64(rowsAffectedMessage, rows).
				Msg("event_type_id successfully set to null for all rows with current value = 1")
		}
		return err
	},
}
