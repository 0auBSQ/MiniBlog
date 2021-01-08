package controllers

import (
  "math/rand"
  u "../utils"
  m "../models"
)

func CommentCreate_c(token string, aid string, content string) int {
  // Check if user is logged
  uid, is_user := UserRights(token)
  if (is_user == false) {
    return 401
  }

  // Check if user is banned
  user, status := m.Userinfo_m(uid)
  if (user.Is_banned == 1) {
    return 403
  }

  // Generate a hash for unique comment id
  hash := make([]byte, 64)
  rand.Read(hash)
  qid := u.MD5Hash(hash)

  // Call the model man
  status = m.CommentCreate_m(qid, aid, uid, content)
  return status
}

func CommentFetch_c(token string, aid string) (coms []m.Comment, status int) {
  // Retrieve uid to check if comment owned (for front purpose)
  uid, admin := AdminRights(token)

  // Retrieve all comments from the given post
  coms, status = m.CommentFetch_m(uid, aid, admin)
  return coms, status
}

func CommentDelete_c(token string, qid string) int {
  // Check if user is logged, admin can delete anyone comments
  uid, is_user := UserRights(token)
  _, is_admin := AdminRights(token)
  if (is_user == false) {
    return 401
  }

  // Retrieve comment owner
  owner, status := m.CommentGetOwner_m(qid)
  if (status == 200 && owner != uid && is_admin == false) {
    status = 401
  }
  if (status != 200) {
    return status
  }

  // Delete the comment
  status = m.CommentDelete_m(qid)
  return status
}

func CommentUpdate_c(token string, qid string, content string) int {
  // Check if user is logged
  uid, is_user := UserRights(token)
  if (is_user == false) {
    return 401
  }

  // Check if user is banned
  user, status := m.Userinfo_m(uid)
  if (user.Is_banned == 1) {
    return 403
  }

  // Retrieve comment owner
  owner, status := m.CommentGetOwner_m(qid)
  if (status == 200 && owner != uid) {
    status = 401
  }
  if (status != 200) {
    return status
  }

  // Update the comment
  status = m.CommentUpdate_m(qid, content)
  return status
}
