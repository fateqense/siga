package usecases

import (
	"github.com/PuerkitoBio/goquery"
	responses "github.com/fateqense/siga/app/responses/student"
	"github.com/fateqense/siga/core/gxstate"
	"github.com/fateqense/siga/core/network"
	"github.com/fateqense/siga/utils"
)

func (StudentUseCase) GetHistoryAction(session string) ([]responses.HistoryResponse, error) {
	authClient := network.NewAuthenticatedClient(session)

	res, err := authClient.Get(network.HISTORY_URL)
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

	completeHistory := make([]responses.HistoryResponse, 0)

	disciplines := utils.FromInterfaceToSliceMap[string, interface{}](gxs.Get("vALU_ALUNONOTAS_SDT"))
	for _, discipline := range disciplines {
		history := responses.HistoryResponse{}

		history.Cod = discipline["ACD_DisciplinaSigla"].(string)
		history.DisiplineName = discipline["ACD_DisciplinaNome"].(string)
		history.Description = discipline["GER_TipoObservacaoHistoricoDescricao"].(string)
		history.FinalGrade = discipline["ACD_AlunoHistoricoItemMediaFinal"].(float64)
		history.TotalAbsences = discipline["ACD_AlunoHistoricoItemQtdFaltas"].(float64)
		history.PresenceFrequency = discipline["ACD_AlunoHistoricoItemFrequencia"].(float64)
		history.RenunciationAt = discipline["ACD_AlunoHistoricoItemDesistenciaData"].(string)
		history.IsApproved = discipline["ACD_AlunoHistoricoItemAprovada"].(float64) == 1

		completeHistory = append(completeHistory, history)
	}

	return completeHistory, nil
}
