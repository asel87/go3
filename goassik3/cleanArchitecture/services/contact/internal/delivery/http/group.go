package http

import (
	"assik3/services/contact/internal/domain/group"
	"assik3/services/contact/internal/domain/group/name"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (d *Delivery) CreateGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var temp struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	name, err := name.New(temp.Name)
	if err != nil {
		http.Error(w, "Invalid group name: "+err.Error(), http.StatusBadRequest)
		return
	}

	group, err := group.New(0, *name)
	if err != nil {
		http.Error(w, "Failed to create group object: "+err.Error(), http.StatusInternalServerError)
		return
	}

	createdGroup, err := d.ucGroup.Create(group)
	if err != nil {
		http.Error(w, "Failed to create group: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(createdGroup)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (d *Delivery) GetGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	group, err := d.ucGroup.ReadByID(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(group)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (d *Delivery) UpdateGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	var temp struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	name, err := name.New(temp.Name)
	if err != nil {
		http.Error(w, "Invalid group name: "+err.Error(), http.StatusBadRequest)
		return
	}

	group, err := group.New(groupID, *name)
	if err != nil {
		http.Error(w, "Failed to create group object: "+err.Error(), http.StatusInternalServerError)
		return
	}

	updatedGroup, err := d.ucGroup.Update(groupID, group)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, _ := json.Marshal(updatedGroup)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (d *Delivery) DeleteGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	groupID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	err = d.ucGroup.Delete(groupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Group deleted successfully"))
}
