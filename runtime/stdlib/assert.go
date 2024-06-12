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

package stdlib

import (
	"fmt"

	"github.com/onflow/cadence/runtime/errors"
	"github.com/onflow/cadence/runtime/interpreter"
	"github.com/onflow/cadence/runtime/sema"
)

// AssertFunction

const assertFunctionDocString = `
Terminates the program if the given condition is false, and reports a message which explains how the condition is false.
Use this function for internal sanity checks.

The message argument is optional.
`

var assertFunctionType = &sema.FunctionType{
	Purity: sema.FunctionPurityView,
	Parameters: []sema.Parameter{
		{
			Label:          sema.ArgumentLabelNotRequired,
			Identifier:     "condition",
			TypeAnnotation: sema.BoolTypeAnnotation,
		},
		{
			Identifier:     "message",
			TypeAnnotation: sema.StringTypeAnnotation,
		},
	},
	ReturnTypeAnnotation: sema.VoidTypeAnnotation,
	// `message` parameter is optional
	Arity: &sema.Arity{Min: 1, Max: 2},
}

var AssertFunction = NewStandardLibraryStaticFunction(
	"assert",
	assertFunctionType,
	assertFunctionDocString,
	func(invocation interpreter.Invocation) interpreter.Value {
		result, ok := invocation.Arguments[0].(interpreter.BoolValue)
		if !ok {
			panic(errors.NewUnreachableError())
		}

		if !result {
			var message string
			if len(invocation.Arguments) > 1 {
				messageValue, ok := invocation.Arguments[1].(*interpreter.StringValue)
				if !ok {
					panic(errors.NewUnreachableError())
				}
				message = messageValue.Str
			}
			panic(AssertionError{
				Message:       message,
				LocationRange: invocation.LocationRange,
			})
		}
		return interpreter.Void
	},
)

// AssertionError

type AssertionError struct {
	interpreter.LocationRange
	Message string
}

var _ errors.UserError = AssertionError{}

func (AssertionError) IsUserError() {}

func (e AssertionError) Error() string {
	const message = "assertion failed"
	if e.Message == "" {
		return message
	}
	return fmt.Sprintf("%s: %s", message, e.Message)
}
