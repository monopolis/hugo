// Copyright 2016 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/hugo/create"
	"github.com/spf13/hugo/hugofs"
	"github.com/spf13/viper"
)

func init() {

}

var diaryCmd = &cobra.Command{
	Use:   "diary",
	Short: "Create a new diary entry",
	Long: `Create a new diary entry in content/diary/<current_time>.md.`,
	RunE: DiaryEntry,
}

func DiaryEntry(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(); err != nil {
		return err
	}

    viper.Set("NewContentEditor", "vim")

	return create.NewDiaryContent(hugofs.Source(), "diary")
}

