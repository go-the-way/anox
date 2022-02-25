// Copyright 2022 anox Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package anox

var (
	MySQLPagination = &mysqlPagination{}
)

type (
	// Pagination defines Pagination plugin
	Pagination interface {
		Page(originalSql string, offset, size int) (string, []interface{})
	}
	// ExecHooker defines exec hooker before and after exec SQL
	ExecHooker interface {
		BeforeExec(model Model, sqlStr *string, ps *[]interface{})
		AfterExec(model Model, sqlStr string, ps []interface{}, err error)
	}
	mysqlPagination struct {
	}
)

func (m *mysqlPagination) Page(originalSql string, offset, size int) (string, []interface{}) {
	return originalSql + " LIMIT ?, ?", []interface{}{offset, size}
}
