package usecases

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	responses "github.com/fateqense/siga/app/responses/student"
	"github.com/fateqense/siga/core/gxstate"
	"github.com/fateqense/siga/core/network"
)

func (StudentUseCase) GetProfileAction(session string) (*responses.ProfileResponse, error) {
	authClient := network.NewAuthenticatedClient(session)

	res, err := authClient.Get(network.HOME_URL)
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

	averageGrade, err := strconv.Atoi(gxs.GetWithPrefix("vACD_ALUNOCURSOINDICEPR").(string))
	if err != nil {
		averageGrade = 0
	}

	progression, err := strconv.Atoi(gxs.GetWithPrefix("vACD_ALUNOCURSOINDICEPP").(string))
	if err != nil {
		progression = 0
	}

	profile := &responses.ProfileResponse{}
	profile.Ra = gxs.GetWithPrefix("vACD_ALUNOCURSOREGISTROACADEMICOCURSO").(string)
	profile.Name = strings.Replace(gxs.GetWithPrefix("vPRO_PESSOALNOME").(string), " -", "", 1)
	profile.PhotoURL = doc.Find(fmt.Sprintf("#%sFOTO > img", gxs.Prefix)).AttrOr("src", "")
	profile.Email = gxs.Get("vPRO_PESSOALEMAIL").(string)
	profile.InstitutionalEmail = gxs.GetWithPrefix("vINSTITUCIONALFATEC").(string)
	profile.AverageGrade = averageGrade
	profile.Progression = progression

	return profile, nil
}
