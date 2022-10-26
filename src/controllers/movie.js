const response = require("../helpers/response")
const model_movie = require("../models/movie")
const verify = require('jsonwebtoken')
const model_users = require('../models/users')
const validate = require("../middleware/validate_users")


var ctrls_movie = {}

ctrls_movie.GETALL = async (req, res) => {
    try {
        const pagination = {
            page: req.query.page || 1,
            limit: req.query.limit || 5,
            order: req.query.order
        }
        const rows = await model_movie.countRows()
        console.log(rows)
        const data = await model_movie.GetAll(pagination)
        return response(res, 200, data)
        } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_movie.GETNAME = async (req, res) => {
    try {
        const data = await model_movie.GetName(req.params.name)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_movie.POST = async (req, res) => {
    try {
        const {name, synopsis, category, director, casts, release, duration} = req.body
        const img = 'http://localhost:8000/'+req.file.path
        const data = await model_movie.Post({img, name, synopsis, category, director, casts, release, duration})
        return validate(req, res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_movie.UPDATE = async (req, res) => {
    try {
        const {name, synopsis, category, director, casts, release, duration} = req.body
        const img = 'http://localhost:8000/'+req.file.path
        const data = await model_movie.Update({img, name, synopsis, category, director, casts, release, duration}, req.params.name)
        return validate(req, res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_movie.DELETE = async (req, res) => {
    try {
        const data = await model_movie.Delete(req.params.name)
        return validate(req, res, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}
module.exports = ctrls_movie