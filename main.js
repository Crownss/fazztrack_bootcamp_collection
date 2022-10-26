require('dotenv').config()
const express = require('express')
const server = express()
const app = require('./src/app')
const pool = require('./src/config/connection')
async function Start() {
   try {
      await pool.connect()
      server.use(express.urlencoded({ extended: true }))
      server.use(express.json())
      server.use(app)
      if(process.env.RUN_PORT != ""){
         return server.listen(process.env.RUN_PORT, () => {
            console.log('running on port',process.env.RUN_PORT)
         })
      }
      return server.listen(8000, () => {
         console.log('running on port 8000')
      })
   } catch (error) {
      return console.log(error)
   }
}
Start()