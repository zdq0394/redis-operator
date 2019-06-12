package redisfailover

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	redisfailoverv1 "github.com/spotahome/redis-operator/api/redisfailover/v1"
)

func (w *RedisFailoverHandler) Ensure(rf *redisfailoverv1.RedisFailover, labels map[string]string, or []metav1.OwnerReference) error {
	if rf.Spec.Redis.Exporter.Enabled {
		if err := w.rfService.EnsureRedisService(rf, labels, or); err != nil {
			return err
		}
	} else {
		if err := w.rfService.EnsureNotPresentRedisService(rf); err != nil {
			return err
		}
	}
	if err := w.rfService.EnsureSentinelService(rf, labels, or); err != nil {
		return err
	}
	if err := w.rfService.EnsureSentinelConfigMap(rf, labels, or); err != nil {
		return err
	}
	if err := w.rfService.EnsureRedisShutdownConfigMap(rf, labels, or); err != nil {
		return err
	}
	if err := w.rfService.EnsureRedisConfigMap(rf, labels, or); err != nil {
		return err
	}
	if err := w.rfService.EnsureRedisStatefulset(rf, labels, or); err != nil {
		return err
	}
	if err := w.rfService.EnsureSentinelDeployment(rf, labels, or); err != nil {
		return err
	}

	return nil
}

func (w *RedisFailoverHandler) EnsureRedisPerceptronDeployment(rf *redisfailoverv1.RedisFailover, hostIPs []string, labels map[string]string, ownerRefs []metav1.OwnerReference) error {
	if err := w.rfService.EnsureRedisPerceptronDeployment(rf, hostIPs, labels, ownerRefs); err != nil {
		return err
	}
	return nil
}
