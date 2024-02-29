package http

import (
	"assik3/pkg/type/phoneNumber"
	"assik3/services/contact/internal/domain/contact"
	"assik3/services/contact/internal/domain/contact/name"
	"assik3/services/contact/internal/domain/contact/patronymic"
	"assik3/services/contact/internal/domain/contact/surname"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (d *Delivery) CreateContactInGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupID, err := strconv.Atoi(ps.ByName("groupId"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	var tempContact struct {
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Patronymic  string `json:"patronymic"`
	}

	if err := json.NewDecoder(r.Body).Decode(&tempContact); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	phoneNumber, err := phoneNumber.New(tempContact.PhoneNumber)
	if err != nil {
		http.Error(w, "Invalid phone number: "+err.Error(), http.StatusBadRequest)
		return
	}

	name, err := name.New(tempContact.Name)
	if err != nil {
		http.Error(w, "Invalid name: "+err.Error(), http.StatusBadRequest)
		return
	}

	surname, err := surname.New(tempContact.Surname)
	if err != nil {
		http.Error(w, "Invalid surname: "+err.Error(), http.StatusBadRequest)
		return
	}

	patronymic, err := patronymic.New(tempContact.Patronymic)
	if err != nil {
		http.Error(w, "Invalid patronymic: "+err.Error(), http.StatusBadRequest)
		return
	}

	contact, err := contact.New(0, *phoneNumber, *name, *surname, *patronymic)
	if err != nil {
		http.Error(w, "Failed to create contact object: "+err.Error(), http.StatusInternalServerError)
		return
	}

	createdContact, err := d.ucGroup.CreateContantIntoGroup(groupID, contact)
	if err != nil {
		http.Error(w, "Failed to create contact in group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(createdContact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (d *Delivery) AddContactToGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupID, err := strconv.Atoi(ps.ByName("groupId"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	contactID, err := strconv.Atoi(ps.ByName("contactId"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	err = d.ucGroup.AddContactToGroup(groupID, contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact added to group successfully"))
}

func (d *Delivery) DeleteContactFromGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupID, err := strconv.Atoi(ps.ByName("groupId"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	contactID, err := strconv.Atoi(ps.ByName("contactId"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	err = d.ucGroup.DeleteContantFromGroup(groupID, contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact removed from group successfully"))
}
