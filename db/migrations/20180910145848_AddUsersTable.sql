
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id 			serial			PRIMARY KEY,
  email 		text			NOT NULL,
  first_name 	text			NOT NULL,
  last_name 	text			NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
