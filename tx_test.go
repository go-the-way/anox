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

import (
	"testing"
)

func TestTx(t *testing.T) {
	o := NewWithTx(new(userModel), false)
	oi := o.Insert()
	ou := o.Update()
	od := o.Delete()
	err := oi.SaveAll(false,
		&userModel{Name: "zhang san"},
		&userModel{Name: "zhang san"},
		&userModel{Name: "zhang sanzhang sanzhang sanzhang sanzhang sanzhang sanzhang san"},
		&userModel{Name: "zhang san"},
		&userModel{Name: "zhang san"},
		&userModel{Name: "zhang san"},
	)
	t.Log(err)
	oi.Save(&userModel{Name: "zhang san1111111"})
	ou.Modify(&userModel{ID: 12})
	od.Remove(&userModel{ID: 4, Name: "zhang san"})
	oi.orm.tx.Commit()
}
