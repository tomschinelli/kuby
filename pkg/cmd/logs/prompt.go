package logs

import (
	"github.com/AlecAivazis/survey/v2"
	"strings"
)

func promptDeployment() (namespace string, name string, error error) {
	deployments, err := validDeployments()
	if err != nil {
		return "", "", err
	}
	label := "Which logs do you want to inspect?"
	var opts []string

	for _, depl := range deployments {
		opts = append(opts, depl.Namespace+"	"+depl.Name)
	}

	var res string
	prompt := &survey.Select{
		Message: label,
		Options: opts,
	}
	err = survey.AskOne(prompt, &res)
	if err != nil {
		return "", "", err
	}
	resChunks := strings.Split(res, "	")
	return resChunks[1], resChunks[0], err
}
