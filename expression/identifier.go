//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package expression

import (
	"fmt"

	"github.com/couchbaselabs/query/value"
)

type Identifier struct {
	ExpressionBase
	identifier string
}

func NewIdentifier(identifier string) Expression {
	return &Identifier{
		identifier: identifier,
	}
}

func (this *Identifier) Evaluate(item value.Value, context Context) (value.Value, error) {
	rv, ok := item.Field(this.identifier)
	if !ok {
		return nil, fmt.Errorf("Unbound identifier %s for value %v.", this.identifier, item)
	}

	return rv, nil
}

func (this *Identifier) EquivalentTo(other Expression) bool {
	switch other := other.(type) {
	case *Identifier:
		return this.identifier == other.identifier
	default:
		return false
	}
}

func (this *Identifier) Dependencies() Expressions {
	return nil
}

func (this *Identifier) Fold() Expression {
	return this
}