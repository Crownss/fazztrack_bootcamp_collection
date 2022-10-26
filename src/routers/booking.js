const express = require('express')
const ctrls_booking = require('../controllers/booking')
const Router = express.Router()


//Movie
Router.get('/', ctrls_booking.GETALL)
Router.get('/:schedule_id', ctrls_booking.GETSCHEDULE)
Router.post('/', ctrls_booking.POST)
Router.put('/:schedule_id', ctrls_booking.UPDATE)
Router.delete('/:schedule_id', ctrls_booking.DELETE)

module.exports = Router