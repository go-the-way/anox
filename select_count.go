// Copyright 2022 anorm Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package anorm

import (
	"github.com/go-the-way/sg"
)

type selectCountOperation[E Entity] struct {
	orm    *Orm[E]
	wheres []sg.Ge
	joins  []sg.Ge
}

func newSelectCountOperation[E Entity](o *Orm[E]) *selectCountOperation[E] {
	return &selectCountOperation[E]{orm: o, wheres: make([]sg.Ge, 0)}
}

// IfWhere if cond is true, append wheres
func (o *selectCountOperation[E]) IfWhere(cond bool, wheres ...sg.Ge) *selectCountOperation[E] {
	if cond {
		return o.Where(wheres...)
	}
	return o
}

// Where append wheres
func (o *selectCountOperation[E]) Where(wheres ...sg.Ge) *selectCountOperation[E] {
	o.wheres = append(o.wheres, wheres...)
	return o
}

// Where append wheres
func (o *selectCountOperation[E]) join(joins ...sg.Ge) *selectCountOperation[E] {
	o.joins = append(o.joins, joins...)
	return o
}

func (o *selectCountOperation[E]) getTableName() sg.Ge {
	return sg.T(entityTableMap[getEntityPkgName(o.orm.entity)])
}

func (o *selectCountOperation[E]) appendWhereGes(entity E) {
	o.wheres = append(o.wheres, o.orm.getWhereGes(entity)...)
}

// Exec select count
//
// Params:
//
// - entity: the orm wrapper entity
//
// Returns:
//
// - count: rows count
//
// - err: exec error
//
func (o *selectCountOperation[E]) Exec(entity E) (count int64, err error) {
	o.appendWhereGes(entity)
	sqlStr, ps := sg.SelectBuilder().
		Select(sg.Alias(sg.C("count(0)"), "c")).
		From(sg.Alias(o.getTableName(), "t")).
		Where(sg.AndGroup(o.wheres...)).
		Join(o.joins...).
		Build()
	queryLog("OpsForSelectCount.Exec", sqlStr, ps)
	row := o.orm.db.QueryRow(sqlStr, ps...)
	queryErrorLog("OpsForSelectCount.Exec", sqlStr, ps, row.Err())
	if err = row.Err(); err != nil {
		return
	}
	err = row.Scan(&count)
	return
}
