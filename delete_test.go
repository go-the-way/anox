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

func init() {
	testInit()
}

func TestDeleteExec(t *testing.T) {
	truncateTestTable()
	o := New(new(userModel))
	{
		if err := insertTest(); err != nil {
			t.Fatalf("TestDeleteExec prepare insert failed: %v\n", err)
		}
		if _, err := o.Delete().IfWhere(false).IfWhere(true, getTestGes()...).Exec(nil); err != nil {
			t.Fatalf("TestDeleteExec failed: %v\n", err)
		}
		if selectUserModelCount() != 0 {
			t.Fatal("TestDeleteExec failed!")
		}
	}
	{
		if err := insertTest(); err != nil {
			t.Fatalf("TestDeleteExec prepare insert failed: %v\n", err)
		}
		if _, err := o.Delete().Where(getTestGes()...).Exec(nil); err != nil {
			t.Fatalf("TestDeleteExec failed: %v\n", err)
		}
		if selectUserModelCount() != 0 {
			t.Fatal("TestDeleteExec failed!")
		}
	}
	{
		if err := insertTest(); err != nil {
			t.Fatalf("TestDeleteExec prepare insert failed: %v\n", err)
		}
		if _, err := o.Delete().Exec(getTest()); err != nil {
			t.Fatalf("TestDeleteExec failed: %v\n", err)
		}
		if selectUserModelCount() != 0 {
			t.Fatal("TestDeleteExec failed!")
		}
	}
}

func TestDeleteExec2(t *testing.T) {
	truncateTestNullTable()
	o := New(new(userModelNull))
	{
		if err := insertNullTest(); err != nil {
			t.Fatalf("TestDeleteExec prepare insert failed: %v\n", err)
		}
		if _, err := o.Delete().IfWhere(true, getTestGes()...).Exec(nil); err != nil {
			t.Fatalf("TestDeleteExec failed: %v\n", err)
		}
		if selectUserModelNullCount() > 0 {
			t.Fatal("TestDeleteExec failed!")
		}
	}
	{
		if err := insertNullTest(); err != nil {
			t.Fatalf("TestDeleteExec prepare insert failed: %v\n", err)
		}
		if _, err := o.Delete().Where(getTestGes()...).Exec(nil); err != nil {
			t.Fatalf("TestDeleteExec failed: %v\n", err)
		}
		if selectUserModelNullCount() > 0 {
			t.Fatal("TestDeleteExec failed!")
		}
	}
	{
		if err := insertNullTest(); err != nil {
			t.Fatalf("TestDeleteExec prepare insert failed: %v\n", err)
		}
		if _, err := o.Delete().Exec(getNullTest()); err != nil {
			t.Fatalf("TestDeleteExec failed: %v\n", err)
		}
		if selectUserModelNullCount() > 0 {
			t.Fatal("TestDeleteExec failed!")
		}
	}
}

func TestDeleteExecTX(t *testing.T) {
	truncateTestTable()
	o := New(new(userModel))
	if err := insertTest(); err != nil {
		t.Fatalf("TestDeleteExecTX prepare insert failed: %v\n", err)
	}
	if err := o.Begin(); err != nil {
		t.Fatalf("TestDeleteExecTX failed: %v\n", err)
	}
	if _, err := o.Delete().IfWhere(false).IfWhere(true, getTestGes()...).Exec(nil); err != nil {
		t.Fatalf("TestDeleteExecTX failed: %v\n", err)
	}
	if err := o.Commit(); err != nil {
		t.Fatalf("TestDeleteExecTX failed: %v\n", err)
	}
	if selectUserModelCount() != 0 {
		t.Fatalf("TestDeleteExecTX failed!")
	}
}
