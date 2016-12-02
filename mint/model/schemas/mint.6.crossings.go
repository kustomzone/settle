// OWNER: stan

package schemas

import "github.com/spolu/settle/mint/model"

const (
	crossingsSQL = `
CREATE TABLE IF NOT EXISTS crossings(
  owner VARCHAR(256) NOT NULL,  -- owner address
  token VARCHAR(256) NOT NULL,  -- token
  created TIMESTAMP NOT NULL,

  offer VARCHAR(256) NOT NULL,                     -- offer id
  amount VARCHAR(64) NOT NULL CHECK (amount > 0),  -- crossing amount

  status VARCHAR(32) NOT NULL,       -- status (reserved, settled, canceled)
  txn VARCHAR(256) NOT NULL,         -- transaction id
  hop SMALLINT NOT NULL,             -- transaction hop

  PRIMARY KEY(owner, token)
);
`
)

func init() {
	model.RegisterSchema(
		"mint",
		"crossings",
		crossingsSQL,
	)
}
