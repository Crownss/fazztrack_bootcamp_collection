const pool = require("../config/connection")
const format = require('pg-format')

var model_movie = {}

model_movie.GetAll = (paginatoin) => {
    return new Promise((resolve, reject) => {
      let queryTemp = format(
         'SELECT * FROM movie'
       )
   
       if (paginatoin.order) {
         queryTemp = format(`${queryTemp} ORDER BY %s`, paginatoin.order)
       } else {
         queryTemp = format(`${queryTemp} ORDER BY id DESC`)
       }
   
       const offset = (paginatoin.page - 1) * paginatoin.limit
       queryTemp = format(
         `${queryTemp} LIMIT %s OFFSET %s`,
         paginatoin.limit,
         offset
       )
   
       const query = db.query(queryTemp)
         .then((data) => {
            resolve(data.rows)
         })
         .catch((err) => {
            reject(err)
         })
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