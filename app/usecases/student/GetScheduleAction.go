package usecases

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	responses "github.com/fateqense/siga/app/responses/student"
	"github.com/fateqense/siga/core/gxstate"
	"github.com/fateqense/siga/core/network"
	"github.com/fateqense/siga/utils"
)

func (StudentUseCase) GetScheduleAction(session string) ([][]responses.ScheduleResponse, error) {
	authClient := network.NewAuthenticatedClient(session)

	res, err := authClient.Get(network.SCHEDULE_URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	doc, _ := goquery.NewDocumentFromReader(res.Body)
	gxstateRaw, _ := doc.Find(`[name="GXState"]`).Attr("value")

	gxs, err := gxstate.NewGXState(gxstateRaw)
	if err != nil {
		return nil, err
	}

	disciplines := make(map[string]responses.Discipline)
	rawDisciplines := utils.FromInterfaceToSliceMap[string, interface{}](gxs.Get("vALU_ALUNOHISTORICOITEM_SDT"))
	for _, discipline := range rawDisciplines {
		dsc := responses.Discipline{}

		splited := strings.Split(discipline["ACD_DisciplinaNome"].(string), "<br&gt;")

		dsc.Cod = strings.Trim(discipline["ACD_DisciplinaSigla"].(string), " ")
		dsc.Name = splited[0]
		dsc.TeacherName = strings.Trim(discipline["Pro_PessoalNome"].(string), " ")
		dsc.Class = discipline["ACD_TurmaLetra"].(string)

		durationStr := strings.Replace(splited[1], "hs/aula", "", 1)
		duration, err := strconv.Atoi(durationStr)
		if err != nil {
			duration = 0
		}

		dsc.Workload = duration * 20
		dsc.TotalAbsencesAllowed = float64(dsc.Workload) - float64(dsc.Workload)*0.75

		disciplines[dsc.Cod] = dsc
	}

	schedule := make([][]responses.ScheduleResponse, 0)
	today := time.Now()

	for i := range 6 {
		lessons := make([]responses.ScheduleResponse, 0)

		gridTag := fmt.Sprintf(`[name="Grid%dContainerDataV"]`, i+2)
		data := doc.Find(gridTag).AttrOr("value", "[]")

		var dayLessons [][]string
		if err := json.Unmarshal([]byte(data), &dayLessons); err != nil {
			continue
		}

		week := i + 1

		for _, lesson := range dayLessons {
			discipline := responses.ScheduleResponse{}

			discipline.Cod = lesson[2]

			horaries := strings.Split(lesson[1], "-")
			startsAt := horaries[0]
			endsAt := horaries[1]

			hours := time.Duration(((week - int(today.Weekday()) + 7) % 7) * 24)

			startsAtDate := today.Add(time.Hour * hours)
			endsAtDate := today.Add(time.Hour * hours)

			discipline.StartsAt = fmt.Sprintf("%s %s", startsAtDate.Format(time.DateOnly), startsAt)
			discipline.EndsAt = fmt.Sprintf("%s %s", endsAtDate.Format(time.DateOnly), endsAt)

			discipline.Discipline = disciplines[discipline.Cod]

			lessons = append(lessons, discipline)
		}

		schedule = append(schedule, lessons)
	}

	return schedule, nil
}
