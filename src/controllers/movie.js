const model_movie = require("../models/movie")
const response = require("../helpers/response")
var ctrls_movie = {}

ctrls_movie.GETALL = async (req, res) => {
    try {
        const {orderby} = req.query
        if(orderby != ''){
            const data = await model_movie.GetAll(orderby)
            return response(res, 200, data)
        }else{
            const data = await model_movie.GetAll()
            return response(res, 200, data)
        }
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
        const {img, name, synopsis, category, director, casts, release, duration} = req.body
        const data = await model_movie.Post({img, name, synopsis, category, director, casts, release, duration})
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_movie.UPDATE = async (req, res) => {
    try {
        const {img, name, synopsis, category, director, casts, release, duration} = req.body
        const data = await model_movie.Update({img, name, synopsis, category, director, casts, release, duration}, req.params.name)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_movie.DELETE = async (req, res) => {
    try {
        const data = await model_movie.Delete(req.params.name)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 400, error, true)
    }
}
module.exports = ctrls_movie