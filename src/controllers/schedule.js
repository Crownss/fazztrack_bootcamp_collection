const model_schedule = require("../models/schedule")
const response = require("../helpers/response")

var ctrls_schedule = {}

ctrls_schedule.GETALL = async (req, res) => {
    try {
        const data = await model_schedule.GetAll()
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_schedule.GETMOVIE = async (req, res) => {
    try {
        const data = await model_schedule.GetMovie(req.params.movie_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_schedule.POST = async (req, res) => {
    try {
        const {movie_id, price, location, date_start, date_end, premier, time} = req.body
        const data = await model_schedule.Post({movie_id, price, location, date_start, date_end, premier, time})
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_schedule.UPDATE = async (req, res) => {
    try {
        const {movie_id, price, location, date_start, date_end, premier, time} = req.body
        const data = await model_schedule.Update({movie_id, price, location, date_start, date_end, premier, time}, req.params.movie_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_schedule.DELETE = async (req, res) => {
    try {
        const data = await model_schedule.Delete(req.params.movie_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}
module.exports = ctrls_schedule