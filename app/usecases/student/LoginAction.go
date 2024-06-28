package usecases

import (
	"errors"
	"net/http"
	"net/url"

	requests "github.com/fateqense/siga/app/requests/student"
	responses "github.com/fateqense/siga/app/responses/student"
	"github.com/fateqense/siga/core/network"
	"github.com/fateqense/siga/utils"
)

func (StudentUseCase) LoginAction(req requests.LoginRequest) (*responses.LoginResponse, error) {
	client := network.NewClient()

	form := url.Values{}
	form.Set("vSIS_USUARIOID", req.Username)
	form.Set("vSIS_USUARIOSENHA", req.Password)
	form.Set("BTCONFIRMA", "Confirmar")
	form.Set("GXState", network.LOGIN_GXSTATE)

	res, err := client.PostForm(network.LOGIN_URL, form)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusSeeOther {
		return nil, errors.New("invalid credentials")
	}

	cookies := utils.ParseSetCookieHeaders(res.Header.Values("Set-Cookie"))
	session, ok := cookies[network.COOKIE_SESSION_KEY]
	if !ok {
		return nil, errors.New("failed to login")
	}

	return &responses.LoginResponse{Token: session}, nil
}
