package utils

import (
  "database/sql"
)


func CheckErr(err error, status int) int {
  if (err != nil) {
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
    return (400)
  }
  return (200)
}
