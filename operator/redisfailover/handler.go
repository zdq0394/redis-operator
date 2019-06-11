package redisfailover

import (
	"context"
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	redisfailoverv1 "github.com/spotahome/redis-operator/api/redisfailover/v1"
	"github.com/spotahome/redis-operator/log"
	"github.com/spotahome/redis-operator/metrics"
	rfservice "github.com/spotahome/redis-operator/operator/redisfailover/service"
	"github.com/spotahome/redis-operator/operator/redisfailover/util"
	"github.com/spotahome/redis-operator/service/k8s"
)

const (
	rfLabelManagedByKey = "app.kubernetes.io/managed-by"
	rfLabelNameKey      = "redisfailovers.databases.spotahome.com/name"
)

var (
	defaultLabels = map[string]string{
		rfLabelManagedByKey: operatorName,
	}
)

// RedisFailoverHandler is the Redis Failover handler. This handler will create the required
// resources that a RF needs.
type RedisFailoverHandler struct {
	config     Config
	k8sservice k8s.Services
	rfService  rfservice.RedisFailoverClient
	rfChecker  rfservice.RedisFailoverCheck
	rfHealer   rfservice.RedisFailoverHeal
	mClient    metrics.Instrumenter
	logger     log.Logger
}

// NewRedisFailoverHandler returns a new RF handler
func NewRedisFailoverHandler(config Config, rfService rfservice.RedisFailoverClient, rfChecker rfservice.RedisFailoverCheck, rfHealer rfservice.RedisFailoverHeal, k8sservice k8s.Services, mClient metrics.Instrumenter, logger log.Logger) *RedisFailoverHandler {
	return &RedisFailoverHandler{
		config:     config,
		rfService:  rfService,
		rfChecker:  rfChecker,
		rfHealer:   rfHealer,
		mClient:    mClient,
		k8sservice: k8sservice,
		logger:     logger,
	}
}

// Add will ensure the redis failover is in the expected state.
func (r *RedisFailoverHandler) Add(_ context.Context, obj runtime.Object) error {
	rf, ok := obj.(*redisfailoverv1.RedisFailover)
	if !ok {
		return fmt.Errorf("can't handle the received object: not a redisfailover")
	}

	if err := rf.Validate(); err != nil {
		r.mClient.SetClusterError(rf.Namespace, rf.Name)
		return err
	}

	// Create owner refs so the objects manager by this handler have ownership to the
	// received RF.
	oRefs := r.createOwnerReferences(rf)

	// Create the labels every object derived from this need to have.
	labels := r.getLabels(rf)

	if err := r.Ensure(rf, labels, oRefs); err != nil {
		r.mClient.SetClusterError(rf.Namespace, rf.Name)
		return err
	}

	if err := r.CheckAndHeal(rf); err != nil {
		r.mClient.SetClusterError(rf.Namespace, rf.Name)
		return err
	}

	if err := r.setRedisFailoverStatus(rf); err != nil {
		r.logger.Warningf("Failed to update redisfailover(%s:%s) status\n", rf.Namespace, rf.Name)
		return err
	}

	r.mClient.SetClusterOK(rf.Namespace, rf.Name)
	return nil
}

// Delete handles the deletion of a RF.
func (r *RedisFailoverHandler) Delete(_ context.Context, name string) error {
	n := strings.Split(name, "/")
	if len(n) >= 2 {
		r.mClient.DeleteCluster(n[0], n[1])
	}
	// No need to do anything, it will be handled by the owner reference done
	// on the creation.
	r.logger.Debugf("ignoring, kubernetes GCs all using the objects OwnerReference metadata")
	return nil
}

// getLabels merges the labels (dynamic and operator static ones).
func (r *RedisFailoverHandler) getLabels(rf *redisfailoverv1.RedisFailover) map[string]string {
	dynLabels := map[string]string{
		rfLabelNameKey: rf.Name,
	}
	return util.MergeLabels(defaultLabels, dynLabels, rf.Labels)
}

func (w *RedisFailoverHandler) createOwnerReferences(rf *redisfailoverv1.RedisFailover) []metav1.OwnerReference {
	rfvk := redisfailoverv1.VersionKind(redisfailoverv1.RFKind)
	return []metav1.OwnerReference{
		*metav1.NewControllerRef(rf, rfvk),
	}
}

func (r *RedisFailoverHandler) setRedisFailoverStatus(rf *redisfailoverv1.RedisFailover) error {
	redisNodes, err := r.rfChecker.GetRedisesIPandHostIPs(rf)
	if err != nil {
		return fmt.Errorf("Get redisfailover status for %s:%s failed.", rf.Namespace, rf.Name)
	}
	rf.Status.RedisNodes = redisNodes
	if _, err := r.k8sservice.UpdateRedisFailovers(rf.Namespace, rf); err != nil {
		return err
	}
	return nil
}
