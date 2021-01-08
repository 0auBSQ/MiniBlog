package models

import (
  u "../utils"
  "log"
)

type Comment struct {
  Id string
  Uid string
  Aid string
  Author string
  Contents string
  Creation_date string
  Owned int
}

func CommentCreate_m(qid string, aid string, uid string, content string) int {
  db = Db_m()
  _, err := db.Exec("INSERT INTO comment(id, uid, aid, contents) VALUES ($1, $2, $3, $4);", qid, uid, aid, content)
  return (u.CheckErr(err, 200))
}

func CommentDelete_m(qid string) int {
  db = Db_m()
  _, err := db.Exec("DELETE FROM comment WHERE id=$1;", qid)
  return (u.CheckErr(err, 200))
}

func CommentUpdate_m(qid string, content string) int {
  db = Db_m()
  _, err := db.Exec("UPDATE comment SET contents=$2 WHERE id=$1;", qid, content)
  return (u.CheckErr(err, 200))
}

func CommentFetch_m(uid string, aid string, admin bool) (coms []Comment, status int) {
  db = Db_m()
  var com Comment
  rows, err := db.Query("SELECT c.id, c.uid, c.aid, c.contents, c.creation_date, u.username FROM comment c INNER JOIN articles a ON c.aid = a.id INNER JOIN users u ON c.uid = u.id WHERE c.aid=$1", aid)
  if (err != nil) {
    log.Println(err)
    return coms, 500
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&com.Id, &com.Uid, &com.Aid, &com.Contents, &com.Creation_date, &com.Author)
    // If admin and post not owned = Ban + Delete icons, if post owned = Delete icons, else no icon
    if (admin == true && uid != com.Uid) {
      com.Owned = 2
    } else if (uid == com.Uid) {
      com.Owned = 1
    }
    if (err != nil) {
      log.Println(err)
      return coms, 500
    }
    coms = append(coms, com)
  }
  return coms, 200
}

func CommentGetOwner_m(qid string) (owner string, status int) {
  db = Db_m()
  status = 404
  rows, err := db.Query("SELECT uid FROM comment WHERE id=$1", qid)
  if (err != nil) {
    log.Println(err)
    return owner, 500
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&owner)
    if (err != nil) {
      log.Println(err)
      return owner, 500
    }
    status = 200
  }
  return owner, status
}
