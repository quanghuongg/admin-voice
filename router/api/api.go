package api

import (
	"crypto"
	"fmt"
	"github.com/labstack/echo"
	"github.com/vtcc/voice-note-admin/model"
	"github.com/vtcc/voice-note-admin/router/base"
	"net/http"
	"strconv"
	"time"
)

func UserStats(c echo.Context) error {

	data := base.TableData{}

	draw, _ := strconv.Atoi(c.QueryParam("draw"))
	start, _ := strconv.Atoi(c.QueryParam("start"))
	size, _ := strconv.Atoi(c.QueryParam("length"))

	query := c.QueryParam("search[value]")

	var user model.User
	var record model.Record

	users, total, err := user.GetAllUsers(start, size, query)
	//total, _ := user.Count()

	if err != nil || users == nil {
		users = make([]model.User, 0)
	} else {
		for k,u := range users {
			// get User
			totalRecord, _ := record.CountByUserId(u.Id)
			users[k].TotalRecord = totalRecord
		}
	}
	data.Data = users

	data.Draw = draw
	data.RecordsTotal = total
	data.RecordsFiltered = total

	return c.JSON(http.StatusOK, data)
}

func RecordStats(c echo.Context) error {

	data := base.TableData{}

	draw, _ := strconv.Atoi(c.QueryParam("draw"))
	start, _ := strconv.Atoi(c.QueryParam("start"))
	size, _ := strconv.Atoi(c.QueryParam("length"))

	query := c.QueryParam("search[value]")
	uid := c.QueryParam("uid")

	var user model.User
	var record model.Record
	var media model.Media
	var feedback model.FeedBack

	records, total, err := record.GetAllRecords(start, size, query, uid)
	//total, _ := record.Count()

	if err != nil || records == nil {
		records = make([]model.Record, 0)
	} else {
		for k,r := range records {
			// get User
			u, _ := user.GetUserById(r.UserId)
			records[k].User = u

			// get Media
			m, _ := media.GetMediaByFileId(r.Id)
			records[k].Duration = m.Duration

			//Calc Hash Id
			records[k].HashId = base.EncodeHashId(r.Id)

			// get Feedback
			f, _ := feedback.GetFeedbackByFileId(r.Id)
			records[k].FeedBack = f

		}
	}
	data.Data = records

	data.Draw = draw

	data.RecordsTotal = total
	data.RecordsFiltered = total

	return c.JSON(http.StatusOK, data)
}

func RecordDetail(c echo.Context) error {

	data := base.JsonData{
		Err: 1,
	}

	fileId, _ := strconv.Atoi(c.QueryParam("fid"))

	var record model.Record

	r, err := record.GetByFileId(fileId)
	if err == nil {
		var media model.Media
		if m, err := media.GetMediaByFileId(r.Id); err == nil {
			t,_ := time.Parse("2006-01-02 15:04:05", r.Created)
			h := crypto.SHA1.New()
			h.Write([]byte(strconv.Itoa(m.Id)))
			r.AudioUrl = fmt.Sprintf("%d/%d/%d/%x.wav", t.Year(), int(t.Month()), t.Day(), h.Sum(nil))
		}

		data.Err = 0
		data.Data = r
	}

	return c.JSON(http.StatusOK, data)
}