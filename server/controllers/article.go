package controllers

import (
  "math/rand"
  u "../utils"
  m "../models"
)

func ArticleCreate_c(token string, title string, content string, img_link string) (string, int) {
  // Check admin rights from stored token
  uid, admin := AdminRights(token)
  if (admin == false) {
    return "", 401
  }

  // Generate a hash for unique article id
  hash := make([]byte, 64)
  rand.Read(hash)
  aid := u.MD5Hash(hash)

  // Call the model man
  status := m.ArticleCreate_m(aid, uid, title, content, img_link)
  return aid, status
}

func ArticleRead_c(token string, aid string) (m.Article, int) {
  // Get article object from database (no authentification needed)
  art, status := m.ArticleRead_m(aid)
  // Check if article is owned (to allow Delete method front-side), token is not mandatory to read it
  if (token != "") {
    id, _ := u.DesDecrypt(token)
    if (id == art.Uid) {
      art.Owned = 1
    }
  }
  return art, status
}

func ArticleFetch_c(search string) ([]m.Article, int){
  // Fetch all articles (no authentification needed)
  arts, status := m.ArticleFetch_m(search)
  return arts, status
}

func ArticleUpdate_c(aid string, token string, title string, content string, img_link string) (string, int) {
  // Check admin rights from stored token
  uid, admin := AdminRights(token)
  if (admin == false) {
    return "", 401
  }

  // Check if article is owned
  owner, status := m.ArticleGetOwner_m(aid)
  if (status == 200 && owner != uid) {
    status = 401
  }
  if (status != 200) {
    return "", status
  }

  // Update the article
  status = m.ArticleUpdate_m(aid, title, content, img_link)
  return aid, status
}

func ArticleDelete_c(aid string, token string) (string, int) {
  // Check admin rights from stored token
  uid, admin := AdminRights(token)
  if (admin == false) {
    return "", 401
  }

  // Check if article is owned
  owner, status := m.ArticleGetOwner_m(aid)
  if (status == 200 && owner != uid) {
    status = 401
  }
  if (status != 200) {
    return "", status
  }

  // Delete the article
  status = m.ArticleDelete_m(aid)
  return aid, status
}
