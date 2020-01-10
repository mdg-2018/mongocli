// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/10gen/mcli/internal/store"
	"github.com/10gen/mcli/internal/utils"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/spf13/cobra"
)

type atlasClustersListOpts struct {
	*globalOpts
	pageNum      int
	itemsPerPage int
	store        store.ClusterLister
}

func (opts *atlasClustersListOpts) init() error {
	if err := opts.loadConfig(); err != nil {
		return err
	}

	if opts.ProjectID() == "" {
		return errMissingProjectID
	}

	s, err := store.New(opts.Config)

	if err != nil {
		return err
	}

	opts.store = s
	return nil
}

func (opts *atlasClustersListOpts) Run() error {
	listOpts := opts.newListOptions()
	result, err := opts.store.ProjectClusters(opts.ProjectID(), listOpts)

	if err != nil {
		return err
	}

	return utils.PrettyJSON(result)
}

func (opts *atlasClustersListOpts) newListOptions() *atlas.ListOptions {
	return &atlas.ListOptions{
		PageNum:      opts.pageNum,
		ItemsPerPage: opts.itemsPerPage,
	}
}

// mcli atlas cluster(s) list --projectId projectId [--page N] [--limit N]
func AtlasClustersListBuilder() *cobra.Command {
	opts := &atlasClustersListOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Command to list Atlas clusters",
		Aliases: []string{"ls"},
		Args:    cobra.ExactArgs(0),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.projectID, flags.ProjectID, "", "Project ID")
	cmd.Flags().IntVar(&opts.pageNum, flags.Page, 0, "Page number")
	cmd.Flags().IntVar(&opts.itemsPerPage, flags.Limit, 0, "Items per page")

	cmd.Flags().StringVar(&opts.profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}