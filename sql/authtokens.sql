-- drop the table if it exists
DROP TABLE IF EXISTS authtokens;

-- and create the new one
CREATE TABLE authtokens(
   id          INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
   token       VARCHAR( 32 ) UNIQUE KEY NOT NULL DEFAULT '',
   whom        VARCHAR( 32 ) NOT NULL DEFAULT '',
   what        VARCHAR( 32 ) NOT NULL DEFAULT '',
   create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) CHARACTER SET utf8 COLLATE utf8_bin;
