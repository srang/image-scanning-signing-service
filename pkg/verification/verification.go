package verification

import (
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/redhat-cop/image-scanning-signing-service/pkg/apis/cop/v1alpha2"
	"github.com/redhat-cop/image-scanning-signing-service/pkg/config"
	"github.com/redhat-cop/image-scanning-signing-service/pkg/util"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"

	corev1 "k8s.io/api/core/v1"
)

func UpdateOnVerificationPodLaunch(message string, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionTrue, v1alpha2.ImageExecutionConditionInitialization)

	imageVerificationRequestCopy.Status.StartTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseRunning)
}

func UpdateOnImageVerificationInitializationFailure(message string, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionFalse, v1alpha2.ImageExecutionConditionInitialization)

	imageVerificationRequestCopy.Status.StartTime = condition.LastTransitionTime
	imageVerificationRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseFailed)
}

func updateImageVerificationRequest(imageVerificationRequest *v1alpha2.ImageVerificationRequest, condition v1alpha2.ImageExecutionCondition, phase v1alpha2.ImageExecutionPhase) error {

	imageVerificationRequest.Status.Conditions = append(imageVerificationRequest.Status.Conditions, condition)
	imageVerificationRequest.Status.Phase = phase

	err := sdk.Update(imageVerificationRequest)

	return err
}

func UpdateOnImageVerificationCompletionError(message string, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionFalse, v1alpha2.ImageExecutionConditionFinished)

	imageVerificationRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseFailed)
}

func UpdateOnImageVerificationCompletionSuccess(message string, signatureStatus v1alpha2.SignatureStatus, imageVerificationRequest v1alpha2.ImageVerificationRequest) error {
	imageVerificationRequestCopy := imageVerificationRequest.DeepCopy()

	condition := util.NewImageExecutionCondition(message, corev1.ConditionTrue, v1alpha2.ImageExecutionConditionFinished)

	imageVerificationRequestCopy.Status.SignatureStatus = signatureStatus
	imageVerificationRequestCopy.Status.EndTime = condition.LastTransitionTime

	return updateImageVerificationRequest(imageVerificationRequestCopy, condition, v1alpha2.PhaseCompleted)
}

func LaunchVerificationPod(config config.Config, image string, ownerID string, ownerReference string, gpgSecretName string, gpgSignedBy string) (string, error) {

	pod, err := createVerificationPod(config.SignScanImage, config.TargetProject, image, ownerID, ownerReference, config.TargetServiceAccount, gpgSecretName, gpgSignedBy)

	if err != nil {
		logrus.Errorf("Error Generating Pod: %v'", err)
		return "", err
	}

	err = sdk.Create(pod)

	if err != nil {
		logrus.Errorf("Error Creating Pod: %v'", err)
		return "", err
	}

	var key string
	if key, err = cache.MetaNamespaceKeyFunc(pod); err != nil {
		runtime.HandleError(err)
		return "", err
	}

	return key, nil
}

func createVerificationPod(signScanImage string, targetProject string, image string, imageDigest string, ownerID string, ownerReference string, serviceAccount string, gpgSecret string, signedBy string) (*corev1.Pod, error) {

	priv := true

	pod := &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        ownerID,
			Namespace:   targetProject,
			Labels:      map[string]string{"type": "image-verification"},
			Annotations: map[string]string{"cop.redhat.com/owner": ownerReference, "cop.redhat.com/type": "image-verification"},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{{
				Name:            "image-verifier",
				Image:           signScanImage,
				ImagePullPolicy: corev1.PullAlways,
				Command:         []string{"/bin/bash", "-c", "mkdir -p ~/.gnupg && cp /root/gpg/* ~/.gnupg && /usr/local/bin/verify-image"},
				Env: []corev1.EnvVar{
					{
						Name:      "NAMESPACE",
						ValueFrom: &corev1.EnvVarSource{FieldRef: &corev1.ObjectFieldSelector{FieldPath: "metadata.namespace"}},
					},
					{
						Name:  "IMAGE",
						Value: image,
					},
					{
						Name:  "DIGEST",
						Value: imageDigest,
					},
					{
						Name:  "SIGNEDBY",
						Value: signedBy,
					},
				},
				SecurityContext: &corev1.SecurityContext{
					Privileged: &priv,
				},
				VolumeMounts: []corev1.VolumeMount{
					{
						Name:      "docker-socket",
						MountPath: "/var/run/docker.sock",
					},
					{
						Name:      "gpg",
						MountPath: "/root/gpg",
					},
				},
			}},
			RestartPolicy:      corev1.RestartPolicyNever,
			ServiceAccountName: serviceAccount,
			Volumes: []corev1.Volume{
				{
					Name: "docker-socket",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/run/docker.sock",
						},
					},
				},
				{
					Name: "gpg",
					VolumeSource: corev1.VolumeSource{
						Secret: &corev1.SecretVolumeSource{
							SecretName: gpgSecret,
						},
					},
				},
			},
		},
	}

	return pod, nil
}
