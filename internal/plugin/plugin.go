package plugin

import (
	"fmt"
	"strings"

	"github.com/cludden/protoc-gen-go-temporal/pkg/strcase"
	j "github.com/dave/jennifer/jen"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	deprecatedComment = "Deprecated: Do not use."
)

var (
	multiLineArgs = j.Options{
		Close:     ")",
		Multi:     true,
		Open:      "(",
		Separator: ",",
	}

	multiLineValues = j.Options{
		Close:     "}",
		Multi:     true,
		Open:      "{",
		Separator: ",",
	}
)

type Config struct {
	CliCategories              bool
	CliEnabled                 bool
	CliV3Enabled               bool
	DisableWorkflowInputRename bool
	DocsOut                    string
	DocsTemplate               string
	EnableCodec                bool
	EnableDebugLogging         bool
	EnablePatchSupport         bool
	EnableXNS                  bool
	NexusEnabled               bool
	NexusExcludeOperationTags  string
	NexusIncludeOperationTags  string
	NexusExcludeServiceTags    string
	NexusIncludeServiceTags    string
	Patches                    string
	IgnoreAcronyms             string
	WorkflowUpdateEnabled      bool
}

// Plugin provides a protoc plugin for generating temporal workers and clients in go
type Plugin struct {
	*protogen.Plugin

	Commit               string
	Version              string
	cfg                  *Config
	flags                *pflag.FlagSet
	caser                *strcase.Caser
	excludeOperationTags map[string]struct{}
	excludeServiceTags   map[string]struct{}
	includeOperationTags map[string]struct{}
	includeServiceTags   map[string]struct{}
}

func Run(commit, version string) {
	p := New(commit, version)
	protogen.Options{ParamFunc: p.Param}.Run(p.Run)
}

func New(commit, version string) *Plugin {
	var cfg Config

	flags := pflag.NewFlagSet("plugin", pflag.ExitOnError)
	flags.BoolVar(&cfg.CliEnabled, "cli-enabled", false, "enable cli generation")
	flags.BoolVar(&cfg.CliV3Enabled, "cli-v3", false, "enable cli generation using urfave/cli/v3")
	flags.BoolVar(&cfg.CliCategories, "cli-categories", true, "enable cli categories")
	flags.BoolVar(&cfg.DisableWorkflowInputRename, "disable-workflow-input-rename", false, "disable renaming of \"<Workflow>WorkflowInput\"")
	flags.StringVar(&cfg.DocsOut, "docs-out", "", "docs output path")
	flags.StringVar(&cfg.DocsTemplate, "docs-template", "basic", "built-in template name or path to custom template file")
	flags.BoolVar(&cfg.EnableCodec, "enable-codec", false, "enables experimental codec support")
	flags.BoolVar(&cfg.EnableDebugLogging, "enable-debug-logging", false, "enables debug logging")
	flags.BoolVar(&cfg.EnablePatchSupport, "enable-patch-support", false, "enables support for alta/protopatch renaming")
	flags.BoolVar(&cfg.EnableXNS, "enable-xns", false, "enable experimental cross-namespace workflow client")
	flags.StringVar(&cfg.IgnoreAcronyms, "ignore-acronyms", "", "semicolon-delimited string of acronyms to ignore when converting generated output to camel case")
	flags.BoolVar(&cfg.NexusEnabled, "nexus", false, "enable nexus handler generation")
	flags.StringVar(&cfg.NexusExcludeOperationTags, "nexus-exclude-operation-tags", "", "semicolon-delimited list of operation tags to exclude from nexus generation")
	flags.StringVar(&cfg.NexusIncludeOperationTags, "nexus-include-operation-tags", "", "semicolon-delimited list of operation tags to include in nexus generation")
	flags.StringVar(&cfg.NexusExcludeServiceTags, "nexus-exclude-service-tags", "", "semicolon-delimited list of service tags to exclude from nexus generation")
	flags.StringVar(&cfg.NexusIncludeServiceTags, "nexus-include-service-tags", "", "semicolon-delimited list of service tags to include in nexus generation")
	// patches is a semicolon-delimited string of
	flags.StringVar(&cfg.Patches, "patches", "", "semicolon-delimited string of <PATCH_VERSION>[_<MODE>] (e.g. --patches=64_MARKER;65_REMOVED)")
	flags.BoolVar(&cfg.WorkflowUpdateEnabled, "workflow-update-enabled", true, "enable experimental workflow update (DEPRECATED)")

	return &Plugin{
		Commit:  commit,
		Version: version,
		cfg:     &cfg,
		flags:   flags,
	}
}

// Param provides a protogen ParamFunc handler
func (p *Plugin) Param(key, value string) error {
	return p.flags.Set(key, value)
}

// Run defines the plugin entrypoint
func (p *Plugin) Run(plugin *protogen.Plugin) (err error) {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_SUPPORTS_EDITIONS | pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2023

	p.Plugin = plugin
	services, err := parse(p)
	if err != nil {
		return err
	}

	var ignoreAcronyms []string
	for _, a := range strings.Split(p.cfg.IgnoreAcronyms, ";") {
		if aa := strings.TrimSpace(a); aa != "" {
			ignoreAcronyms = append(ignoreAcronyms, aa)
		}
	}
	p.caser = strcase.NewCaser(strcase.WithAcronyms(ignoreAcronyms...))
	p.excludeOperationTags = p.nexusGetTags(p.cfg.NexusExcludeOperationTags)
	p.excludeServiceTags = p.nexusGetTags(p.cfg.NexusExcludeServiceTags)
	p.includeOperationTags = p.nexusGetTags(p.cfg.NexusIncludeOperationTags)
	p.includeServiceTags = p.nexusGetTags(p.cfg.NexusIncludeServiceTags)

	return services.render()
}

func commentf[T any, PT interface {
	*T
	Comment(string) *j.Statement
	Commentf(string, ...any) *j.Statement
}](c PT, methods []*protogen.Method, defaultMsg string, a ...any) {
	var deprecated bool
	for _, method := range methods {
		deprecated = isDeprecated(method)
		if deprecated {
			break
		}
	}
	c.Commentf(defaultMsg, a...)
	if deprecated {
		c.Comment(" ")
		c.Comment(deprecatedComment)
	}
}

func commentWithDefaultf[T any, PT interface {
	*T
	Comment(string) *j.Statement
	Commentf(string, ...any) *j.Statement
}](c PT, methods []*protogen.Method, defaultMsg string, a ...any) {
	var deprecated bool
	for _, method := range methods {
		deprecated = isDeprecated(method)
		if deprecated {
			break
		}
	}
	if len(methods) == 1 && methods[0].Comments.Leading.String() != "" {
		comment := strings.TrimSuffix(methods[0].Comments.Leading.String(), "\n")
		c.Comment(comment)
		if deprecated && !strings.Contains(comment, "Deprecated:") {
			c.Comment(" ")
			c.Comment(deprecatedComment)
		}
	} else {
		c.Commentf(defaultMsg, a...)
		if deprecated {
			c.Comment(" ")
			c.Comment(deprecatedComment)
		}
	}
}

func genCodeGenerationHeader(p *Plugin, f *j.File, target *protogen.File) {
	f.PackageComment("Code generated by protoc-gen-go_temporal. DO NOT EDIT.")
	f.PackageComment("versions: ")
	f.PackageComment(fmt.Sprintf("    protoc-gen-go_temporal %s (%s)", p.Version, p.Commit))
	compilerVersion := p.Plugin.Request.CompilerVersion
	if compilerVersion != nil {
		f.PackageComment(fmt.Sprintf("    protoc %s", compilerVersion.String()))
	} else {
		f.PackageComment("    protoc (unknown)")
	}

	f.PackageComment(fmt.Sprintf("source: %s", target.Desc.Path()))
}

func isDeprecated(method *protogen.Method) bool {
	return method.Desc.Options().(*descriptorpb.MethodOptions).GetDeprecated()
}

// isEmpty checks if the message is a google.protobuf.Empty message
func isEmpty(m *protogen.Message) bool {
	return m.Desc.FullName() == "google.protobuf.Empty"
}

func methodSet(methods ...*protogen.Method) []*protogen.Method {
	return methods
}
