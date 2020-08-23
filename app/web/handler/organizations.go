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

// OrganizationsHandler ..
var (
	OrganizationsHandler organizationsHandler = organizationsHandler{}
)

type organizationsHandler struct{}

// Search ..
func (organizationsHandler) Search(c echo.Context) (err error) {

	type myResponse struct {
		Organization   model.Organizations `json:"organization"`
		TicketSubjects []string            `json:"ticket_subjects,omitempty"`
		UserNames      []string            `json:"user_names,omitempty"`
	}
	type myRequest struct {
		Term  string `json:"term" query:"term" validate:"required"`
		Value string `json:"value" query:"value"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = c.Validate(request); err != nil {
		return
	}
	// TODO: validate term, value type
	organizationReflect := reflect.ValueOf(&model.Organizations{}).Elem()
	typeOfT := organizationReflect.Type()
	organizationCondition := bson.M{}

	for i := 0; i < organizationReflect.NumField(); i++ {
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
		switch organizationReflect.Field(i).Kind().String() {
		case "int", "[]int":
			valueInt, errParse := strconv.Atoi(request.Value)
			if errParse != nil {
				return fmt.Errorf("Wrong format value for term: %s must be int", term)
			}
			organizationCondition[term] = valueInt
		case "string", "[]string":
			organizationCondition[term] = request.Value
		case "bool", "[]bool":
			valueBool, errParse := strconv.ParseBool(request.Value)
			if errParse != nil {
				return fmt.Errorf("Wrong format value for term: %s must be bool", term)
			}
			organizationCondition[term] = valueBool
		}
	}
	if len(organizationCondition) == 0 {
		return fmt.Errorf("Term not support search")
	}
	// TODO: Get organizations with condition
	organizations, err := organizationsRepo.New().All(organizationCondition)
	if err != nil {
		return fmt.Errorf("Connect to mongo fail: %s", err)
	}
	if len(organizations) == 0 {
		return fmt.Errorf("Not found organization with conditions")
	}
	response := []myResponse{}
	for _, organization := range organizations {
		temp := myResponse{
			Organization: organization,
		}
		// Get ticket subjects
		ticketCondition := bson.M{"organization_id": organization.ID}
		tickets, err := ticketsRepo.New().All(ticketCondition)
		if err != nil {
			return fmt.Errorf("Connect to mongo fail: %s", err)
		}
		for _, ticket := range tickets {
			temp.TicketSubjects = append(temp.TicketSubjects, ticket.Subject)
		}
		// Get user names
		userCondition := bson.M{"organization_id": organization.ID}
		users, err := usersRepo.New().All(userCondition)
		if err != nil {
			return fmt.Errorf("Connect to mongo fail: %s", err)
		}
		for _, user := range users {
			temp.UserNames = append(temp.UserNames, user.Name)
		}
		response = append(response, temp)
	}

	return c.JSON(success(response))
}
