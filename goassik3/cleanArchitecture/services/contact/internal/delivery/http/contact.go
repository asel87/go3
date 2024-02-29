package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"assik3/pkg/type/phoneNumber"
	"assik3/services/contact/internal/domain/contact"
	"assik3/services/contact/internal/domain/contact/name"
	"assik3/services/contact/internal/domain/contact/patronymic"
	"assik3/services/contact/internal/domain/contact/surname"

	"github.com/julienschmidt/httprouter"
)

func (d *Delivery) CreateContact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var temp struct {
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Patronymic  string `json:"patronymic"`
	}

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	phoneNumber, err := phoneNumber.New(temp.PhoneNumber)
	if err != nil {
		http.Error(w, "Invalid phone number: "+err.Error(), http.StatusBadRequest)
		return
	}

	name, err := name.New(temp.Name)
	if err != nil {
		http.Error(w, "Invalid name: "+err.Error(), http.StatusBadRequest)
		return
	}

	surname, err := surname.New(temp.Surname)
	if err != nil {
		http.Error(w, "Invalid surname: "+err.Error(), http.StatusBadRequest)
		return
	}

	patronymic, err := patronymic.New(temp.Patronymic)
	if err != nil {
		http.Error(w, "Invalid patronymic: "+err.Error(), http.StatusBadRequest)
		return
	}

	contact, err := contact.New(0, *phoneNumber, *name, *surname, *patronymic)
	if err != nil {
		http.Error(w, "Failed to create contact object: "+err.Error(), http.StatusInternalServerError)
		return
	}

	createdContact, err := d.ucContact.Create(contact)
	if err != nil {
		http.Error(w, "Failed to create contact: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createdContact)
	if err != nil {
		http.Error(w, "Failed to serialize response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (d *Delivery) GetContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contactID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := d.ucContact.ReadByID(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(contact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (d *Delivery) UpdateContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contactID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var temp struct {
		PhoneNumber string `json:"phoneNumber"`
		Name        string `json:"name"`
		Surname     string `json:"surname"`
		Patronymic  string `json:"patronymic"`
	}

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	phoneNumber, err := phoneNumber.New(temp.PhoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name, err := name.New(temp.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	surname, err := surname.New(temp.Surname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	patronymic, err := patronymic.New(temp.Patronymic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact, err := contact.New(contactID, *phoneNumber, *name, *surname, *patronymic)
	if err != nil {
		http.Error(w, "Failed to create contact object", http.StatusInternalServerError)
		return
	}

	updatedContact, err := d.ucContact.Update(contactID, contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(updatedContact)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (d *Delivery) DeleteContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contactID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	err = d.ucContact.Delete(contactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact deleted successfully"))
}
