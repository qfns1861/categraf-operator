/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	qfnsv1 "categraf-operator/api/v1"
)

// CategrafMonitorReconciler reconciles a CategrafMonitor object
type CategrafMonitorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=qfns.categraf-operator,resources=categrafmonitors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=qfns.categraf-operator,resources=categrafmonitors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=qfns.categraf-operator,resources=categrafmonitors/finalizers,verbs=update

var nodeName string

func writeFile(filename string, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %s\n", filename, err.Error())
	} else {
		fmt.Printf("File %s written successfully\n", filename)
	}
}

func (r *CategrafMonitorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	logger := log.FromContext(ctx)
	logger.V(2).Info("categrafMonitor event received")
	defer func() { logger.V(2).Info("categrafMonitor event handling completed") }()

	var categrafMonitor qfnsv1.CategrafMonitor
	// 使用controller-runtime获取CR信息
	if err := r.Get(ctx, req.NamespacedName, &categrafMonitor); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("categrafMonitor not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request
		logger.Error(err, "unable to fetch categrafMonitor, requeue")
		return ctrl.Result{}, err
	}
	for _, monitor := range categrafMonitor.Spec.MonitorSuper {

		substrings := strings.Split(monitor.Name, ".")
		leftPart := substrings[0]
		folderPath := "input." + leftPart
		fileName := strings.Split(monitor.Name, ".")[0] + ".toml"
		filePath := filepath.Join(folderPath, fileName)

		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return ctrl.Result{}, err
		}

		tmpl, err := template.New("config").Parse(`interval = {{ .Interval }}
[mappings]
{{ .Mappings }}
{{- range .Instances}}
[[instances]]
{{ .Data }}
{{- end }}
`)
		if err != nil {
			return ctrl.Result{}, err
		}

		file, err := os.Create(filePath)
		if err != nil {
			return ctrl.Result{}, err
		}
		defer file.Close()

		err = tmpl.Execute(file, monitor)
		if err != nil {
			return ctrl.Result{}, err
		}
		fmt.Println("文件写入成功:", filePath)
	}

	for _, monitor := range categrafMonitor.Spec.MonitorLite {

		substrings := strings.Split(monitor.Name, ".")
		leftPart := substrings[0]
		folderPath := "input." + leftPart
		fileName := strings.Split(monitor.Name, ".")[0] + ".toml"
		filePath := filepath.Join(folderPath, fileName)

		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return ctrl.Result{}, err
		}

		file, err := os.Create(filePath)
		if err != nil {
			return ctrl.Result{}, err
		}
		defer file.Close()

		_, err = io.WriteString(file, monitor.Data)
		if err != nil {
			return ctrl.Result{}, err
		}

		fmt.Println("文件写入成功:", filePath)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CategrafMonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&qfnsv1.CategrafMonitor{}).
		Complete(r)
}
