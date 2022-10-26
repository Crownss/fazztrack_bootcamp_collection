require('dotenv').config()
const express = require('express')
const server = express()
const app = require('./src/app')
const pool = require('./src/config/connection')
const multer = require('multer')
const upload = multer({dest: 'static/images'})
const cors = require('cors')

async function Start() {
   try {
      await pool.connect()
      server.use(express.urlencoded({ extended: true }))
      server.use(express.json())
      server.use('/static/images', express.static('static/images'))
      const whitelist = [`http://localhost:${process.env.APP_PORT}`]
      const corsOptionsDelegate = function (req, callback) {
        let corsOptions
        if (whitelist.indexOf(req.header('Origin')) !== -1) {
          corsOptions = { origin: true }
        } else {
          corsOptions = { origin: false }
        }
        callback(null, corsOptions)
      }
      server.use('/api/v1', cors(corsOptionsDelegate), app)
      if(process.env.RUN_PORT != ""){
         return server.listen(process.env.RUN_PORT, () => {
            console.log(`running on port ${process.env.RUN_PORT}`)
         })
      }
      return server.listen(8000, () => {
         console.log(`running on port ${8000}`)
      })
   } catch (error) {
      return console.log(error)
   }
}
Start()