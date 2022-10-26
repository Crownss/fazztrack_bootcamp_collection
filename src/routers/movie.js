const express = require('express')
const ctrls_movie = require('../controllers/movie')
const Router = express.Router()
const upload = require('../middleware/uploader')


//Movie
Router.get('/', ctrls_movie.GETALL)
Router.get('/:name', ctrls_movie.GETNAME)
Router.post('/', upload.single('img'),ctrls_movie.POST)
Router.put('/:name',upload.single('img'), ctrls_movie.UPDATE)
Router.delete('/:name', ctrls_movie.DELETE)

module.exports = Router