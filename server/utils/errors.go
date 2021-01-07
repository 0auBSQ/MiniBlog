package utils

import (
  "database/sql"
)

func CheckErr(err error, status int) int {
  if (err != nil) {
    // Handles most internal server errors
    return (500)
  }
  return (status)
}

func CheckCountUnique(rows *sql.Rows) (count int) {
  for rows.Next() {
    err := rows.Scan(&count)
    if (CheckErr(err, 200) != 200) {
      return (500)
    }
  }
  if (count > 0) {
    // Value (ex : user account) already exists : bad request
    return (400)
  }
  return (200)
}
