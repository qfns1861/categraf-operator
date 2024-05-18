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
	"sync"
	"text/template"

	"github.com/go-logr/logr"
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

func (r *CategrafMonitorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.V(2).Info("categrafMonitor event received")
	defer func() { logger.V(2).Info("categrafMonitor event handling completed") }()

	var categrafMonitor qfnsv1.CategrafMonitor
	if err := r.Get(ctx, req.NamespacedName, &categrafMonitor); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("categrafMonitor not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "unable to fetch categrafMonitor, requeue")
		return ctrl.Result{}, err
	}

	envNode := os.Getenv("NODE_NAME")
	fmt.Printf("%+v\n", envNode)
	crNode := strings.ReplaceAll(categrafMonitor.Spec.Node, "*", "")
	node_bool := strings.Contains(envNode, crNode)

	if node_bool {
		var wg sync.WaitGroup
		var mutex sync.Mutex // Mutex for file write operations

		// Process MonitorSuper
		for _, monitor := range categrafMonitor.Spec.MonitorSuper {
			wg.Add(1)
			go func(m qfnsv1.MonitorSuper) {
				defer wg.Done()
				processMonitorSuper(logger, m, &mutex)
			}(monitor)
		}

		// Process MonitorLite
		for _, monitor := range categrafMonitor.Spec.MonitorLite {
			wg.Add(1)
			go func(m qfnsv1.MonitorLite) {
				defer wg.Done()
				processMonitorLite(logger, m, &mutex)
			}(monitor)
		}

		wg.Wait() // Wait for all goroutines to complete
	}

	return ctrl.Result{}, nil
}

func processMonitorSuper(logger logr.Logger, monitor qfnsv1.MonitorSuper, mutex *sync.Mutex) {
	substrings := strings.Split(monitor.Name, ".")
	leftPart := substrings[0]
	folderPath := "categraf/input." + leftPart
	fileName := substrings[0] + ".toml"
	filePath := filepath.Join(folderPath, fileName)

	mutex.Lock()
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		logger.Error(err, "Failed to create directory", "path", folderPath)
		mutex.Unlock()
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		logger.Error(err, "Failed to create file", "path", filePath)
		mutex.Unlock()
		return
	}
	defer file.Close()
	mutex.Unlock()

	tmpl, err := template.New("config").Parse(`interval = {{ .Interval }}
[mappings]
{{ .Mappings }}
{{- range .Instances}}
[[instances]]
{{ .Data }}
{{- end }}
`)
	if err != nil {
		logger.Error(err, "Failed to parse template")
		return
	}

	err = tmpl.Execute(file, monitor)
	if err != nil {
		logger.Error(err, "Failed to execute template", "file", filePath)
		return
	}

	logger.Info("File written successfully", "file", filePath)
}
func processMonitorLite(logger logr.Logger, monitor qfnsv1.MonitorLite, mutex *sync.Mutex) {
	substrings := strings.Split(monitor.Name, ".")
	leftPart := substrings[0]
	folderPath := "categraf/input." + leftPart
	fileName := substrings[0] + ".toml"
	filePath := filepath.Join(folderPath, fileName)

	mutex.Lock()
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		logger.Error(err, "Failed to create directory", "path", folderPath)
		mutex.Unlock()
		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		logger.Error(err, "Failed to create file", "path", filePath)
		mutex.Unlock()
		return
	}
	defer file.Close()
	mutex.Unlock()

	_, err = io.WriteString(file, monitor.Data)
	if err != nil {
		logger.Error(err, "Failed to write data to file", "file", filePath)
		return
	}
	logger.Info("File written successfully", "file", filePath)
}

// SetupWithManager sets up the controller with the Manager.
func (r *CategrafMonitorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&qfnsv1.CategrafMonitor{}).
		Complete(r)
}
