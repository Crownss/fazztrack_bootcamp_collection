const express = require('express')
const ctrls_movie = require('../controllers/movie')
const Router = express.Router()


//Movie
Router.get('/', ctrls_movie.GETALL)
Router.get('/:name', ctrls_movie.GETNAME)
Router.post('/', ctrls_movie.POST)
Router.put('/:name', ctrls_movie.UPDATE)
Router.delete('/:name', ctrls_movie.DELETE)

module.exports = Router