package webhook

import (
	"github.com/jenkins-x/jx-gitops/pkg/cmd/webhook/verify"
	"github.com/jenkins-x/jx-helpers/pkg/cobras"
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/spf13/cobra"
)

// NewCmdWebhook creates the new command
func NewCmdWebhook() *cobra.Command {
	command := &cobra.Command{
		Use:     "webhook",
		Short:   "Commands for working with WebHooks on your source repositories",
		Aliases: []string{"webhooks", "hook", "hooks"},
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	command.AddCommand(cobras.SplitCommand(verify.NewCmdWebHookVerify()))
	return command
}
