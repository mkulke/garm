// Copyright 2022 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cloudbase/garm/apiserver/params"
	gErrors "github.com/cloudbase/garm/errors"
	runnerParams "github.com/cloudbase/garm/params"

	"github.com/gorilla/mux"
)

func (a *APIController) CreateOrgHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var repoData runnerParams.CreateOrgParams
	if err := json.NewDecoder(r.Body).Decode(&repoData); err != nil {
		handleError(w, gErrors.ErrBadRequest)
		return
	}

	repo, err := a.r.CreateOrganization(ctx, repoData)
	if err != nil {
		log.Printf("error creating repository: %+v", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(repo); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) ListOrgsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	orgs, err := a.r.ListOrganizations(ctx)
	if err != nil {
		log.Printf("listing orgs: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) GetOrgByIDHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	orgID, ok := vars["orgID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	org, err := a.r.GetOrganizationByID(ctx, orgID)
	if err != nil {
		log.Printf("fetching org: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(org); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) DeleteOrgHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	orgID, ok := vars["orgID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	if err := a.r.DeleteOrganization(ctx, orgID); err != nil {
		log.Printf("removing org: %+v", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (a *APIController) UpdateOrgHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	orgID, ok := vars["orgID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	var updatePayload runnerParams.UpdateRepositoryParams
	if err := json.NewDecoder(r.Body).Decode(&updatePayload); err != nil {
		handleError(w, gErrors.ErrBadRequest)
		return
	}

	org, err := a.r.UpdateOrganization(ctx, orgID, updatePayload)
	if err != nil {
		log.Printf("error updating organization: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(org); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) CreateOrgPoolHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	orgID, ok := vars["orgID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	var poolData runnerParams.CreatePoolParams
	if err := json.NewDecoder(r.Body).Decode(&poolData); err != nil {
		log.Printf("failed to decode: %s", err)
		handleError(w, gErrors.ErrBadRequest)
		return
	}

	pool, err := a.r.CreateOrgPool(ctx, orgID, poolData)
	if err != nil {
		log.Printf("error creating organization pool: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pool); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) ListOrgPoolsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	orgID, ok := vars["orgID"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	pools, err := a.r.ListOrgPools(ctx, orgID)
	if err != nil {
		log.Printf("listing pools: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pools); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) GetOrgPoolHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	orgID, repoOk := vars["orgID"]
	poolID, poolOk := vars["poolID"]
	if !repoOk || !poolOk {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org or pool ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	pool, err := a.r.GetOrgPoolByID(ctx, orgID, poolID)
	if err != nil {
		log.Printf("listing pools: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pool); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}

func (a *APIController) DeleteOrgPoolHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	orgID, orgOk := vars["orgID"]
	poolID, poolOk := vars["poolID"]
	if !orgOk || !poolOk {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org or pool ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	if err := a.r.DeleteOrgPool(ctx, orgID, poolID); err != nil {
		log.Printf("removing pool: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func (a *APIController) UpdateOrgPoolHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	orgID, orgOk := vars["orgID"]
	poolID, poolOk := vars["poolID"]
	if !orgOk || !poolOk {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(params.APIErrorResponse{
			Error:   "Bad Request",
			Details: "No org or pool ID specified",
		}); err != nil {
			log.Printf("failed to encode response: %q", err)
		}
		return
	}

	var poolData runnerParams.UpdatePoolParams
	if err := json.NewDecoder(r.Body).Decode(&poolData); err != nil {
		log.Printf("failed to decode: %s", err)
		handleError(w, gErrors.ErrBadRequest)
		return
	}

	pool, err := a.r.UpdateOrgPool(ctx, orgID, poolID, poolData)
	if err != nil {
		log.Printf("error creating organization pool: %s", err)
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(pool); err != nil {
		log.Printf("failed to encode response: %q", err)
	}
}
