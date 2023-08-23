package services

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/devanfer02/gokers/configs"
	"github.com/devanfer02/gokers/helpers"
	"github.com/devanfer02/gokers/helpers/res"
	"github.com/devanfer02/gokers/helpers/status"
	"github.com/devanfer02/gokers/models"
)

type ClassService struct {
	Db *configs.Database
}

func (classSvc *ClassService) GetClasses(class []models.Class, queries []string) res.Response {
	var err error 

	if queries[0] != "" {
		err = classSvc.Db.PreloadByCondition([]string{"Course", "Lecturer"}, &class, "course_code = ?", queries[0]);
	} else if queries[1] != "" {
		err = classSvc.Db.PreloadByCondition([]string{"Course", "Lecturer"}, &class, "major = ?", queries[1]);
	} else {
		err = classSvc.Db.FindAll(&class);
	}

	if err != nil {
		return res.CreateResponseErr(status.ServerError,
			"internal server error", 
			err,
		)
	}

	return res.CreateResponse(status.Ok, "successfully fetch classes", class)
}

func (classSvc *ClassService) GetClass(class models.Class) res.Response {
	if err := classSvc.Db.PreloadByPK([]string{"Course","Lecturer"}, &class, class.ID); err != nil {
		return res.CreateResponseErr(status.BadRequest, "foreign entity assosiated not found body", err)
	}

	return res.CreateResponse(status.Ok, "test", nil)
}

func (classSvc *ClassService) RegisterClass(class models.Class) res.Response {
	var course models.Course
	var lecturer models.Lecturer

	if _, err := govalidator.ValidateStruct(&class); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if count := classSvc.Db.Count("class_name", models.Class{}, class.ClassName); count > 0 {
		return res.CreateResponseErr(status.Conflict, "class with that name already exist", nil)
	}

	if err := classSvc.Db.FirstByPK(&course, class.CourseID); err != nil {
		return res.CreateResponseErr(status.NotFound, "course_id not found", err)
	}

	if err := classSvc.Db.FirstByPK(&lecturer, class.LecturerID); err != nil {
		return res.CreateResponseErr(status.NotFound, "lecturer_id not found", err)
	}

	class.ID = helpers.GenerateUUID()

	if err := classSvc.Db.Create(&class); err != nil {
		return res.CreateResponseErr(status.ServerError, "internal sever error", err)
	}

	return res.CreateResponse(status.Ok, "class registered to system", gin.H {
		"record_data": class,
	})
}

func (classSvc *ClassService) UpdateClass(class models.Class) res.Response {
	if _, err := govalidator.ValidateStruct(class); err != nil {
		return res.CreateResponseErr(status.BadRequest, "bad body request", err)
	}

	if classSvc.Db.Update("id = ?", &class, class.ID) == 0 {
		return res.CreateResponseErr(
			status.ServerError, 
			"failed to update data",
			fmt.Errorf("data value doesnt change or data doesnt exist"),
		)
	}

	return res.CreateResponse(status.Ok, "class data update", gin.H {
		"record_data": class, 
	})
}

func (classSvc *ClassService) DeleteClass(class models.Class) res.Response {
	if classSvc.Db.Delete("id = ?", class, class.ID) == 0 {
		return res.CreateResponseErr(
			status.ServerError,
			"failed to delete data",
			fmt.Errorf("data didnt exist"),
		)
	}

	return res.CreateResponse(status.Ok, "successfully delete class", gin.H {
		"deleted_record_data": class,
	})
}