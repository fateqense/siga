package usecases

import (
	"github.com/PuerkitoBio/goquery"
	responses "github.com/fateqense/siga/app/responses/student"
	"github.com/fateqense/siga/core/gxstate"
	"github.com/fateqense/siga/core/network"
	"github.com/fateqense/siga/utils"
)

func (StudentUseCase) GetPartialAbsencesAction(session string) ([]responses.PartialAbsenceRespose, error) {
	authClient := network.NewAuthenticatedClient(session)

	res, err := authClient.Get(network.PARTIAL_ABSENCES_URL)
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

	absences := make([]responses.PartialAbsenceRespose, 0)

	disciplines := utils.FromInterfaceToSliceMap[string, interface{}](gxs.Get("vFALTAS"))
	for _, discipline := range disciplines {
		absence := responses.PartialAbsenceRespose{}

		absence.Cod = discipline["ACD_DisciplinaSigla"].(string)
		absence.DisiplineName = discipline["ACD_DisciplinaNome"].(string)
		absence.TotalPresences = discipline["TotalPresencas"].(float64)
		absence.TotalAbsences = discipline["TotalAusencias"].(float64)

		lessons := utils.FromInterfaceToSliceMap[string, interface{}](discipline["Aulas"])
		for _, lesson := range lessons {
			lsn := responses.Lesson{}

			lsn.Title = lesson["ACD_PlanoEnsinoConteudoTituloAula"].(string)
			lsn.Date = lesson["ACD_PlanoEnsinoConteudoDataAula"].(string)
			lsn.Presences = lesson["Presencas"].(float64)
			lsn.Absences = lesson["Ausencias"].(float64)

			absence.Lessons = append(absence.Lessons, lsn)
		}

		absences = append(absences, absence)
	}

	return absences, nil
}
