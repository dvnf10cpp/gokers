package controllers

import (
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
	"github.com/devanfer02/gokers/services"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
	Router *gin.Engine
	Service services.CourseService
}

func (courseCtr *CourseController) RegisterCourse(ctx *gin.Context) {
	var course models.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return 
	}

	response := courseCtr.Service.RegisterCourse(course)

	res.SendResponse(ctx, response)
}

func (courseCtr *CourseController) GetCourses(ctx *gin.Context) {
	var courses []models.Course

	typequery := ctx.Query("type")
	response := courseCtr.Service.GetCourses(courses, typequery)

	res.SendResponse(ctx, response)
}

func (courseCtr *CourseController) GetCourse(ctx *gin.Context) {
	var course models.Course

	id, err := helpers.GetParamID(ctx)

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	course.ID = id 

	response := courseCtr.Service.GetCourse(course)

	res.SendResponse(ctx, response)
}

func (courseCtr *CourseController) UpdateCourse(ctx *gin.Context) {
	var course models.Course

	if err := ctx.ShouldBindJSON(&course); err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad body request", err))
		return
	}

	id, err := helpers.GetParamID(ctx)

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	course.ID = id 

	response := courseCtr.Service.UpdateCourse(course)

	res.SendResponse(ctx, response)
}

func (courseCtr *CourseController) DeleteCourse(ctx *gin.Context) {
	id, err := helpers.GetParamID(ctx)

	if err != nil {
		res.SendResponse(ctx, res.CreateResponseErr(status.BadRequest, "bad param request", err))
		return
	}

	var course models.Course
	course.ID = id 

	response := courseCtr.Service.DeleteService(course)

	res.SendResponse(ctx, response)
}