/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright Flow Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package interpreter_utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/cadence/interpreter"
	. "github.com/onflow/cadence/test_utils/common_utils"
)

func NewTestInterpreter(tb testing.TB) *interpreter.Interpreter {
	storage := NewUnmeteredInMemoryStorage()
	return NewTestInterpreterWithStorage(tb, storage)
}

func NewTestInterpreterWithStorage(tb testing.TB, storage interpreter.Storage) *interpreter.Interpreter {
	return NewTestInterpreterWithStorageAndAtreeValidationConfig(tb, storage, true, true)
}

func NewTestInterpreterWithStorageAndAtreeValidationConfig(
	tb testing.TB,
	storage interpreter.Storage,
	atreeValueValidationEnabled bool,
	atreeStorageValidationEnabled bool,
) *interpreter.Interpreter {
	inter, err := interpreter.NewInterpreter(
		nil,
		TestLocation,
		&interpreter.Config{
			Storage:                       storage,
			AtreeValueValidationEnabled:   atreeValueValidationEnabled,
			AtreeStorageValidationEnabled: atreeStorageValidationEnabled,
		},
	)
	require.NoError(tb, err)

	return inter
}

func NewUnmeteredInMemoryStorage() interpreter.Storage {
	return interpreter.NewInMemoryStorage(nil)
}
