// Code generated by github.com/ervitis/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ervitis/gqlgen/plugin/federation/fedruntime"
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) ([]fedruntime.Entity, error) {
	list := []fedruntime.Entity{}
	for _, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			return nil, errors.New("__typename must be an existing string")
		}
		switch typeName {

		case "User":
			id0, err := ec.unmarshalNID2string(ctx, rep["id"])
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Field %s undefined in schema.", "id"))
			}

			entity, err := ec.resolvers.Entity().FindUserByID(ctx,
				id0)
			if err != nil {
				return nil, err
			}

			list = append(list, entity)

		default:
			return nil, errors.New("unknown type: " + typeName)
		}
	}
	return list, nil
}
