package devfile

import (
	"path/filepath"
	"testing"

	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/devfile/library/v2/pkg/devfile/parser/data/v2/common"
	"github.com/google/go-cmp/cmp"
	"k8s.io/utils/pointer"
)

func TestParse(t *testing.T) {
	type args struct {
		devfileName string
		flatten     *bool
	}
	tests := []struct {
		name                string
		args                args
		wantAutoBuild       *bool
		wantDeployByDefault *bool
		wantErr             bool
	}{
		{
			name: "autoBuild=false, deployByDefault=true - flatten=nil => assumes flatten=true",
			args: args{
				devfileName: "devfile-with-both-autobuild-and-deploybydefault-set.yaml",
				flatten:     nil,
			},
			wantErr:             false,
			wantAutoBuild:       pointer.Bool(false),
			wantDeployByDefault: pointer.Bool(true),
		},
		{
			name: "autoBuild=false, deployByDefault=true - flatten=false",
			args: args{
				devfileName: "devfile-with-both-autobuild-and-deploybydefault-set.yaml",
				flatten:     pointer.Bool(false),
			},
			wantErr:             false,
			wantAutoBuild:       pointer.Bool(false),
			wantDeployByDefault: pointer.Bool(true),
		},
		{
			name: "autoBuild=false, deployByDefault=true - flatten=true",
			args: args{
				devfileName: "devfile-with-both-autobuild-and-deploybydefault-set.yaml",
				flatten:     pointer.Bool(true),
			},
			wantErr:             false,
			wantAutoBuild:       pointer.Bool(false),
			wantDeployByDefault: pointer.Bool(true),
		},

		{
			name: "autoBuild=true, deployByDefault=nil - flatten=false",
			args: args{
				devfileName: "devfile-with-autobuild-set-deploybydefault-unset.yaml",
				flatten:     pointer.Bool(false),
			},
			wantErr:             false,
			wantAutoBuild:       pointer.Bool(true),
			wantDeployByDefault: nil,
		},
		// The 2 tests below will fail because deployByDefault will be set to its default value (false) because of flatten=true.
		// But setting flatten to false causes issues with potentially incomplete Devfile validation when there is a parent-child relationship.
		{
			name: "autoBuild=true, deployByDefault=nil - flatten=nil => assumes flatten=true",
			args: args{
				devfileName: "devfile-with-autobuild-set-deploybydefault-unset.yaml",
				flatten:     nil,
			},
			wantErr:             false,
			wantAutoBuild:       pointer.Bool(true),
			wantDeployByDefault: nil,
		},
		{
			name: "autoBuild=true, deployByDefault=nil - flatten=true",
			args: args{
				devfileName: "devfile-with-autobuild-set-deploybydefault-unset.yaml",
				flatten:     pointer.Bool(true),
			},
			wantErr:             false,
			wantAutoBuild:       pointer.Bool(true),
			wantDeployByDefault: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := Parse(filepath.Join("testdata", tt.args.devfileName), tt.args.flatten)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			imageComps, err := d.Data.GetComponents(common.DevfileOptions{
				ComponentOptions: common.ComponentOptions{
					ComponentType: v1alpha2.ImageComponentType,
				}})
			if err != nil {
				t.Errorf("error while determining Image component: %v", err)
				return
			}
			if len(imageComps) != 1 {
				t.Errorf("unexpected number of image components: %d", len(imageComps))
				return
			}
			gotAutoBuild := imageComps[0].Image.AutoBuild
			if diff := cmp.Diff(tt.wantAutoBuild, gotAutoBuild); diff != "" {
				t.Errorf("autoBuild mismatch: %s\n", diff)
			}

			k8sComps, err := d.Data.GetComponents(common.DevfileOptions{
				ComponentOptions: common.ComponentOptions{
					ComponentType: v1alpha2.KubernetesComponentType,
				}})
			if err != nil {
				t.Errorf("error while determining Kubernetes component: %v", err)
				return
			}
			if len(k8sComps) != 1 {
				t.Errorf("unexpected number of Kubernetes components: %d", len(k8sComps))
				return
			}
			gotDeployByDefault := k8sComps[0].Kubernetes.DeployByDefault
			if diff := cmp.Diff(tt.wantDeployByDefault, gotDeployByDefault); diff != "" {
				t.Errorf("deployByDefault mismatch: %s\n", diff)
			}
		})
	}
}
