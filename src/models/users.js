const pool = require("../config/connection")

var model_users = {}

model_users.GetAll = () =>{
    return new Promise((resolve, reject) => {
        pool.query('SELECT * FROM users')
           .then((data) => {
              resolve(data.rows)
           })
           .catch((err) => {
              reject(err)
           })
     })
}

model_users.GetUsername = (username) =>{
   return new Promise((resolve, reject) => {
      pool.query('SELECT * FROM users WHERE username = $1', [username])
         .then((data) => {
            resolve(data.rows)
         })
         .catch((err) => {
            reject(err)
         })
   })
}

model_users.Post = (obj) =>{
   return new Promise((resolve, reject) => {
      pool.query(`INSERT INTO users (profile, email, username, password, no_telp, is_admin) VALUES($1, $2, $3, $4, $5, $6)`, [
         obj.profile,
         obj.email,
         obj.username,
         obj.password_hash,
         obj.no_telp,
         obj.is_admin
      ])
         .then((data) => {
            resolve()
         })
         .catch((err) => {
            reject(err)
         })
   })
}

model_users.Delete = (username) =>{
   return new Promise((resolve, reject) => {
      pool.query('DELETE FROM users WHERE username = $1',[username])
         .then((data) => {
            resolve()
         })
         .catch((err) => {
            reject(err)
         })
   })
}

model_users.Update = (obj, username) =>{
   return new Promise((resolve, reject) => {
      pool.query(`UPDATE users SET profile = COALESCE(NULLIF($1, ''), profile), email = COALESCE(NULLIF($2, ''), email), username = COALESCE(NULLIF($3, ''), username), password = COALESCE(NULLIF($4, ''), password), no_telp = COALESCE(NULLIF($5, ''), no_telp), is_admin = COALESCE(NULLIF($6, false), is_admin) WHERE username = $7 RETURNING *`,
      [
         obj.profile,
         obj.email,
         obj.username,
         obj.password_hash,
         obj.no_telp,
         obj.is_admin,
         username
      ])
         .then((data) => {
            resolve()
         })
         .catch((err) => {
            reject(err)
         })
   })
}

module.exports = model_users