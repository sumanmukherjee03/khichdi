package k8s

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/sumanmukherjee03/gotils/cmd/utils"
)

const (
	ARTIFACT_BUILDER_POD = "artifact_builder_pod"
)

var (
	generateK8sTemplateShortDesc = "Provides k8s template generation capability"
	generateK8sTemplateLongDesc  = `Lets you generate k8s templates.
		The templates generated by this tool can help you create various resources in a kubernetes cluster.`
	generateK8sTemplateExample = `
	### Available commands for k8s
	# gotils k8s generate TEMPLATE_KIND
	gotils k8s generate artifact_builder_pod -d $(pwd)/artifact_builder_pod.json`
	validK8sTemplates = map[string]func() (string, error){
		ARTIFACT_BUILDER_POD: GenArtifactBuilderPodTemplate,
	}
	dest      string
	namespace string
	imageName string
	imageTag  string
	appPort   int
)

func NewK8sGenerator() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generate TEMPLATE_KIND",
		Short:   generateK8sTemplateShortDesc,
		Long:    generateK8sTemplateLongDesc,
		Example: generateK8sTemplateExample,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return utils.RaiseCmdErr(cmd, "Kind of k8s template not provided")
			}
			if len(args) > 1 {
				return utils.RaiseCmdErr(cmd, "Too many args")
			}
			if _, found := validK8sTemplates[args[0]]; !found {
				return utils.RaiseCmdErr(cmd, "Wrong type of k8s template provided")
			}
			return nil
		},
		TraverseChildren: true,
		Run: func(cmd *cobra.Command, args []string) {
			genTemplate(args[0])
		},
	}
	cmd.Flags().StringVarP(&dest, "dest", "d", "", "Full path to the output file")
	cmd.MarkFlagRequired("dest")
	cmd.Flags().StringVarP(&namespace, "namespace", "", "", "Namespace on which to perform the operations")
	cmd.MarkFlagRequired("namespace")
	cmd.Flags().StringVarP(&imageName, "image", "", "", "Docker image name")
	cmd.MarkFlagRequired("image")
	cmd.Flags().StringVarP(&imageTag, "tag", "", "", "Docker image tag")
	cmd.MarkFlagRequired("tag")
	cmd.Flags().IntVarP(&appPort, "port", "", 0, "Application port to be exposed")
	return cmd
}

////////////////////////// Unexported funcs //////////////////////////

func genTemplate(key string) {
	data, err := validK8sTemplates[key]()
	if err != nil {
		utils.CheckErr(err.Error())
	}

	f, err := os.Create(dest)
	if err != nil {
		utils.CheckErr(err.Error())
	}
	defer f.Close()

	_, err = f.Write([]byte(data))
	if err != nil {
		utils.CheckErr(err.Error())
	}
}
