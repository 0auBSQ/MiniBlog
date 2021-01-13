package models

import (
  u "../utils"
  "log"
)

type Article struct {
  Id string
  Uid string
  Author string
  Title string
  Contents string
  Img_link string
  Creation_date string
  Owned int
}

func ArticleCreate_m(aid string, uid string, title string, content string, img_link string) int {
  db = Db_m()
  _, err := db.Exec("INSERT INTO articles(id, uid, titre, contents, img_link) VALUES ($1, $2, $3, $4, $5);", aid, uid, title, content, img_link)
  return (u.CheckErr(err, 200))
}

func ArticleUpdate_m(aid string, title string, content string, img_link string) int {
  db = Db_m()
  _, err := db.Exec("UPDATE articles SET titre=$2, contents=$3, img_link=$4 WHERE id=$1;", aid, title, content, img_link)
  return (u.CheckErr(err, 200))
}

func ArticleDelete_m(aid string) int {
  db = Db_m()
  _, err := db.Exec("DELETE FROM articles WHERE id=$1;", aid)
  return (u.CheckErr(err, 200))
}

func ArticleRead_m(aid string) (art Article, status int) {
  db = Db_m()
  status = 404
  rows, err := db.Query("SELECT a.id, uid, titre, contents, img_link, creation_date, u.username FROM articles a INNER JOIN users u ON a.uid = u.id WHERE a.id=$1", aid)
  if (err != nil) {
    log.Println(err)
    return art, 500
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&art.Id, &art.Uid, &art.Title, &art.Contents, &art.Img_link, &art.Creation_date, &art.Author)
    if (err != nil) {
      log.Println(err)
      return art, 500
    }
    status = 200
  }
  return art, status
}

func ArticleGetOwner_m(aid string) (string, int) {
  art, status := ArticleRead_m(aid)
  return art.Uid, status
}

func ArticleFetch_m(search string) (arts []Article, status int) {
  db = Db_m()
  var art Article
  rows, err := db.Query("SELECT a.id, uid, titre, contents, img_link, creation_date, u.username FROM articles a INNER JOIN users u ON a.uid = u.id WHERE a.titre LIKE '%' || $1 || '%'", search)
  if (err != nil) {
    log.Println(err)
    return arts, 500
  }
  defer rows.Close()
  for rows.Next() {
    err := rows.Scan(&art.Id, &art.Uid, &art.Title, &art.Contents, &art.Img_link, &art.Creation_date, &art.Author)
    if (err != nil) {
      log.Println(err)
      return arts, 500
    }
    arts = append(arts, art)
  }
  return arts, 200
}
