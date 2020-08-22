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

// TicketsHandler ..
var (
	TicketsHandler ticketsHandler = ticketsHandler{}
)

type ticketsHandler struct{}

// Search ..
func (ticketsHandler) Search(c echo.Context) (err error) {

	type myResponse struct {
		Ticket            model.Tickets `json:"ticket"`
		OrganizationNames []string      `json:"organization_names,omitempty"`
		AssigneeNames     []string      `json:"assignee_names,omitempty"`
		SubmitterNames    []string      `json:"submitter_names,omitempty"`
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
	ticketReflect := reflect.ValueOf(&model.Tickets{}).Elem()
	typeOfT := ticketReflect.Type()
	ticketCondition := bson.M{}

	for i := 0; i < ticketReflect.NumField(); i++ {
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
		switch ticketReflect.Field(i).Kind().String() {
		case "int", "[]int":
			valueInt, errParse := strconv.Atoi(request.Value)
			if errParse != nil {
				return fmt.Errorf("Wrong format value for term: %s must be int", term)
			}
			ticketCondition[term] = valueInt
		case "string", "[]string":
			ticketCondition[term] = request.Value
		case "bool", "[]bool":
			valueBool, errParse := strconv.ParseBool(request.Value)
			if errParse != nil {
				return fmt.Errorf("Wrong format value for term: %s must be bool", term)
			}
			ticketCondition[term] = valueBool
		}
	}
	if len(ticketCondition) == 0 {
		return fmt.Errorf("Term not support search")
	}
	// TODO: Get tickets with condition
	tickets, err := ticketsRepo.New().All(ticketCondition)
	if err != nil {
		return fmt.Errorf("Connect to mongo fail: %s", err)
	}
	if len(tickets) == 0 {
		return fmt.Errorf("Not found ticket with conditions")
	}
	response := []myResponse{}
	for _, ticket := range tickets {
		temp := myResponse{
			Ticket: ticket,
		}
		// Get assignees
		assigneeCondition := bson.M{"_id": ticket.AssigneeID}
		assigners, err := usersRepo.New().All(assigneeCondition)
		if err != nil {
			return fmt.Errorf("Connect to mongo fail: %s", err)
		}
		for _, assignee := range assigners {
			temp.AssigneeNames = append(temp.AssigneeNames, assignee.Name)
		}
		// Get submiters
		submitterCondition := bson.M{"_id": ticket.SubmitterID}
		submitters, err := usersRepo.New().All(submitterCondition)
		if err != nil {
			return fmt.Errorf("Connect to mongo fail: %s", err)
		}
		for _, submitter := range submitters {
			temp.SubmitterNames = append(temp.SubmitterNames, submitter.Name)
		}
		// Get organizations
		organizationCondition := bson.M{
			"_id": ticket.OrganizationID,
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
