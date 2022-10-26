const express = require('express')
const Routers = express.Router()
const movie = require('./routers/movie')
const schedule = require('./routers/schedule')
const booking = require('./routers/booking')
const auth = require('./routers/auth')

Routers.use('/movie', movie)
Routers.use('/schedule', schedule)
Routers.use('/booking', booking)
Routers.use('/auth', auth)

module.exports = Routers