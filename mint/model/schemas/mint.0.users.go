package schemas

import "github.com/spolu/settle/mint/model"

const (
	usersSQL = `
CREATE TABLE IF NOT EXISTS users(
  token VARCHAR(256) NOT NULL,
  created TIMESTAMP NOT NULL,

  username VARCHAR(256) NOT NULL,
  password_hash VARCHAR(256) NOT NULL,

  PRIMARY KEY(token),
  CONSTRAINT users_usernameu UNIQUE (username)
);
`
)

func init() {
	model.RegisterSchema(
		"mint",
		"users",
		usersSQL,
	)
}
