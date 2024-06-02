package module

const createUsersTable string = `
  CREATE TABLE IF NOT EXISTS users (
  id INTEGER NOT NULL PRIMARY KEY,
  balance FLOAT NOT NULL
  );`
