const model_booking = require("../models/booking")

var ctrls_booking = {}

ctrls_booking.GETALL = async (req, res) => {
    try {
        const pagination = {
            page: req.query.page || 1,
            limit: req.query.limit || 5,
            order: req.query.order
        }
        const data = await model_booking.GetAll(pagination)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_booking.GETSCHEDULE = async (req, res) => {
    try {
        const data = await model_booking.GetSchedule(req.params.schedule_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}

ctrls_booking.POST = async (req, res) => {
    try {
        const {schedule_id, seats} = req.body
        const data = await model_booking.Post({schedule_id, seats})
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_booking.UPDATE = async (req, res) => {
    try {
        const {schedule_id, seats} = req.body
        const data = await model_booking.Update({schedule_id, seats}, req.params.schedule_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 500, error, true)
    }
}

ctrls_booking.DELETE = async (req, res) => {
    try {
        const data = await model_booking.Delete(req.params.schedule_id)
        return response(res, 200, data)
    } catch (error) {
        return response(res, 404, error, true)
    }
}
module.exports = ctrls_booking