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
	"testing"
)

func TestRegisterNil(t *testing.T) {
	defer func() {
		if re := recover(); re != errCannotRegisterNilModel {
			t.Fatal("when register a nil model, expect get errCannotRegisterNilModel")
		}
	}()
	Register(nil)
}

func TestRegisterDuplicate(t *testing.T) {
	Configuration.Migrate = false
	model := new(userModel)
	defer func() {
		if re := recover(); re != errDuplicateRegisterModel {
			t.Fatal("when register a duplicate model, expect get errDuplicateRegisterModel")
		}
	}()
	Register(model)
	Register(model)
}

func TestRegister(t *testing.T) {
	Configuration.Migrate = true
	Configuration.TableNameStrategy = CamelCase
}
