// OWNER: stan

package schemas

import "github.com/spolu/settle/mint/model"

const (
	assetsSQL = `
CREATE TABLE IF NOT EXISTS assets(
  owner VARCHAR(256) NOT NULL,       -- owner address
  token VARCHAR(256) NOT NULL,       -- token
  created TIMESTAMP NOT NULL,
  propagation VARCHAR(32) NOT NULL,  -- propagation type (canonical, propagated)

  code VARCHAR(64) NOT NULL,    -- the code of the asset
  scale SMALLINT,               -- factor by which the asset native is scaled

  PRIMARY KEY(owner, token),
  CONSTRAINT assets_owner_code_u UNIQUE (owner, code) -- not propagated
);
`
)

func init() {
	model.RegisterSchema(
		"mint",
		"assets",
		assetsSQL,
	)
}
