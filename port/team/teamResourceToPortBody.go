package team

import (
	"context"

	"github.com/port-labs/terraform-provider-port-labs/internal/cli"
)

func TeamResourceToPortBody(ctx context.Context, state *TeamModel) (*cli.Team, error) {
	tp := &cli.Team{
		Name:        state.Name.ValueString(),
		Description: state.Description.ValueString(),
	}
	if state.Users != nil {
		tp.Users = make([]string, len(state.Users))
		for i, t := range state.Users {
			tp.Users[i] = t.ValueString()
		}
	}

	return tp, nil
}