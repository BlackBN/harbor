/*
   Copyright (c) 2016 VMware, Inc. All Rights Reserved.
   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at
       http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package models

import (
	"time"
	// klar_clair "github.com/optiopay/klar/clair"
)

// RepoRecord holds the record of an repository in DB, all the infors are from the registry notification event.
type RepoRecord struct {
	RepositoryID string    `orm:"column(repository_id);pk" json:"repository_id"`
	Name         string    `orm:"column(name)" json:"name"`
	OwnerName    string    `orm:"-"`
	OwnerID      int64     `orm:"column(owner_id)"  json:"owner_id"`
	ProjectName  string    `orm:"-"`
	ProjectID    int64     `orm:"column(project_id)"  json:"project_id"`
	Manager      string    `orm:"column(manager)" json:"manager"`
	Description  string    `orm:"column(description)" json:"description"`
	PullCount    int64     `orm:"column(pull_count)" json:"pull_count"`
	StarCount    int64     `orm:"column(star_count)" json:"star_count"`
	CreationTime time.Time `orm:"column(creation_time);auto_now_add" json:"creation_time"`
	UpdateTime   time.Time `orm:"column(update_time);auto_now" json:"update_time"`
}

//TableName is required by by beego orm to map RepoRecord to table repository
func (rp *RepoRecord) TableName() string {
	return "repository"
}

// ImageVulnerability holds the vulnerability of an image in DB.
// CPH, TODO, As mysql 5.6 not support JSON type, so use string to save vulnerabilities
type ImageVulnerability struct {
	RVID               string `orm:"column(rv_id);pk" json:"rv_id"`
	RepoName           string `orm:"column(repo_name)" json:"repo_name"`
	Tag                string `orm:"column(tag)" json:"tag"`
	VulnerabilityCount int    `orm:"column(v_count)" json:"v_count"`
	Vulnerabilities    string `orm:"column(vulnerabilities)" json:"vulnerabilities"`
}
