// Copyright 2023 Edson Michaque
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/edsonmichaque/tibctl/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func cmdProfile(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use:   "profile",
		Short: "Manage profiles",
		Example: heredoc.Doc(`
			tibctl profile
			tibctl profile --output=json
			tibctl profile --output=yaml
			tibctl profile --output=json --query="[].id"
		`),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load()
			if err != nil {
				return newError(1, err)
			}

			cmd.Println(cfg)

			return nil
		},
	}

	return newCmd(
		cmd,
		withOutputFlag(formatTable),
		withQueryFlag(),
		withSubcommand(cmdProfileList(opts)),
		withSubcommand(cmdProfileCreate(opts)),
		withSubcommand(cmdProfileGet(opts)),
		withSubcommand(cmdProfileUpdate(opts)),
		withSubcommand(cmdProfileDelete(opts)),
		withOptions(opts),
	)
}

func cmdProfileList(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return newCmd(cmd, withOptions(opts))
}

func cmdProfileGet(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return newCmd(cmd, withOptions(opts))
}

func cmdProfileCreate(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use: "create",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return newCmd(cmd, withOptions(opts))
}

func cmdProfileUpdate(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return newCmd(cmd, withOptions(opts))
}

func cmdProfileDelete(opts *Options) *Cmd {
	cmd := &cobra.Command{
		Use: "delete",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return newCmd(cmd, withOptions(opts))
}
