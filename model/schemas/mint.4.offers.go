package schemas

import "github.com/spolu/peer-currencies/model"

const (
	offersSQL = `
CREATE TABLE IF NOT EXISTS offers(
  token VARCHAR(256) NOT NULL,
  created TIMESTAMP NOT NULL,
  livemode BOOL NOT NULL,

  owner VARCHAR(256) NOT NULL,       -- the offer's owner's address
  base_asset VARCHAR(256) NOT NULL,  -- the base asset
  quote_asset VARCHAR(256) NOT NULL, -- the quote asset

  type VARCHAR(32) NOT NULL,         -- the type (bid, ask)

  base_price VARCHAR(64) NOT NULL,   -- the base asset price
  quote_price VARCHAR(64) NOT NULL,  -- the quote asset price
  amount VARCHAR(64) NOT NULL,       -- the amount of quote asset offered

  status VARCHAR(32) NOT NULL,       -- the status (active, closed)

  PRIMARY KEY(token)
);
`
)

func init() {
	model.RegisterSchema(
		"mint",
		"offers",
		offersSQL,
	)
}
