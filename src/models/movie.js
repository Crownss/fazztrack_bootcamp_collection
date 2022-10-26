const pool = require("../config/connection")

var model_movie = {}

model_movie.GetAll = (sortby) => {
    return new Promise((resolve, reject) => {
      if(sortby != ''){
         pool.query(`SELECT * FROM movie ORDER BY ${sortby} ASC`)
         .then((data) => {
            resolve(data.rows)
         })
         .catch((err) => {
            reject(err)
         })
      }else{
        pool.query('SELECT * FROM movie')
           .then((data) => {
              resolve(data.rows)
           })
           .catch((err) => {
              reject(err)
           })
         }
     })
}

model_movie.GetName = (data) => {
    return new Promise((resolve, reject) => {
        pool.query('SELECT * FROM movie WHERE name = $1', [data])
           .then((data) => {
              resolve(data.rows)
           })
           .catch((err) => {
              reject(err)
           })
     })
}

model_movie.Post = (data) => {
    return new Promise((resolve, reject) => {
       pool.query(`INSERT INTO movie (img, "name", synopsis, category, director, casts, release, duration) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`, [
         data.img,
         data.name,
         data.synopsis,
         data.category,
         data.director,
         data.casts,
         data.release,
         data.duration
       ])
          .then((data) => {
             resolve('berhasil memposting')
          })
          .catch((err) => {
             reject(err)
             console.log(err)
          })
    })
}

model_movie.Delete = (data) => {
    return new Promise((resolve, reject) => {
        pool.query('DELETE FROM movie WHERE name = $1',[data])
           .then((data) => {
              resolve(`berhasil menghapus movie nama ${data.name}`)
           })
           .catch((err) => {
              reject(err)
           })
     })
}

model_movie.Update = (data, params) => {
    return new Promise((resolve, reject) => {
       pool.query(`UPDATE movie SET img=$1, name=$2, synopsis=$3, category=$4, director=$5, casts=$6, release=$7, duration=$8 WHERE name=$9`, [
         data.img,
         data.name,
         data.synopsis,
         data.category,
         data.director,
         data.casts,
         data.release,
         data.duration,
         params
       ])
          .then((data) => {
            resolve(`berhasil memperbarui movie dengan nama ${params}`)
          })
          .catch((err) => {
            reject(err)
            console.log(err)
          })
    })
}

module.exports = model_movie