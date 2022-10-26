const pool = require("../config/connection")

var model_booking = {}

model_booking.GetAll = () => {
   return new Promise((resolve, reject) => {
      let queryTemp = format(
         'SELECT * FROM booking'
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

model_booking.GetSchedule = (data) => {
    return new Promise((resolve, reject) => {
        pool.query('SELECT * FROM booking WHERE schedule_id = $1', [data])
           .then((data) => {
              resolve(data.rows)
           })
           .catch((err) => {
              reject(err)
           })
     })
}

model_booking.Post = (data) => {
    return new Promise((resolve, reject) => {
       pool.query(`INSERT INTO booking (schedule_id, seats) VALUES($1, $2)`, [
         data.schedule_id,
         data.seats,
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

model_booking.Delete = (data) => {
    return new Promise((resolve, reject) => {
        pool.query('DELETE FROM booking WHERE id = $1',[data])
           .then((data) => {
            resolve(`berhasil menghapus`)
           })
           .catch((err) => {
            reject(err)
           })
     })
}

model_booking.Update = (data, params) => {
    return new Promise((resolve, reject) => {
       pool.query(`UPDATE booking SET schedule_id=$1, seats=$2 WHERE schedule_id=$3`, [
         data.schedule_id,
         data.seats,
         params
       ])
          .then((data) => {
            resolve(`berhasil memperbarui booking dengan id schedule ${params}`)
          })
          .catch((err) => {
            reject(err)
            console.log(err)
          })
    })
}

module.exports = model_booking