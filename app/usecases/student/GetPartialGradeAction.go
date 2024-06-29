package usecases

import (
	"github.com/PuerkitoBio/goquery"
	responses "github.com/fateqense/siga/app/responses/student"
	"github.com/fateqense/siga/core/gxstate"
	"github.com/fateqense/siga/core/network"
	"github.com/fateqense/siga/utils"
)

func (StudentUseCase) GetPartialGradesAction(session string) ([]*responses.PartialGradeResponse, error) {
	authClient := network.NewAuthenticatedClient(session)

	res, err := authClient.Get(network.PARTIAL_GRADES_URL)
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

	grades := make([]*responses.PartialGradeResponse, 0)

	disciplines := utils.FromInterfaceToSliceMap[string, interface{}](gxs.Get("vACD_ALUNONOTASPARCIAISRESUMO_SDT"))
	for _, discipline := range disciplines {
		grade := &responses.PartialGradeResponse{}

		grade.Cod = discipline["ACD_DisciplinaSigla"].(string)
		grade.DisiplineName = discipline["ACD_DisciplinaNome"].(string)
		grade.AverageGrade = discipline["ACD_AlunoHistoricoItemMediaFinal"].(float64)

		dates := utils.FromInterfaceToSliceMap[string, interface{}](discipline["Datas"])
		for _, date := range dates {
			exam := &responses.Exam{}

			exam.Title = date["ACD_PlanoEnsinoAvaliacaoTitulo"].(string)
			exam.StartsAt = date["ACD_PlanoEnsinoAvaliacaoDataPrevista"].(string)

			exams := utils.FromInterfaceToSliceMap[string, int](date["Avaliacoes"])
			if len(exams) > 0 {
				exam.Grade = exams[0]["ACD_PlanoEnsinoAvaliacaoParcialNota"]
			} else {
				exam.Grade = 0
			}

			grade.Exams = append(grade.Exams, exam)
		}

		grades = append(grades, grade)
	}

	return grades, nil
}
