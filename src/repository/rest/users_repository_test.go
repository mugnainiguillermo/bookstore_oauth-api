package rest

import (
	goer "errors"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/domain/user"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/utils/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterNoResponder(func(request *http.Request) (*http.Response, error) {
		panic("no responder registered for this scenario, please register a responder that matches testing scenario")
	})
	os.Exit(m.Run())
}

func TestLoginUser(t *testing.T) {
	responderMethod := http.MethodPost
	responderUrl := fmt.Sprintf("%s%s", client.BaseURL, "/users/login")

	t.Run("ErrorDuringRequest", func(t *testing.T) {
		httpmock.RegisterResponder(
			responderMethod,
			responderUrl,
			httpmock.NewErrorResponder(goer.New("some error")),
		)
		repository := usersRepository{}

		user, restErr := repository.LoginUser("email@example.org", "password")

		assert.Nil(t, user)
		assert.Equal(t, errors.NewInternalServerError("error during client request"), restErr)
		httpmock.Reset()
	})
	t.Run("RestError", func(t *testing.T) {
		mockRestErr := errors.NewBadRequestError("generic rest error")
		httpmock.RegisterResponder(
			responderMethod,
			responderUrl,
			httpmock.NewJsonResponderOrPanic(http.StatusBadRequest, &mockRestErr),
		)
		repository := usersRepository{}

		user, restErr := repository.LoginUser("email@example.org", "password")

		assert.Nil(t, user)
		assert.Equal(t, mockRestErr, restErr)
		httpmock.Reset()
	})
	t.Run("Success", func(t *testing.T) {
		mockUser := user.User{
			Id:        123,
			FirstName: "Carla",
			LastName:  "Smith",
			Email:     "carla.smith@example.org",
		}
		httpmock.RegisterResponder(
			responderMethod,
			responderUrl,
			httpmock.NewJsonResponderOrPanic(http.StatusOK, mockUser),
		)
		repository := usersRepository{}

		user, restErr := repository.LoginUser("email@example.org", "password")

		assert.Equal(t, &mockUser, user)
		assert.Nil(t, restErr)
		httpmock.Reset()
	})
}
