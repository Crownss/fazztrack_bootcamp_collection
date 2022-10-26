const pool = require("../config/connection")

var model_schedule = {}

model_schedule.GetAll = () => {
   return new Promise((resolve, reject) => {
      let queryTemp = format(
         'SELECT * FROM schedule'
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

model_schedule.GetSchedule = (data) => {
    return new Promise((resolve, reject) => {
        pool.query('SELECT * FROM schedule WHERE movie_id = $1', [data])
           .then((data) => {
              resolve(data.rows)
           })
           .catch((err) => {
              reject(err)
           })
     })
}

model_schedule.Post = (data) => {
    return new Promise((resolve, reject) => {
       pool.query(`INSERT INTO schedule (movie_id, price, location, date_start, date_end, premier, time) VALUES($1, $2, $3, $4, $5, $6, $7)`, [
         data.movie_id,
         data.price,
         data.location,
         data.date_start,
         data.date_end,
         data.premier,
         data.time,
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

model_schedule.Delete = (data) => {
    return new Promise((resolve, reject) => {
        pool.query('DELETE FROM schedule WHERE id = $1',[data])
           .then((data) => {
            resolve(`berhasil menghapus`)
           })
           .catch((err) => {
            reject(err)
           })
     })
}

model_schedule.Update = (data, params) => {
    return new Promise((resolve, reject) => {
       pool.query(`UPDATE schedule SET movie_id=$1, price=$2, location=$3, date_start=$4, date_end=$5, premier=$6, time=$7 WHERE movie_id=$8`, [
         data.movie_id,
         data.price,
         data.location,
         data.date_start,
         data.date_end,
         data.premier,
         data.time,
         params
       ])
          .then((data) => {
            resolve(`berhasil memperbarui schedule dengan id movie ${params}`)
          })
          .catch((err) => {
            reject(err)
            console.log(err)
          })
    })
}

module.exports = model_schedule