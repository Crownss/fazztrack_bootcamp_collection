const express = require('express')
const ctrls_schedule = require('../controllers/schedule')
const Router = express.Router()


//Movie
Router.get('/', ctrls_schedule.GETALL)
Router.get('/:movie_id', ctrls_schedule.GETMOVIE)
Router.post('/', ctrls_schedule.POST)
Router.put('/:movie_id', ctrls_schedule.UPDATE)
Router.delete('/:movie_id', ctrls_schedule.DELETE)

module.exports = Router