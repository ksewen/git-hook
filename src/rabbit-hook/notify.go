package main

import (
	"fmt"
	"gopkg.in/go-playground/webhooks.v5/gitlab"
	"strings"
)

const (
	pushEventTitle   = "%s %s branch [%s](%s) of [%s](%s)"
	pushEventCompare = "([Compare changes](%s))"
	pushEventMessage = ">[%s](%s): %s"

	issueEventTitle = "%s (%s) %s issue#%d:[%s](%s) in [%s](%s)"
	eventAssign     = ">assignee: %s"

	mergeRequestTitle = "%s (%s) %s	[!%d %s](%s) in [%s](%s)"

	emptyFlag = "0000000000000000000000000000000000000000"
)

func buildPushNotify(model gitlab.PushEventPayload) string {
	var pushAction = "pushed to"
	var notifyMessage = ""
	if model.Before == emptyFlag {
		pushAction = "pushed to new"
	} else if model.After == emptyFlag {
		pushAction = "delete"
	}

	ref := model.Ref
	refSplit := strings.Split(ref, "/")
	branch := refSplit[len(refSplit)-1]
	branchUrl := model.Project.Homepage + "/tree/" + branch
	notifyMessage = fmt.Sprintf(pushEventTitle, model.UserName, pushAction, branch, branchUrl, model.Project.PathWithNamespace, model.Project.Homepage)
	if pushAction == "pushed to" && len(model.Commits) > 0 {
		compareUrl := model.Project.Homepage + "/commit/" + model.Commits[0].ID
		notifyMessage = notifyMessage + fmt.Sprintf(pushEventCompare, compareUrl)

		for i := 0; i < len(model.Commits); i++ {
			shortCommitId := string([]rune(model.Commits[i].ID)[0:8])
			commitUrl := model.Project.Homepage + "/commit/" + model.Commits[i].ID
			gitMessageArr := strings.Split(model.Commits[i].Message, "\n")
			gitMessage := gitMessageArr[0]
			if len(gitMessageArr) > 1 {
				for i := 1; i < len(gitMessageArr); i++ {
					if gitMessageArr[i] == "" {
						continue
					}
					gitMessage = gitMessage + fmt.Sprintf("\n>\\- %s", gitMessageArr[i])
				}
			}
			notifyMessage = notifyMessage + "\n" + fmt.Sprintf(pushEventMessage, shortCommitId, commitUrl, gitMessage)
		}
	}
	return notifyMessage
}

func buildIssueNotify(model gitlab.IssueEventPayload) string {
	action := model.ObjectAttributes.Action
	issueUrl := fmt.Sprintf(model.Project.Homepage+"/issues/%d", model.ObjectAttributes.IID)
	issueMessage := fmt.Sprintf(issueEventTitle, model.User.Name, model.User.UserName, action, model.ObjectAttributes.IID, model.ObjectAttributes.Title, issueUrl, model.Project.Name, model.Project.Homepage)
	if model.Assignees != nil && len(model.Assignees) != 0 && model.Assignees[0].Username != "" {
		issueMessage = issueMessage + "\n" + fmt.Sprintf(eventAssign, model.Assignees[0].Username)
	}
	return issueMessage
}

func buildMergeRequestNotify(model gitlab.MergeRequestEventPayload) string {
	mergeRequestUrl := fmt.Sprintf(model.Project.Homepage+"/merge_requests/%d", model.ObjectAttributes.IID)
	mergeRequestMessage := fmt.Sprintf(mergeRequestTitle, model.User.Name, model.User.UserName, model.ObjectAttributes.Action, model.ObjectAttributes.IID, model.ObjectAttributes.Title, mergeRequestUrl, model.Project.PathWithNamespace, model.Project.Homepage)
	if model.Assignee.Username != "" {
		mergeRequestMessage = mergeRequestMessage + "\n" + fmt.Sprintf(eventAssign, model.Assignee.Username)
	}
	return mergeRequestMessage
}
