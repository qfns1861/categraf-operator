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
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	qfnsv1 "categraf-operator/api/v1"
)

// CategrafglobalReconciler reconciles a Categrafglobal object
type CategrafglobalReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=qfns.categraf-operator,resources=categrafglobals,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=qfns.categraf-operator,resources=categrafglobals/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=qfns.categraf-operator,resources=categrafglobals/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Categrafglobal object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.2/pkg/reconcile
func (r *CategrafglobalReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// 输出日志
	logger := log.FromContext(ctx)
	logger.V(2).Info("Categrafglobal event received")
	defer func() { logger.V(2).Info("Categrafglobal event handling completed") }()

	// TODO(user): your logic here

	var categrafglobal qfnsv1.Categrafglobal
	// 使用controller-runtime获取CR信息
	if err := r.Get(ctx, req.NamespacedName, &categrafglobal); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("Categrafglobal not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request
		logger.Error(err, "unable to fetch Categrafglobal, requeue")
		return ctrl.Result{}, err
	}

	jsonBytes, err := json.Marshal(&categrafglobal)
	if err != nil {
		logger.Error(err, "json Marshal error")
	}

	filePath := filepath.Join("/opt", "categrafglobal.json")
	err = ioutil.WriteFile(filePath, jsonBytes, 0644)
	if err != nil {
		logger.Error(err, "ioutil WriteFile error")
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CategrafglobalReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&qfnsv1.Categrafglobal{}).
		Complete(r)
}
