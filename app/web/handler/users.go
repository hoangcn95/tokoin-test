package handler

import (
	"app/model"
	organizationsRepo "app/repo/organizations"
	ticketsRepo "app/repo/tickets"
	usersRepo "app/repo/users"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// UsersHandler ..
var (
	UsersHandler usersHandler = usersHandler{}
)

type usersHandler struct{}

// Search ..
func (usersHandler) Search(c echo.Context) (err error) {

	type myResponse struct {
		User              model.Users `json:"user"`
		OrganizationNames []string    `json:"organization_names,omitempty"`
		TicketNames       []string    `json:"ticket_names,omitempty"`
	}
	type myRequest struct {
		Term  string `json:"term" query:"term" validate:"required"`
		Value string `json:"value" query:"value" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	// TODO: validate term, value type
	userReflect := reflect.ValueOf(&model.Users{}).Elem()
	typeOfT := userReflect.Type()
	userCondition := bson.M{}

	for i := 0; i < userReflect.NumField(); i++ {
		var term string
		keys := strings.Split(typeOfT.Field(i).Tag.Get("json"), ",")
		if len(keys) > 0 {
			term = keys[0]
		} else {
			term = typeOfT.Field(i).Tag.Get("json")
		}

		// validate term
		if request.Term != term {
			continue
		}

		// validate value, convert request value by type of term
		switch userReflect.Field(i).Kind().String() {
		case "int", "[]int":
			valueInt, errParse := strconv.Atoi(request.Value)
			if errParse != nil {
				return fmt.Errorf("Wrong format value for term: %s must be int", term)
			}
			userCondition[term] = valueInt
		case "string", "[]string":
			userCondition[term] = request.Value
		case "bool", "[]bool":
			valueBool, errParse := strconv.ParseBool(request.Value)
			if errParse != nil {
				return fmt.Errorf("Wrong format value for term: %s must be bool", term)
			}
			userCondition[term] = valueBool
		}
	}
	if len(userCondition) == 0 {
		return fmt.Errorf("Term must in: %s", "_id, url, external_id, name, alias, created_at, active, verified, shared, locale, timezone, last_login_at, email, phone, signature, organization_id, tags, suspended, role")
	}
	// TODO: Get users with condition
	users, err := usersRepo.New().All(userCondition)
	if err != nil {
		return fmt.Errorf("Connect to mongo fail: %s", err)
	}
	if len(users) == 0 {
		return fmt.Errorf("Not found user with conditions")
	}
	response := []myResponse{}
	for _, user := range users {
		temp := myResponse{
			User: user,
		}
		// Get Tickets by submitter_id or assignee_id
		ticketCondition := bson.M{
			"$or": []bson.M{
				bson.M{"submitter_id": user.ID},
				bson.M{"assignee_id": user.ID},
			},
		}
		tickets, err := ticketsRepo.New().All(ticketCondition)
		if err != nil {
			return fmt.Errorf("Connect to mongo fail: %s", err)
		}
		for _, ticket := range tickets {
			temp.TicketNames = append(temp.TicketNames, ticket.Subject)
		}
		// Get Organization by organization_id
		organizationCondition := bson.M{
			"_id": user.OrganizationID,
		}
		organizations, err := organizationsRepo.New().All(organizationCondition)
		if err != nil {
			return fmt.Errorf("Connect to mongo fail: %s", err)
		}
		for _, organization := range organizations {
			temp.OrganizationNames = append(temp.OrganizationNames, organization.Name)
		}
		response = append(response, temp)
	}

	return c.JSON(success(response))
}
