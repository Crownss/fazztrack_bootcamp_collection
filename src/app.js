const express = require('express')
const Routers = express.Router()
const movie = require('./routers/movie')
const schedule = require('./routers/schedule')
const booking = require('./routers/booking')

Routers.use('/movie', movie)
Routers.use('/schedule', schedule)
Routers.use('/booking', booking)

module.exports = Routers