// Code generated for package assets by go-bindata DO NOT EDIT. (@generated)
// sources:
// manifests/bootstrap-pod-v2.yaml
// manifests/controllerconfig.crd.yaml
// manifests/machineconfigcontroller/clusterrole.yaml
// manifests/machineconfigcontroller/clusterrolebinding.yaml
// manifests/machineconfigcontroller/controllerconfig.yaml
// manifests/machineconfigcontroller/deployment.yaml
// manifests/machineconfigcontroller/events-clusterrole.yaml
// manifests/machineconfigcontroller/events-rolebinding-default.yaml
// manifests/machineconfigcontroller/events-rolebinding-target.yaml
// manifests/machineconfigcontroller/sa.yaml
// manifests/machineconfigdaemon/clusterrole.yaml
// manifests/machineconfigdaemon/clusterrolebinding.yaml
// manifests/machineconfigdaemon/cookie-secret.yaml
// manifests/machineconfigdaemon/daemonset.yaml
// manifests/machineconfigdaemon/events-clusterrole.yaml
// manifests/machineconfigdaemon/events-rolebinding-default.yaml
// manifests/machineconfigdaemon/events-rolebinding-target.yaml
// manifests/machineconfigdaemon/sa.yaml
// manifests/machineconfigserver/clusterrole.yaml
// manifests/machineconfigserver/clusterrolebinding.yaml
// manifests/machineconfigserver/csr-bootstrap-role-binding.yaml
// manifests/machineconfigserver/csr-renewal-role-binding.yaml
// manifests/machineconfigserver/daemonset.yaml
// manifests/machineconfigserver/kube-apiserver-serving-ca-configmap.yaml
// manifests/machineconfigserver/node-bootstrapper-sa.yaml
// manifests/machineconfigserver/node-bootstrapper-token.yaml
// manifests/machineconfigserver/sa.yaml
// manifests/master.machineconfigpool.yaml
// manifests/on-prem/coredns-corefile.tmpl
// manifests/on-prem/coredns.yaml
// manifests/on-prem/keepalived.conf.tmpl
// manifests/on-prem/keepalived.yaml
// manifests/worker.machineconfigpool.yaml
package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _manifestsBootstrapPodV2Yaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: bootstrap-machine-config-operator
  namespace: {{.TargetNamespace}}
  annotations:
    target.workload.openshift.io/management: '{"effect": "PreferredDuringScheduling"}'
spec:
  initContainers:
  - name: machine-config-controller
    image: {{.Images.MachineConfigOperator}}
    command: ["/usr/bin/machine-config-controller"]
    args:
    - "bootstrap"
    - "--manifest-dir=/etc/mcc/bootstrap"
    - "--dest-dir=/etc/mcs/bootstrap"
    - "--pull-secret=/etc/mcc/bootstrap/machineconfigcontroller-pull-secret"
    resources:
      limits:
        memory: 50Mi
      requests:
        cpu: 20m
        memory: 50Mi
    securityContext:
      privileged: true
    terminationMessagePolicy: FallbackToLogsOnError
    volumeMounts:
    - name: bootstrap-manifests
      mountPath: /etc/mcc/bootstrap
    - name: server-basedir
      mountPath: /etc/mcs/bootstrap
  containers:
  - name: machine-config-server
    image: {{.Images.MachineConfigOperator}}
    command: ["/usr/bin/machine-config-server"]
    args:
      - "bootstrap"
    terminationMessagePolicy: FallbackToLogsOnError
    volumeMounts:
    - name: server-certs
      mountPath: /etc/ssl/mcs
    - name: bootstrap-kubeconfig
      mountPath: /etc/kubernetes/kubeconfig
    - name: server-basedir
      mountPath: /etc/mcs/bootstrap
    securityContext:
      privileged: true
  hostNetwork: true
  tolerations:
    - key: node-role.kubernetes.io/master
      operator: Exists
      effect: NoSchedule
  restartPolicy: Always
  volumes:
  - name: server-certs
    hostPath:
      path: /etc/ssl/mcs
  - name: bootstrap-kubeconfig
    hostPath:
      path: /etc/mcs/kubeconfig
  - name: server-basedir
    hostPath:
      path: /etc/mcs/bootstrap
  - name: bootstrap-manifests
    hostPath:
      path: /etc/mcc/bootstrap
`)

func manifestsBootstrapPodV2YamlBytes() ([]byte, error) {
	return _manifestsBootstrapPodV2Yaml, nil
}

func manifestsBootstrapPodV2Yaml() (*asset, error) {
	bytes, err := manifestsBootstrapPodV2YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/bootstrap-pod-v2.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsControllerconfigCrdYaml = []byte(`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  # name must match the spec fields below, and be in the form: <plural>.<group>
  name: controllerconfigs.machineconfiguration.openshift.io
  labels:
    "openshift.io/operator-managed": ""
spec:
  # group name to use for REST API: /apis/<group>/<version>
  group: machineconfiguration.openshift.io
  # either Namespaced or Cluster
  scope: Cluster
  names:
    # plural name to be used in the URL: /apis/<group>/<version>/<plural>
    plural: controllerconfigs
    # singular name to be used as an alias on the CLI and for display
    singular: controllerconfig
    # kind is normally the PascalCased singular type. Your resource manifests use this.
    kind: ControllerConfig
  # list of versions supported by this CustomResourceDefinition
  versions:
  - name: v1
    # Each version can be enabled/disabled by Served flag.
    served: true
    # One and only one version must be marked as the storage version.
    storage: true
    subresources:
      status: {}
    schema:
      openAPIV3Schema:
        description: ControllerConfig describes configuration for MachineConfigController.
          This is currently only used to drive the MachineConfig objects generated
          by the TemplateController.
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ControllerConfigSpec is the spec for ControllerConfig resource.
            properties:
              additionalTrustBundle:
                description: additionalTrustBundle is a certificate bundle that will
                  be added to the nodes trusted certificate store.
                format: byte
                nullable: true
                type: string
              cloudProviderCAData:
                description: cloudProvider specifies the cloud provider CA data
                format: byte
                nullable: true
                type: string
              cloudProviderConfig:
                description: cloudProviderConfig is the configuration for the given
                  cloud provider
                type: string
              clusterDNSIP:
                description: clusterDNSIP is the cluster DNS IP address
                type: string
              dns:
                description: dns holds the cluster dns details
                nullable: true
                properties:
                  apiVersion:
                    description: 'APIVersion defines the versioned schema of this
                      representation of an object. Servers should convert recognized
                      schemas to the latest internal value, and may reject unrecognized
                      values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
                    type: string
                  kind:
                    description: 'Kind is a string value representing the REST resource
                      this object represents. Servers may infer this from the endpoint
                      the client submits requests to. Cannot be updated. In CamelCase.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  metadata:
                    type: object
                  spec:
                    description: spec holds user settable values for configuration
                    properties:
                      baseDomain:
                        description: "baseDomain is the base domain of the cluster.
                          All managed DNS records will be sub-domains of this base.
                          \n For example, given the base domain `+"`"+`openshift.example.com`+"`"+`,
                          an API server DNS record may be created for `+"`"+`cluster-api.openshift.example.com`+"`"+`.
                          \n Once set, this field cannot be changed."
                        type: string
                      privateZone:
                        description: "privateZone is the location where all the DNS
                          records that are only available internally to the cluster
                          exist. \n If this field is nil, no private records should
                          be created. \n Once set, this field cannot be changed."
                        properties:
                          id:
                            description: "id is the identifier that can be used to
                              find the DNS hosted zone. \n on AWS zone can be fetched
                              using `+"`"+`ID`+"`"+` as id in [1] on Azure zone can be fetched
                              using `+"`"+`ID`+"`"+` as a pre-determined name in [2], on GCP zone
                              can be fetched using `+"`"+`ID`+"`"+` as a pre-determined name in
                              [3]. \n [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options
                              [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show
                              [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get"
                            type: string
                          tags:
                            additionalProperties:
                              type: string
                            description: "tags can be used to query the DNS hosted
                              zone. \n on AWS, resourcegroupstaggingapi [1] can be
                              used to fetch a zone using `+"`"+`Tags`+"`"+` as tag-filters, \n
                              [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options"
                            type: object
                        type: object
                      publicZone:
                        description: "publicZone is the location where all the DNS
                          records that are publicly accessible to the internet exist.
                          \n If this field is nil, no public records should be created.
                          \n Once set, this field cannot be changed."
                        properties:
                          id:
                            description: "id is the identifier that can be used to
                              find the DNS hosted zone. \n on AWS zone can be fetched
                              using `+"`"+`ID`+"`"+` as id in [1] on Azure zone can be fetched
                              using `+"`"+`ID`+"`"+` as a pre-determined name in [2], on GCP zone
                              can be fetched using `+"`"+`ID`+"`"+` as a pre-determined name in
                              [3]. \n [1]: https://docs.aws.amazon.com/cli/latest/reference/route53/get-hosted-zone.html#options
                              [2]: https://docs.microsoft.com/en-us/cli/azure/network/dns/zone?view=azure-cli-latest#az-network-dns-zone-show
                              [3]: https://cloud.google.com/dns/docs/reference/v1/managedZones/get"
                            type: string
                          tags:
                            additionalProperties:
                              type: string
                            description: "tags can be used to query the DNS hosted
                              zone. \n on AWS, resourcegroupstaggingapi [1] can be
                              used to fetch a zone using `+"`"+`Tags`+"`"+` as tag-filters, \n
                              [1]: https://docs.aws.amazon.com/cli/latest/reference/resourcegroupstaggingapi/get-resources.html#options"
                            type: object
                        type: object
                    type: object
                  status:
                    description: status holds observed values from the cluster. They
                      may not be overridden.
                    type: object
                required:
                - spec
                type: object
              etcdDiscoveryDomain:
                description: etcdDiscoveryDomain is deprecated, use Infra.Status.EtcdDiscoveryDomain
                  instead
                type: string
              images:
                additionalProperties:
                  type: string
                description: images is map of images that are used by the controller
                  to render templates under ./templates/
                type: object
              infra:
                description: infra holds the infrastructure details
                nullable: true
                properties:
                  apiVersion:
                    description: 'APIVersion defines the versioned schema of this
                      representation of an object. Servers should convert recognized
                      schemas to the latest internal value, and may reject unrecognized
                      values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
                    type: string
                  kind:
                    description: 'Kind is a string value representing the REST resource
                      this object represents. Servers may infer this from the endpoint
                      the client submits requests to. Cannot be updated. In CamelCase.
                      More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  metadata:
                    type: object
                  spec:
                    description: spec holds user settable values for configuration
                    properties:
                      cloudConfig:
                        description: "cloudConfig is a reference to a ConfigMap containing
                          the cloud provider configuration file. This configuration
                          file is used to configure the Kubernetes cloud provider
                          integration when using the built-in cloud provider integration
                          or the external cloud controller manager. The namespace
                          for this config map is openshift-config. \n cloudConfig
                          should only be consumed by the kube_cloud_config controller.
                          The controller is responsible for using the user configuration
                          in the spec for various platforms and combining that with
                          the user provided ConfigMap in this field to create a stitched
                          kube cloud config. The controller generates a ConfigMap
                          `+"`"+`kube-cloud-config`+"`"+` in `+"`"+`openshift-config-managed`+"`"+` namespace
                          with the kube cloud config is stored in `+"`"+`cloud.conf`+"`"+` key.
                          All the clients are expected to use the generated ConfigMap
                          only."
                        properties:
                          key:
                            description: Key allows pointing to a specific key/value
                              inside of the configmap.  This is useful for logical
                              file references.
                            type: string
                          name:
                            type: string
                        type: object
                      platformSpec:
                        description: platformSpec holds desired information specific
                          to the underlying infrastructure provider.
                        properties:
                          aws:
                            description: AWS contains settings specific to the Amazon
                              Web Services infrastructure provider.
                            properties:
                              serviceEndpoints:
                                description: serviceEndpoints list contains custom
                                  endpoints which will override default service endpoint
                                  of AWS Services. There must be only one ServiceEndpoint
                                  for a service.
                                items:
                                  description: AWSServiceEndpoint store the configuration
                                    of a custom url to override existing defaults
                                    of AWS Services.
                                  properties:
                                    name:
                                      description: name is the name of the AWS service.
                                        The list of all the service names can be found
                                        at https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html
                                        This must be provided and cannot be empty.
                                      pattern: ^[a-z0-9-]+$
                                      type: string
                                    url:
                                      description: url is fully qualified URI with
                                        scheme https, that overrides the default generated
                                        endpoint for a client. This must be provided
                                        and cannot be empty.
                                      pattern: ^https://
                                      type: string
                                  type: object
                                type: array
                            type: object
                          azure:
                            description: Azure contains settings specific to the Azure
                              infrastructure provider.
                            type: object
                          baremetal:
                            description: BareMetal contains settings specific to the
                              BareMetal platform.
                            type: object
                          equinixMetal:
                            description: EquinixMetal contains settings specific to
                              the Equinix Metal infrastructure provider.
                            type: object
                          gcp:
                            description: GCP contains settings specific to the Google
                              Cloud Platform infrastructure provider.
                            type: object
                          ibmcloud:
                            description: IBMCloud contains settings specific to the
                              IBMCloud infrastructure provider.
                            type: object
                          kubevirt:
                            description: Kubevirt contains settings specific to the
                              kubevirt infrastructure provider.
                            type: object
                          openstack:
                            description: OpenStack contains settings specific to the
                              OpenStack infrastructure provider.
                            type: object
                          ovirt:
                            description: Ovirt contains settings specific to the oVirt
                              infrastructure provider.
                            type: object
                          type:
                            description: type is the underlying infrastructure provider
                              for the cluster. This value controls whether infrastructure
                              automation such as service load balancers, dynamic volume
                              provisioning, machine creation and deletion, and other
                              integrations are enabled. If None, no infrastructure
                              automation is enabled. Allowed values are "AWS", "Azure",
                              "BareMetal", "GCP", "Libvirt", "OpenStack", "VSphere",
                              "oVirt", "KubeVirt", "EquinixMetal", and "None". Individual
                              components may not support all platforms, and must handle
                              unrecognized platforms as None if they do not support
                              that platform.
                            enum:
                            - ""
                            - AWS
                            - Azure
                            - BareMetal
                            - GCP
                            - Libvirt
                            - OpenStack
                            - None
                            - VSphere
                            - oVirt
                            - IBMCloud
                            - KubeVirt
                            - EquinixMetal
                            type: string
                          vsphere:
                            description: VSphere contains settings specific to the
                              VSphere infrastructure provider.
                            type: object
                        type: object
                    type: object
                  status:
                    description: status holds observed values from the cluster. They
                      may not be overridden.
                    properties:
                      apiServerInternalURI:
                        description: apiServerInternalURL is a valid URI with scheme
                          'https', address and optionally a port (defaulting to 443).  apiServerInternalURL
                          can be used by components like kubelets, to contact the
                          Kubernetes API server using the infrastructure provider
                          rather than Kubernetes networking.
                        type: string
                      apiServerURL:
                        description: apiServerURL is a valid URI with scheme 'https',
                          address and optionally a port (defaulting to 443).  apiServerURL
                          can be used by components like the web console to tell users
                          where to find the Kubernetes API.
                        type: string
                      controlPlaneTopology:
                        default: HighlyAvailable
                        description: controlPlaneTopology expresses the expectations
                          for operands that normally run on control nodes. The default
                          is 'HighlyAvailable', which represents the behavior operators
                          have in a "normal" cluster. The 'SingleReplica' mode will
                          be used in single-node deployments and the operators should
                          not configure the operand for highly-available operation
                          The 'External' mode indicates that the control plane is
                          hosted externally to the cluster and that its components
                          are not visible within the cluster.
                        enum:
                        - HighlyAvailable
                        - SingleReplica
                        - External
                        type: string
                      etcdDiscoveryDomain:
                        description: 'etcdDiscoveryDomain is the domain used to fetch
                          the SRV records for discovering etcd servers and clients.
                          For more info: https://github.com/etcd-io/etcd/blob/329be66e8b3f9e2e6af83c123ff89297e49ebd15/Documentation/op-guide/clustering.md#dns-discovery
                          deprecated: as of 4.7, this field is no longer set or honored.  It
                          will be removed in a future release.'
                        type: string
                      infrastructureName:
                        description: infrastructureName uniquely identifies a cluster
                          with a human friendly name. Once set it should not be changed.
                          Must be of max length 27 and must have only alphanumeric
                          or hyphen characters.
                        type: string
                      infrastructureTopology:
                        default: HighlyAvailable
                        description: infrastructureTopology expresses the expectations
                          for infrastructure services that do not run on control plane
                          nodes, usually indicated by a node selector for a `+"`"+`role`+"`"+`
                          value other than `+"`"+`master`+"`"+`. The default is 'HighlyAvailable',
                          which represents the behavior operators have in a "normal"
                          cluster. The 'SingleReplica' mode will be used in single-node
                          deployments and the operators should not configure the operand
                          for highly-available operation
                        enum:
                        - HighlyAvailable
                        - SingleReplica
                        - External
                        type: string
                      platform:
                        description: "platform is the underlying infrastructure provider
                          for the cluster. \n Deprecated: Use platformStatus.type
                          instead."
                        enum:
                        - ""
                        - AWS
                        - Azure
                        - BareMetal
                        - GCP
                        - Libvirt
                        - OpenStack
                        - None
                        - VSphere
                        - oVirt
                        - IBMCloud
                        - KubeVirt
                        - EquinixMetal
                        type: string
                      platformStatus:
                        description: platformStatus holds status information specific
                          to the underlying infrastructure provider.
                        properties:
                          aws:
                            description: AWS contains settings specific to the Amazon
                              Web Services infrastructure provider.
                            properties:
                              region:
                                description: region holds the default AWS region for
                                  new AWS resources created by the cluster.
                                type: string
                              resourceTags:
                                description: resourceTags is a list of additional
                                  tags to apply to AWS resources created for the cluster.
                                  See https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
                                  for information on tagging AWS resources. AWS supports
                                  a maximum of 50 tags per resource. OpenShift reserves
                                  25 tags for its use, leaving 25 tags available for
                                  the user.
                                items:
                                  description: AWSResourceTag is a tag to apply to
                                    AWS resources created for the cluster.
                                  properties:
                                    key:
                                      description: key is the key of the tag
                                      maxLength: 128
                                      minLength: 1
                                      pattern: ^[0-9A-Za-z_.:/=+-@]+$
                                      type: string
                                    value:
                                      description: value is the value of the tag.
                                        Some AWS service do not support empty values.
                                        Since tags are added to resources in many
                                        services, the length of the tag value must
                                        meet the requirements of all services.
                                      maxLength: 256
                                      minLength: 1
                                      pattern: ^[0-9A-Za-z_.:/=+-@]+$
                                      type: string
                                  required:
                                  - key
                                  - value
                                  type: object
                                maxItems: 25
                                type: array
                              serviceEndpoints:
                                description: ServiceEndpoints list contains custom
                                  endpoints which will override default service endpoint
                                  of AWS Services. There must be only one ServiceEndpoint
                                  for a service.
                                items:
                                  description: AWSServiceEndpoint store the configuration
                                    of a custom url to override existing defaults
                                    of AWS Services.
                                  properties:
                                    name:
                                      description: name is the name of the AWS service.
                                        The list of all the service names can be found
                                        at https://docs.aws.amazon.com/general/latest/gr/aws-service-information.html
                                        This must be provided and cannot be empty.
                                      pattern: ^[a-z0-9-]+$
                                      type: string
                                    url:
                                      description: url is fully qualified URI with
                                        scheme https, that overrides the default generated
                                        endpoint for a client. This must be provided
                                        and cannot be empty.
                                      pattern: ^https://
                                      type: string
                                  type: object
                                type: array
                            type: object
                          azure:
                            description: Azure contains settings specific to the Azure
                              infrastructure provider.
                            properties:
                              armEndpoint:
                                description: armEndpoint specifies a URL to use for
                                  resource management in non-soverign clouds such
                                  as Azure Stack.
                                type: string
                              cloudName:
                                description: cloudName is the name of the Azure cloud
                                  environment which can be used to configure the Azure
                                  SDK with the appropriate Azure API endpoints. If
                                  empty, the value is equal to `+"`"+`AzurePublicCloud`+"`"+`.
                                enum:
                                - ""
                                - AzurePublicCloud
                                - AzureUSGovernmentCloud
                                - AzureChinaCloud
                                - AzureGermanCloud
                                - AzureStackCloud
                                type: string
                              networkResourceGroupName:
                                description: networkResourceGroupName is the Resource
                                  Group for network resources like the Virtual Network
                                  and Subnets used by the cluster. If empty, the value
                                  is same as ResourceGroupName.
                                type: string
                              resourceGroupName:
                                description: resourceGroupName is the Resource Group
                                  for new Azure resources created for the cluster.
                                type: string
                            type: object
                          baremetal:
                            description: BareMetal contains settings specific to the
                              BareMetal platform.
                            properties:
                              apiServerInternalIP:
                                description: apiServerInternalIP is an IP address
                                  to contact the Kubernetes API server that can be
                                  used by components inside the cluster, like kubelets
                                  using the infrastructure rather than Kubernetes
                                  networking. It is the IP that the Infrastructure.status.apiServerInternalURI
                                  points to. It is the IP for a self-hosted load balancer
                                  in front of the API servers.
                                type: string
                              ingressIP:
                                description: ingressIP is an external IP which routes
                                  to the default ingress controller. The IP is a suitable
                                  target of a wildcard DNS record used to resolve
                                  default route host names.
                                type: string
                              nodeDNSIP:
                                description: nodeDNSIP is the IP address for the internal
                                  DNS used by the nodes. Unlike the one managed by
                                  the DNS operator, `+"`"+`NodeDNSIP`+"`"+` provides name resolution
                                  for the nodes themselves. There is no DNS-as-a-service
                                  for BareMetal deployments. In order to minimize
                                  necessary changes to the datacenter DNS, a DNS service
                                  is hosted as a static pod to serve those hostnames
                                  to the nodes in the cluster.
                                type: string
                            type: object
                          equinixMetal:
                            description: EquinixMetal contains settings specific to
                              the Equinix Metal infrastructure provider.
                            properties:
                              apiServerInternalIP:
                                description: apiServerInternalIP is an IP address
                                  to contact the Kubernetes API server that can be
                                  used by components inside the cluster, like kubelets
                                  using the infrastructure rather than Kubernetes
                                  networking. It is the IP that the Infrastructure.status.apiServerInternalURI
                                  points to. It is the IP for a self-hosted load balancer
                                  in front of the API servers.
                                type: string
                              ingressIP:
                                description: ingressIP is an external IP which routes
                                  to the default ingress controller. The IP is a suitable
                                  target of a wildcard DNS record used to resolve
                                  default route host names.
                                type: string
                            type: object
                          gcp:
                            description: GCP contains settings specific to the Google
                              Cloud Platform infrastructure provider.
                            properties:
                              projectID:
                                description: resourceGroupName is the Project ID for
                                  new GCP resources created for the cluster.
                                type: string
                              region:
                                description: region holds the region for new GCP resources
                                  created for the cluster.
                                type: string
                            type: object
                          ibmcloud:
                            description: IBMCloud contains settings specific to the
                              IBMCloud infrastructure provider.
                            properties:
                              cisInstanceCRN:
                                description: CISInstanceCRN is the CRN of the Cloud
                                  Internet Services instance managing the DNS zone
                                  for the cluster's base domain
                                type: string
                              location:
                                description: Location is where the cluster has been
                                  deployed
                                type: string
                              providerType:
                                description: ProviderType indicates the type of cluster
                                  that was created
                                type: string
                              resourceGroupName:
                                description: ResourceGroupName is the Resource Group
                                  for new IBMCloud resources created for the cluster.
                                type: string
                            type: object
                          kubevirt:
                            description: Kubevirt contains settings specific to the
                              kubevirt infrastructure provider.
                            properties:
                              apiServerInternalIP:
                                description: apiServerInternalIP is an IP address
                                  to contact the Kubernetes API server that can be
                                  used by components inside the cluster, like kubelets
                                  using the infrastructure rather than Kubernetes
                                  networking. It is the IP that the Infrastructure.status.apiServerInternalURI
                                  points to. It is the IP for a self-hosted load balancer
                                  in front of the API servers.
                                type: string
                              ingressIP:
                                description: ingressIP is an external IP which routes
                                  to the default ingress controller. The IP is a suitable
                                  target of a wildcard DNS record used to resolve
                                  default route host names.
                                type: string
                            type: object
                          openstack:
                            description: OpenStack contains settings specific to the
                              OpenStack infrastructure provider.
                            properties:
                              apiServerInternalIP:
                                description: apiServerInternalIP is an IP address
                                  to contact the Kubernetes API server that can be
                                  used by components inside the cluster, like kubelets
                                  using the infrastructure rather than Kubernetes
                                  networking. It is the IP that the Infrastructure.status.apiServerInternalURI
                                  points to. It is the IP for a self-hosted load balancer
                                  in front of the API servers.
                                type: string
                              cloudName:
                                description: cloudName is the name of the desired
                                  OpenStack cloud in the client configuration file
                                  (`+"`"+`clouds.yaml`+"`"+`).
                                type: string
                              ingressIP:
                                description: ingressIP is an external IP which routes
                                  to the default ingress controller. The IP is a suitable
                                  target of a wildcard DNS record used to resolve
                                  default route host names.
                                type: string
                              nodeDNSIP:
                                description: nodeDNSIP is the IP address for the internal
                                  DNS used by the nodes. Unlike the one managed by
                                  the DNS operator, `+"`"+`NodeDNSIP`+"`"+` provides name resolution
                                  for the nodes themselves. There is no DNS-as-a-service
                                  for OpenStack deployments. In order to minimize
                                  necessary changes to the datacenter DNS, a DNS service
                                  is hosted as a static pod to serve those hostnames
                                  to the nodes in the cluster.
                                type: string
                            type: object
                          ovirt:
                            description: Ovirt contains settings specific to the oVirt
                              infrastructure provider.
                            properties:
                              apiServerInternalIP:
                                description: apiServerInternalIP is an IP address
                                  to contact the Kubernetes API server that can be
                                  used by components inside the cluster, like kubelets
                                  using the infrastructure rather than Kubernetes
                                  networking. It is the IP that the Infrastructure.status.apiServerInternalURI
                                  points to. It is the IP for a self-hosted load balancer
                                  in front of the API servers.
                                type: string
                              ingressIP:
                                description: ingressIP is an external IP which routes
                                  to the default ingress controller. The IP is a suitable
                                  target of a wildcard DNS record used to resolve
                                  default route host names.
                                type: string
                              nodeDNSIP:
                                description: 'deprecated: as of 4.6, this field is
                                  no longer set or honored.  It will be removed in
                                  a future release.'
                                type: string
                            type: object
                          type:
                            description: "type is the underlying infrastructure provider
                              for the cluster. This value controls whether infrastructure
                              automation such as service load balancers, dynamic volume
                              provisioning, machine creation and deletion, and other
                              integrations are enabled. If None, no infrastructure
                              automation is enabled. Allowed values are \"AWS\", \"Azure\",
                              \"BareMetal\", \"GCP\", \"Libvirt\", \"OpenStack\",
                              \"VSphere\", \"oVirt\", \"EquinixMetal\", and \"None\".
                              Individual components may not support all platforms,
                              and must handle unrecognized platforms as None if they
                              do not support that platform. \n This value will be
                              synced with to the `+"`"+`status.platform`+"`"+` and `+"`"+`status.platformStatus.type`+"`"+`.
                              Currently this value cannot be changed once set."
                            enum:
                            - ""
                            - AWS
                            - Azure
                            - BareMetal
                            - GCP
                            - Libvirt
                            - OpenStack
                            - None
                            - VSphere
                            - oVirt
                            - IBMCloud
                            - KubeVirt
                            - EquinixMetal
                            type: string
                          vsphere:
                            description: VSphere contains settings specific to the
                              VSphere infrastructure provider.
                            properties:
                              apiServerInternalIP:
                                description: apiServerInternalIP is an IP address
                                  to contact the Kubernetes API server that can be
                                  used by components inside the cluster, like kubelets
                                  using the infrastructure rather than Kubernetes
                                  networking. It is the IP that the Infrastructure.status.apiServerInternalURI
                                  points to. It is the IP for a self-hosted load balancer
                                  in front of the API servers.
                                type: string
                              ingressIP:
                                description: ingressIP is an external IP which routes
                                  to the default ingress controller. The IP is a suitable
                                  target of a wildcard DNS record used to resolve
                                  default route host names.
                                type: string
                              nodeDNSIP:
                                description: nodeDNSIP is the IP address for the internal
                                  DNS used by the nodes. Unlike the one managed by
                                  the DNS operator, `+"`"+`NodeDNSIP`+"`"+` provides name resolution
                                  for the nodes themselves. There is no DNS-as-a-service
                                  for vSphere deployments. In order to minimize necessary
                                  changes to the datacenter DNS, a DNS service is
                                  hosted as a static pod to serve those hostnames
                                  to the nodes in the cluster.
                                type: string
                            type: object
                        type: object
                    type: object
                required:
                - spec
                type: object
              ipFamilies:
                description: ipFamilies indicates the IP families in use by the cluster
                  network
                type: string
              kubeAPIServerServingCAData:
                description: kubeAPIServerServingCAData managed Kubelet to API Server
                  Cert... Rotated automatically
                format: byte
                type: string
              networkType:
                description: 'networkType holds the type of network the cluster is
                  using XXX: this is temporary and will be dropped as soon as possible
                  in favor of a better support to start network related services the
                  proper way. Nobody is also changing this once the cluster is up
                  and running the first time, so, disallow regeneration if this changes.'
                type: string
              osImageURL:
                description: osImageURL is the location of the container image that
                  contains the OS update payload. Its value is taken from the data.osImageURL
                  field on the machine-config-osimageurl ConfigMap.
                type: string
              platform:
                description: platform is deprecated, use Infra.Status.PlatformStatus.Type
                  instead
                type: string
              proxy:
                description: proxy holds the current proxy configuration for the nodes
                nullable: true
                properties:
                  httpProxy:
                    description: httpProxy is the URL of the proxy for HTTP requests.
                    type: string
                  httpsProxy:
                    description: httpsProxy is the URL of the proxy for HTTPS requests.
                    type: string
                  noProxy:
                    description: noProxy is a comma-separated list of hostnames and/or
                      CIDRs for which the proxy should not be used.
                    type: string
                type: object
              pullSecret:
                description: pullSecret is the default pull secret that needs to be
                  installed on all machines.
                properties:
                  apiVersion:
                    description: API version of the referent.
                    type: string
                  fieldPath:
                    description: 'If referring to a piece of an object instead of
                      an entire object, this string should contain a valid JSON/Go
                      field access statement, such as desiredState.manifest.containers[2].
                      For example, if the object reference is to a container within
                      a pod, this would take on a value like: "spec.containers{name}"
                      (where "name" refers to the name of the container that triggered
                      the event) or if no container name is specified "spec.containers[2]"
                      (container with index 2 in this pod). This syntax is chosen
                      only to have some well-defined way of referencing a part of
                      an object. TODO: this design is not final and this field is
                      subject to change in the future.'
                    type: string
                  kind:
                    description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                  resourceVersion:
                    description: 'Specific resourceVersion to which this reference
                      is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                    type: string
                  uid:
                    description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                    type: string
                type: object
              releaseImage:
                description: releaseImage is the image used when installing the cluster
                type: string
              rootCAData:
                description: rootCAData specifies the root CA data
                format: byte
                type: string
            required:
            - additionalTrustBundle
            - cloudProviderCAData
            - cloudProviderConfig
            - clusterDNSIP
            - dns
            - images
            - infra
            - ipFamilies
            - kubeAPIServerServingCAData
            - osImageURL
            - proxy
            - releaseImage
            - rootCAData
            type: object
          status:
            description: ControllerConfigStatus is the status for ControllerConfig
            properties:
              conditions:
                description: conditions represents the latest available observations
                  of current state.
                items:
                  description: ControllerConfigStatusCondition contains condition
                    information for ControllerConfigStatus
                  properties:
                    lastTransitionTime:
                      description: lastTransitionTime is the time of the last update
                        to the current status object.
                      format: date-time
                      nullable: true
                      type: string
                    message:
                      description: message provides additional information about the
                        current condition. This is only to be consumed by humans.
                      type: string
                    reason:
                      description: reason is the reason for the condition's last transition.  Reasons
                        are PascalCase
                      type: string
                    status:
                      description: status of the condition, one of True, False, Unknown.
                      type: string
                    type:
                      description: type specifies the state of the operator's reconciliation
                        functionality.
                      type: string
                  required:
                  - lastTransitionTime
                  - status
                  - type
                  type: object
                type: array
              observedGeneration:
                description: observedGeneration represents the generation observed
                  by the controller.
                format: int64
                type: integer
            type: object
        required:
        - spec
        type: object
`)

func manifestsControllerconfigCrdYamlBytes() ([]byte, error) {
	return _manifestsControllerconfigCrdYaml, nil
}

func manifestsControllerconfigCrdYaml() (*asset, error) {
	bytes, err := manifestsControllerconfigCrdYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/controllerconfig.crd.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch", "patch"]
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["*"]
  verbs: ["*"]
- apiGroups: [""]
  resources: ["configmaps", "secrets"]
  verbs: ["*"]
- apiGroups: ["config.openshift.io"]
  resources: ["images", "clusterversions", "featuregates"]
  verbs: ["*"]
- apiGroups: ["config.openshift.io"]
  resources: ["schedulers", "apiservers"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["operator.openshift.io"]
  resources: ["imagecontentsourcepolicies"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["operator.openshift.io"]
  resources: ["etcds"]
  verbs: ["get", "list", "watch"]
`)

func manifestsMachineconfigcontrollerClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerClusterroleYaml, nil
}

func manifestsMachineconfigcontrollerClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-controller
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerClusterrolebindingYaml, nil
}

func manifestsMachineconfigcontrollerClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerControllerconfigYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: ControllerConfig
metadata:
  name: machine-config-controller
  annotations:
    machineconfiguration.openshift.io/generated-by-version: "{{ .Version }}"
spec:
{{toYAML .ControllerConfig | toString | indent 2}}
`)

func manifestsMachineconfigcontrollerControllerconfigYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerControllerconfigYaml, nil
}

func manifestsMachineconfigcontrollerControllerconfigYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerControllerconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/controllerconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: machine-config-controller
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-controller
  template:
    metadata:
      labels:
        k8s-app: machine-config-controller
      annotations:
        target.workload.openshift.io/management: '{"effect": "PreferredDuringScheduling"}'
    spec:
      containers:
      - name: machine-config-controller
        image: {{.Images.MachineConfigOperator}}
        command: ["/usr/bin/machine-config-controller"]
        args:
        - "start"
        - "--resourcelock-namespace={{.TargetNamespace}}"
        - "--v=2"
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
        terminationMessagePolicy: FallbackToLogsOnError
      serviceAccountName: machine-config-controller
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      restartPolicy: Always
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Exists
        effect: "NoSchedule"
      - key: "node.kubernetes.io/unreachable"
        operator: "Exists"
        effect: "NoExecute"
        tolerationSeconds: 120
      - key: "node.kubernetes.io/not-ready"
        operator: "Exists"
        effect: "NoExecute"
        tolerationSeconds: 120
`)

func manifestsMachineconfigcontrollerDeploymentYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerDeploymentYaml, nil
}

func manifestsMachineconfigcontrollerDeploymentYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerEventsClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-controller-events
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "patch"]
`)

func manifestsMachineconfigcontrollerEventsClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerEventsClusterroleYaml, nil
}

func manifestsMachineconfigcontrollerEventsClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerEventsClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/events-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerEventsRolebindingDefaultYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: machine-config-controller-events
  namespace: default
roleRef:
  kind: ClusterRole
  name: machine-config-controller-events
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerEventsRolebindingDefaultYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerEventsRolebindingDefaultYaml, nil
}

func manifestsMachineconfigcontrollerEventsRolebindingDefaultYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerEventsRolebindingDefaultYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/events-rolebinding-default.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerEventsRolebindingTargetYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: machine-config-controller-events
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-controller-events
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerEventsRolebindingTargetYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerEventsRolebindingTargetYaml, nil
}

func manifestsMachineconfigcontrollerEventsRolebindingTargetYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerEventsRolebindingTargetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/events-rolebinding-target.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigcontrollerSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-controller
`)

func manifestsMachineconfigcontrollerSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigcontrollerSaYaml, nil
}

func manifestsMachineconfigcontrollerSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigcontrollerSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigcontroller/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["nodes"]
  verbs: ["get", "list", "watch", "patch", "update"]
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["*"]
- apiGroups: ["extensions"]
  resources: ["daemonsets"]
  verbs: ["get"]
- apiGroups: ["apps"]
  resources: ["daemonsets"]
  verbs: ["get"]
- apiGroups: [""]
  resources: ["pods/eviction"]
  verbs: ["create"]
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["machineconfigs"]
  verbs: ["*"]
- apiGroups:
  - authentication.k8s.io
  resources:
  - tokenreviews
  - subjectaccessreviews
  verbs:
  - create
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create
`)

func manifestsMachineconfigdaemonClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonClusterroleYaml, nil
}

func manifestsMachineconfigdaemonClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-daemon
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
---
# Bind auth-delegator role to the MCD service account
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  apiGroup: rbac.authorization.k8s.io
  name: system:auth-delegator
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonClusterrolebindingYaml, nil
}

func manifestsMachineconfigdaemonClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonCookieSecretYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  name: cookie-secret
  namespace: {{.TargetNamespace}}
type: Opaque
data:
  cookie-secret: {{.GenerateProxyCookieSecret}}
`)

func manifestsMachineconfigdaemonCookieSecretYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonCookieSecretYaml, nil
}

func manifestsMachineconfigdaemonCookieSecretYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonCookieSecretYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/cookie-secret.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: machine-config-daemon
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-daemon
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 10%
  template:
    metadata:
      name: machine-config-daemon
      labels:
        k8s-app: machine-config-daemon
      annotations:
        target.workload.openshift.io/management: '{"effect": "PreferredDuringScheduling"}'
    spec:
      containers:
      - name: machine-config-daemon
        image: {{.Images.MachineConfigOperator}}
        command: ["/usr/bin/machine-config-daemon"]
        args:
          - "start"
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
        securityContext:
          privileged: true
        terminationMessagePolicy: FallbackToLogsOnError
        volumeMounts:
          - mountPath: /rootfs
            name: rootfs
        env:
          - name: NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          {{if .ControllerConfig.Proxy}}
          {{if .ControllerConfig.Proxy.HTTPProxy}}
          - name: HTTP_PROXY
            value: {{.ControllerConfig.Proxy.HTTPProxy}}
          {{end}}
          {{if .ControllerConfig.Proxy.HTTPSProxy}}
          - name: HTTPS_PROXY
            value: {{.ControllerConfig.Proxy.HTTPSProxy}}
          {{end}}
          {{if .ControllerConfig.Proxy.NoProxy}}
          - name: NO_PROXY
            value: "{{.ControllerConfig.Proxy.NoProxy}}"
          {{end}}
          {{end}}
      - name: oauth-proxy
        image: {{.Images.OauthProxy}}
        ports:
        - containerPort: 9001
          name: metrics
          protocol: TCP
        args:
        - --https-address=:9001
        - --provider=openshift
        - --openshift-service-account=machine-config-daemon
        - --upstream=http://127.0.0.1:8797
        - --tls-cert=/etc/tls/private/tls.crt
        - --tls-key=/etc/tls/private/tls.key
        - --cookie-secret-file=/etc/tls/cookie-secret/cookie-secret
        - '--openshift-sar={"resource": "namespaces", "verb": "get"}'
        - '--openshift-delegate-urls={"/": {"resource": "namespaces", "verb": "get"}}'
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
        volumeMounts:
        - mountPath: /etc/tls/private
          name: proxy-tls
        - mountPath: /etc/tls/cookie-secret
          name: cookie-secret
      hostNetwork: true
      hostPID: true
      serviceAccountName: machine-config-daemon
      terminationGracePeriodSeconds: 600
      nodeSelector:
        kubernetes.io/os: linux
      priorityClassName: "system-node-critical"
      volumes:
        - name: rootfs
          hostPath:
            path: /
        - name: proxy-tls
          secret:
            secretName: proxy-tls
        - name: cookie-secret
          secret:
            secretName: cookie-secret
      tolerations:
      # MCD needs to run everywhere. Tolerate all taints.
      - operator: Exists
`)

func manifestsMachineconfigdaemonDaemonsetYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonDaemonsetYaml, nil
}

func manifestsMachineconfigdaemonDaemonsetYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonEventsClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-daemon-events
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: [""]
  resources: ["events"]
  verbs: ["create", "patch"]
`)

func manifestsMachineconfigdaemonEventsClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonEventsClusterroleYaml, nil
}

func manifestsMachineconfigdaemonEventsClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonEventsClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/events-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonEventsRolebindingDefaultYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: machine-config-daemon-events
  namespace: default
roleRef:
  kind: ClusterRole
  name: machine-config-daemon-events
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonEventsRolebindingDefaultYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonEventsRolebindingDefaultYaml, nil
}

func manifestsMachineconfigdaemonEventsRolebindingDefaultYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonEventsRolebindingDefaultYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/events-rolebinding-default.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonEventsRolebindingTargetYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: machine-config-daemon-events
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-daemon-events
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonEventsRolebindingTargetYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonEventsRolebindingTargetYaml, nil
}

func manifestsMachineconfigdaemonEventsRolebindingTargetYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonEventsRolebindingTargetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/events-rolebinding-target.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigdaemonSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-daemon
`)

func manifestsMachineconfigdaemonSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigdaemonSaYaml, nil
}

func manifestsMachineconfigdaemonSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigdaemonSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigdaemon/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
rules:
- apiGroups: ["machineconfiguration.openshift.io"]
  resources: ["machineconfigs", "machineconfigpools"]
  verbs: ["*"]
`)

func manifestsMachineconfigserverClusterroleYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverClusterroleYaml, nil
}

func manifestsMachineconfigserverClusterroleYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverClusterrolebindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
roleRef:
  kind: ClusterRole
  name: machine-config-server
subjects:
- kind: ServiceAccount
  namespace: {{.TargetNamespace}}
  name: machine-config-server
`)

func manifestsMachineconfigserverClusterrolebindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverClusterrolebindingYaml, nil
}

func manifestsMachineconfigserverClusterrolebindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverClusterrolebindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/clusterrolebinding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverCsrBootstrapRoleBindingYaml = []byte(`# system-bootstrap-node-bootstrapper lets serviceaccount `+"`"+`openshift-machine-config-operator/node-bootstrapper`+"`"+` tokens and nodes request CSRs.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system-bootstrap-node-bootstrapper
subjects:
- kind: ServiceAccount
  name: node-bootstrapper
  namespace: openshift-machine-config-operator
roleRef:
  kind: ClusterRole
  name: system:node-bootstrapper
  apiGroup: rbac.authorization.k8s.io`)

func manifestsMachineconfigserverCsrBootstrapRoleBindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverCsrBootstrapRoleBindingYaml, nil
}

func manifestsMachineconfigserverCsrBootstrapRoleBindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverCsrBootstrapRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/csr-bootstrap-role-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverCsrRenewalRoleBindingYaml = []byte(`# CSRRenewalRoleBindingTemplate instructs the csrapprover controller to
# automatically approve all CSRs made by nodes to renew their client
# certificates.
#
# This binding should be altered in the future to hold a list of node
# names instead of targeting `+"`"+`system:nodes`+"`"+` so we can revoke invidivual
# node's ability to renew its certs.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: system-bootstrap-node-renewal
subjects:
- kind: Group
  name: system:nodes
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: ClusterRole
  name: system:certificates.k8s.io:certificatesigningrequests:selfnodeclient
  apiGroup: rbac.authorization.k8s.io`)

func manifestsMachineconfigserverCsrRenewalRoleBindingYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverCsrRenewalRoleBindingYaml, nil
}

func manifestsMachineconfigserverCsrRenewalRoleBindingYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverCsrRenewalRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/csr-renewal-role-binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverDaemonsetYaml = []byte(`apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: machine-config-server
  namespace: {{.TargetNamespace}}
spec:
  selector:
    matchLabels:
      k8s-app: machine-config-server
  template:
    metadata:
      name: machine-config-server
      labels:
        k8s-app: machine-config-server
      annotations:
        target.workload.openshift.io/management: '{"effect": "PreferredDuringScheduling"}'
    spec:
      containers:
      - name: machine-config-server
        image: {{.Images.MachineConfigOperator}}
        command: ["/usr/bin/machine-config-server"]
        args:
          - "start"
          - "--apiserver-url={{.APIServerURL}}"
        resources:
          requests:
            cpu: 20m
            memory: 50Mi
        terminationMessagePolicy: FallbackToLogsOnError
        volumeMounts:
        - name: certs
          mountPath: /etc/ssl/mcs
        - name: node-bootstrap-token
          mountPath: /etc/mcs/bootstrap-token
      hostNetwork: true
      nodeSelector:
        node-role.kubernetes.io/master: ""
      priorityClassName: "system-cluster-critical"
      serviceAccountName: machine-config-server
      tolerations:
      - key: node-role.kubernetes.io/master
        operator: Exists
        effect: NoSchedule
      - key: node-role.kubernetes.io/etcd
        operator: Exists
        effect: NoSchedule
      volumes:
      - name: node-bootstrap-token
        secret:
          secretName: node-bootstrapper-token
      - name: certs
        secret:
          secretName: machine-config-server-tls
`)

func manifestsMachineconfigserverDaemonsetYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverDaemonsetYaml, nil
}

func manifestsMachineconfigserverDaemonsetYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/daemonset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  name: initial-kube-apiserver-server-ca
  namespace: openshift-config
data:
  ca-bundle.crt: |
{{.KubeAPIServerServingCA | indent 4}}
`)

func manifestsMachineconfigserverKubeApiserverServingCaConfigmapYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml, nil
}

func manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverKubeApiserverServingCaConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/kube-apiserver-serving-ca-configmap.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverNodeBootstrapperSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: node-bootstrapper
`)

func manifestsMachineconfigserverNodeBootstrapperSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverNodeBootstrapperSaYaml, nil
}

func manifestsMachineconfigserverNodeBootstrapperSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverNodeBootstrapperSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/node-bootstrapper-sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverNodeBootstrapperTokenYaml = []byte(`apiVersion: v1
kind: Secret
metadata:
  annotations:
    kubernetes.io/service-account.name: node-bootstrapper
  name: node-bootstrapper-token
  namespace: {{.TargetNamespace}}
type: kubernetes.io/service-account-token
`)

func manifestsMachineconfigserverNodeBootstrapperTokenYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverNodeBootstrapperTokenYaml, nil
}

func manifestsMachineconfigserverNodeBootstrapperTokenYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverNodeBootstrapperTokenYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/node-bootstrapper-token.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMachineconfigserverSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: {{.TargetNamespace}}
  name: machine-config-server
`)

func manifestsMachineconfigserverSaYamlBytes() ([]byte, error) {
	return _manifestsMachineconfigserverSaYaml, nil
}

func manifestsMachineconfigserverSaYaml() (*asset, error) {
	bytes, err := manifestsMachineconfigserverSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/machineconfigserver/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsMasterMachineconfigpoolYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfigPool
metadata:
  name: master
  labels:
    "operator.machineconfiguration.openshift.io/required-for-upgrade": ""
    "machineconfiguration.openshift.io/mco-built-in": ""
    "pools.operator.machineconfiguration.openshift.io/master": ""
spec:
  machineConfigSelector:
    matchLabels:
      "machineconfiguration.openshift.io/role": "master"
  nodeSelector:
    matchLabels:
      node-role.kubernetes.io/master: ""`)

func manifestsMasterMachineconfigpoolYamlBytes() ([]byte, error) {
	return _manifestsMasterMachineconfigpoolYaml, nil
}

func manifestsMasterMachineconfigpoolYaml() (*asset, error) {
	bytes, err := manifestsMasterMachineconfigpoolYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/master.machineconfigpool.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOnPremCorednsCorefileTmpl = []byte(`. {
    errors
    health :18080
    forward . {{`+"`"+`{{- range $upstream := .DNSUpstreams}} {{$upstream}}{{- end}}`+"`"+`}} {
        policy sequential
    }
    cache 30
    reload
    template IN {{`+"`"+`{{ .Cluster.IngressVIPRecordType }}`+"`"+`}} {{ .ControllerConfig.DNS.Spec.BaseDomain }} {
        match .*.apps.{{ .ControllerConfig.DNS.Spec.BaseDomain }}
        answer "{{`+"`"+`{{"{{ .Name }}"}}`+"`"+`}} 60 in {{`+"`"+`{{"{{ .Type }}"}}`+"`"+`}} {{ onPremPlatformIngressIP .ControllerConfig }}"
        fallthrough
    }
    template IN {{`+"`"+`{{ .Cluster.IngressVIPEmptyType }}`+"`"+`}} {{ .ControllerConfig.DNS.Spec.BaseDomain }} {
        match .*.apps.{{ .ControllerConfig.DNS.Spec.BaseDomain }}
        fallthrough
    }
    template IN {{`+"`"+`{{ .Cluster.APIVIPRecordType }}`+"`"+`}} {{ .ControllerConfig.DNS.Spec.BaseDomain }} {
        match api.{{ .ControllerConfig.DNS.Spec.BaseDomain }}
        answer "{{`+"`"+`{{"{{ .Name }}"}}`+"`"+`}} 60 in {{`+"`"+`{{"{{ .Type }}"}}`+"`"+`}} {{ onPremPlatformAPIServerInternalIP .ControllerConfig }}"
        fallthrough
    }
    template IN {{`+"`"+`{{ .Cluster.APIVIPEmptyType }}`+"`"+`}} {{ .ControllerConfig.DNS.Spec.BaseDomain }} {
        match api.{{ .ControllerConfig.DNS.Spec.BaseDomain }}
        fallthrough
    }
    template IN {{`+"`"+`{{ .Cluster.APIVIPRecordType }}`+"`"+`}} {{ .ControllerConfig.DNS.Spec.BaseDomain }} {
        match api-int.{{ .ControllerConfig.DNS.Spec.BaseDomain }}
        answer "{{`+"`"+`{{"{{ .Name }}"}}`+"`"+`}} 60 in {{`+"`"+`{{"{{ .Type }}"}}`+"`"+`}} {{ onPremPlatformAPIServerInternalIP .ControllerConfig }}"
        fallthrough
    }
    template IN {{`+"`"+`{{ .Cluster.APIVIPEmptyType }}`+"`"+`}} {{ .ControllerConfig.DNS.Spec.BaseDomain }} {
        match api-int.{{ .ControllerConfig.DNS.Spec.BaseDomain }}
        fallthrough
    }
}
`)

func manifestsOnPremCorednsCorefileTmplBytes() ([]byte, error) {
	return _manifestsOnPremCorednsCorefileTmpl, nil
}

func manifestsOnPremCorednsCorefileTmpl() (*asset, error) {
	bytes, err := manifestsOnPremCorednsCorefileTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/on-prem/coredns-corefile.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOnPremCorednsYaml = []byte(`---
kind: Pod
apiVersion: v1
metadata:
  name: coredns
  namespace: openshift-{{ onPremPlatformShortName .ControllerConfig }}-infra
  creationTimestamp:
  deletionGracePeriodSeconds: 65
  labels:
    app: {{ onPremPlatformShortName .ControllerConfig }}-infra-coredns
  annotations:
    target.workload.openshift.io/management: '{"effect": "PreferredDuringScheduling"}'
spec:
  volumes:
  - name: resource-dir
    hostPath:
      path: "/etc/kubernetes/static-pod-resources/coredns"
  - name: kubeconfig
    hostPath:
      path: "/etc/kubernetes/kubeconfig"
  - name: conf-dir
    empty-dir: {}
  - name: manifests
    hostPath:
      path: "/opt/openshift/manifests"
  initContainers:
  - name: render-config
    image: {{ .Images.BaremetalRuntimeCfgBootstrap }}
    command:
    - runtimecfg
    - render
    - "/etc/kubernetes/kubeconfig"
    - "--api-vip"
    - "{{ onPremPlatformAPIServerInternalIP .ControllerConfig }}"
    - "--ingress-vip"
    - "{{ onPremPlatformIngressIP .ControllerConfig }}"
    - "/config"
    - "--out-dir"
    - "/etc/coredns"
    - "--cluster-config"
    - "/opt/openshift/manifests/cluster-config.yaml"
    resources: {}
    volumeMounts:
    - name: kubeconfig
      mountPath: "/etc/kubernetes/kubeconfig"
    - name: resource-dir
      mountPath: "/config"
    - name: conf-dir
      mountPath: "/etc/coredns"
    - name: manifests
      mountPath: "/opt/openshift/manifests"
    imagePullPolicy: IfNotPresent
  containers:
  - name: coredns
    securityContext:
      privileged: true
    image: {{ .Images.CorednsBootstrap }}
    args:
    - "--conf"
    - "/etc/coredns/Corefile"
    resources:
      requests:
        cpu: 100m
        memory: 200Mi
    volumeMounts:
    - name: conf-dir
      mountPath: "/etc/coredns"
    livenessProbe:
      httpGet:
        path: /health
        port: 18080
        scheme: HTTP
      initialDelaySeconds: 60
      timeoutSeconds: 5
      successThreshold: 1
      failureThreshold: 5
    terminationMessagePolicy: FallbackToLogsOnError
  hostNetwork: true
  tolerations:
  - operator: Exists
  priorityClassName: system-node-critical
status: {}
`)

func manifestsOnPremCorednsYamlBytes() ([]byte, error) {
	return _manifestsOnPremCorednsYaml, nil
}

func manifestsOnPremCorednsYaml() (*asset, error) {
	bytes, err := manifestsOnPremCorednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/on-prem/coredns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOnPremKeepalivedConfTmpl = []byte(`# Configuration template for Keepalived, which is used to manage the DNS and
# API VIPs.
#
# For more information, see installer/data/data/bootstrap/baremetal/README.md
# in the installer repo.

{{`+"`"+`vrrp_instance {{.Cluster.Name}}_API {
    state BACKUP
    interface {{.VRRPInterface}}
    virtual_router_id {{.Cluster.APIVirtualRouterID }}
    priority 70
    advert_int 1
    {{ if .EnableUnicast }}
    unicast_src_ip {{.NonVirtualIP}}
    unicast_peer {
        {{range .LBConfig.Backends -}}
        {{.Address}}
        {{end}}
    }
    {{end}}
    authentication {
        auth_type PASS
        auth_pass {{.Cluster.Name}}_api_vip
    }
    virtual_ipaddress {
        {{ .Cluster.APIVIP }}/{{ .Cluster.VIPNetmask }}
    }
}`+"`"+`}}
`)

func manifestsOnPremKeepalivedConfTmplBytes() ([]byte, error) {
	return _manifestsOnPremKeepalivedConfTmpl, nil
}

func manifestsOnPremKeepalivedConfTmpl() (*asset, error) {
	bytes, err := manifestsOnPremKeepalivedConfTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/on-prem/keepalived.conf.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsOnPremKeepalivedYaml = []byte(`---
kind: Pod
apiVersion: v1
metadata:
  name: keepalived
  namespace: openshift-{{ onPremPlatformShortName .ControllerConfig }}-infra
  creationTimestamp:
  deletionGracePeriodSeconds: 65
  labels:
    app: {{ onPremPlatformShortName .ControllerConfig }}-infra-vrrp
  annotations:
    target.workload.openshift.io/management: '{"effect": "PreferredDuringScheduling"}'
spec:
  volumes:
  - name: resource-dir
    hostPath:
      path: "/etc/kubernetes/static-pod-resources/keepalived"
  - name: kubeconfig
    hostPath:
      path: "/etc/kubernetes/kubeconfig"
  - name: conf-dir
    hostPath:
      path: "/etc/keepalived"
  - name: manifests
    hostPath:
      path: "/opt/openshift/manifests"
  - name: run-dir
    empty-dir: {}
  containers:
  - name: keepalived
    securityContext:
      privileged: true
    image: {{.Images.KeepalivedBootstrap}}
    env:
      - name: NSS_SDB_USE_CACHE
        value: "no"
    command:
    - /bin/bash
    - -c
    - |
      #/bin/bash
      reload_keepalived()
      {
        if pid=$(pgrep -o keepalived); then
            kill -s SIGHUP "$pid"
        else
            /usr/sbin/keepalived -f /etc/keepalived/keepalived.conf --dont-fork --vrrp --log-detail --log-console &
        fi
      }
      stop_keepalived()
      {
        echo "Keepalived process stopped" >> /var/run/keepalived/stopped
        if pid=$(pgrep -o keepalived); then
            kill -s TERM "$pid"
        fi
      }

      msg_handler()
      {
        while read -r line; do
          echo "The client sent: $line" >&2
          # currently only 'reload' and 'stop' msgs are supported
          if [ "$line" = reload ]; then
              reload_keepalived
          elif  [ "$line" = stop ]; then
              stop_keepalived
          fi
        done
      }
      set -ex
      declare -r keepalived_sock="/var/run/keepalived/keepalived.sock"
      export -f msg_handler
      export -f reload_keepalived
      export -f stop_keepalived

      while [ -s "/var/run/keepalived/stopped" ]; do
         echo "Container stopped"
         sleep 60
      done
      if [ -s "/etc/keepalived/keepalived.conf" ]; then
          /usr/sbin/keepalived -f /etc/keepalived/keepalived.conf --dont-fork --vrrp --log-detail --log-console &
      fi
      rm -f "$keepalived_sock"
      socat UNIX-LISTEN:${keepalived_sock},fork system:'bash -c msg_handler'
    resources:
      requests:
        cpu: 100m
        memory: 200Mi
    volumeMounts:
    - name: conf-dir
      mountPath: "/etc/keepalived"
    - name: run-dir
      mountPath: "/var/run/keepalived"
    livenessProbe:
      exec:
        command:
        - /bin/bash
        - -c
        - |
          [[ -s /etc/keepalived/keepalived.conf ]] || \
          [[ -s /var/run/keepalived/stopped ]] || \
          kill -s SIGUSR1 "$(pgrep -o keepalived)" && ! grep -q "State = FAULT" /tmp/keepalived.data
      initialDelaySeconds: 20
    terminationMessagePolicy: FallbackToLogsOnError
    imagePullPolicy: IfNotPresent
  - name: keepalived-monitor
    image: {{ .Images.BaremetalRuntimeCfgBootstrap }}
    env:
      - name: ENABLE_UNICAST
        value: "{{ onPremPlatformKeepalivedEnableUnicast .ControllerConfig }}"
      - name: IS_BOOTSTRAP
        value: "yes"
    command:
    - dynkeepalived
    - "/etc/kubernetes/kubeconfig"
    - "/config/keepalived.conf.tmpl"
    - "/etc/keepalived/keepalived.conf"
    - "--api-vip"
    - "{{ onPremPlatformAPIServerInternalIP .ControllerConfig }}"
    - "--ingress-vip"
    - "{{ onPremPlatformIngressIP .ControllerConfig }}"
    - "--cluster-config"
    - "/opt/openshift/manifests/cluster-config.yaml"
    - "--check-interval"
    - "5s"
    resources:
      requests:
        cpu: 100m
        memory: 200Mi
    volumeMounts:
    - name: resource-dir
      mountPath: "/config"
    - name: kubeconfig
      mountPath: "/etc/kubernetes/kubeconfig"
    - name: conf-dir
      mountPath: "/etc/keepalived"
    - name: run-dir
      mountPath: "/var/run/keepalived"
    - name: manifests
      mountPath: "/opt/openshift/manifests"
    imagePullPolicy: IfNotPresent
  hostNetwork: true
  tolerations:
  - operator: Exists
  priorityClassName: system-node-critical
status: {}
`)

func manifestsOnPremKeepalivedYamlBytes() ([]byte, error) {
	return _manifestsOnPremKeepalivedYaml, nil
}

func manifestsOnPremKeepalivedYaml() (*asset, error) {
	bytes, err := manifestsOnPremKeepalivedYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/on-prem/keepalived.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsWorkerMachineconfigpoolYaml = []byte(`apiVersion: machineconfiguration.openshift.io/v1
kind: MachineConfigPool
metadata:
  name: worker
  labels:
    "machineconfiguration.openshift.io/mco-built-in": ""
    "pools.operator.machineconfiguration.openshift.io/worker": ""
spec:
  machineConfigSelector:
    matchLabels:
      "machineconfiguration.openshift.io/role": "worker"
  nodeSelector:
    matchLabels:
      node-role.kubernetes.io/worker: ""`)

func manifestsWorkerMachineconfigpoolYamlBytes() ([]byte, error) {
	return _manifestsWorkerMachineconfigpoolYaml, nil
}

func manifestsWorkerMachineconfigpoolYaml() (*asset, error) {
	bytes, err := manifestsWorkerMachineconfigpoolYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/worker.machineconfigpool.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"manifests/bootstrap-pod-v2.yaml":                                        manifestsBootstrapPodV2Yaml,
	"manifests/controllerconfig.crd.yaml":                                    manifestsControllerconfigCrdYaml,
	"manifests/machineconfigcontroller/clusterrole.yaml":                     manifestsMachineconfigcontrollerClusterroleYaml,
	"manifests/machineconfigcontroller/clusterrolebinding.yaml":              manifestsMachineconfigcontrollerClusterrolebindingYaml,
	"manifests/machineconfigcontroller/controllerconfig.yaml":                manifestsMachineconfigcontrollerControllerconfigYaml,
	"manifests/machineconfigcontroller/deployment.yaml":                      manifestsMachineconfigcontrollerDeploymentYaml,
	"manifests/machineconfigcontroller/events-clusterrole.yaml":              manifestsMachineconfigcontrollerEventsClusterroleYaml,
	"manifests/machineconfigcontroller/events-rolebinding-default.yaml":      manifestsMachineconfigcontrollerEventsRolebindingDefaultYaml,
	"manifests/machineconfigcontroller/events-rolebinding-target.yaml":       manifestsMachineconfigcontrollerEventsRolebindingTargetYaml,
	"manifests/machineconfigcontroller/sa.yaml":                              manifestsMachineconfigcontrollerSaYaml,
	"manifests/machineconfigdaemon/clusterrole.yaml":                         manifestsMachineconfigdaemonClusterroleYaml,
	"manifests/machineconfigdaemon/clusterrolebinding.yaml":                  manifestsMachineconfigdaemonClusterrolebindingYaml,
	"manifests/machineconfigdaemon/cookie-secret.yaml":                       manifestsMachineconfigdaemonCookieSecretYaml,
	"manifests/machineconfigdaemon/daemonset.yaml":                           manifestsMachineconfigdaemonDaemonsetYaml,
	"manifests/machineconfigdaemon/events-clusterrole.yaml":                  manifestsMachineconfigdaemonEventsClusterroleYaml,
	"manifests/machineconfigdaemon/events-rolebinding-default.yaml":          manifestsMachineconfigdaemonEventsRolebindingDefaultYaml,
	"manifests/machineconfigdaemon/events-rolebinding-target.yaml":           manifestsMachineconfigdaemonEventsRolebindingTargetYaml,
	"manifests/machineconfigdaemon/sa.yaml":                                  manifestsMachineconfigdaemonSaYaml,
	"manifests/machineconfigserver/clusterrole.yaml":                         manifestsMachineconfigserverClusterroleYaml,
	"manifests/machineconfigserver/clusterrolebinding.yaml":                  manifestsMachineconfigserverClusterrolebindingYaml,
	"manifests/machineconfigserver/csr-bootstrap-role-binding.yaml":          manifestsMachineconfigserverCsrBootstrapRoleBindingYaml,
	"manifests/machineconfigserver/csr-renewal-role-binding.yaml":            manifestsMachineconfigserverCsrRenewalRoleBindingYaml,
	"manifests/machineconfigserver/daemonset.yaml":                           manifestsMachineconfigserverDaemonsetYaml,
	"manifests/machineconfigserver/kube-apiserver-serving-ca-configmap.yaml": manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml,
	"manifests/machineconfigserver/node-bootstrapper-sa.yaml":                manifestsMachineconfigserverNodeBootstrapperSaYaml,
	"manifests/machineconfigserver/node-bootstrapper-token.yaml":             manifestsMachineconfigserverNodeBootstrapperTokenYaml,
	"manifests/machineconfigserver/sa.yaml":                                  manifestsMachineconfigserverSaYaml,
	"manifests/master.machineconfigpool.yaml":                                manifestsMasterMachineconfigpoolYaml,
	"manifests/on-prem/coredns-corefile.tmpl":                                manifestsOnPremCorednsCorefileTmpl,
	"manifests/on-prem/coredns.yaml":                                         manifestsOnPremCorednsYaml,
	"manifests/on-prem/keepalived.conf.tmpl":                                 manifestsOnPremKeepalivedConfTmpl,
	"manifests/on-prem/keepalived.yaml":                                      manifestsOnPremKeepalivedYaml,
	"manifests/worker.machineconfigpool.yaml":                                manifestsWorkerMachineconfigpoolYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"manifests": &bintree{nil, map[string]*bintree{
		"bootstrap-pod-v2.yaml":     &bintree{manifestsBootstrapPodV2Yaml, map[string]*bintree{}},
		"controllerconfig.crd.yaml": &bintree{manifestsControllerconfigCrdYaml, map[string]*bintree{}},
		"machineconfigcontroller": &bintree{nil, map[string]*bintree{
			"clusterrole.yaml":                &bintree{manifestsMachineconfigcontrollerClusterroleYaml, map[string]*bintree{}},
			"clusterrolebinding.yaml":         &bintree{manifestsMachineconfigcontrollerClusterrolebindingYaml, map[string]*bintree{}},
			"controllerconfig.yaml":           &bintree{manifestsMachineconfigcontrollerControllerconfigYaml, map[string]*bintree{}},
			"deployment.yaml":                 &bintree{manifestsMachineconfigcontrollerDeploymentYaml, map[string]*bintree{}},
			"events-clusterrole.yaml":         &bintree{manifestsMachineconfigcontrollerEventsClusterroleYaml, map[string]*bintree{}},
			"events-rolebinding-default.yaml": &bintree{manifestsMachineconfigcontrollerEventsRolebindingDefaultYaml, map[string]*bintree{}},
			"events-rolebinding-target.yaml":  &bintree{manifestsMachineconfigcontrollerEventsRolebindingTargetYaml, map[string]*bintree{}},
			"sa.yaml":                         &bintree{manifestsMachineconfigcontrollerSaYaml, map[string]*bintree{}},
		}},
		"machineconfigdaemon": &bintree{nil, map[string]*bintree{
			"clusterrole.yaml":                &bintree{manifestsMachineconfigdaemonClusterroleYaml, map[string]*bintree{}},
			"clusterrolebinding.yaml":         &bintree{manifestsMachineconfigdaemonClusterrolebindingYaml, map[string]*bintree{}},
			"cookie-secret.yaml":              &bintree{manifestsMachineconfigdaemonCookieSecretYaml, map[string]*bintree{}},
			"daemonset.yaml":                  &bintree{manifestsMachineconfigdaemonDaemonsetYaml, map[string]*bintree{}},
			"events-clusterrole.yaml":         &bintree{manifestsMachineconfigdaemonEventsClusterroleYaml, map[string]*bintree{}},
			"events-rolebinding-default.yaml": &bintree{manifestsMachineconfigdaemonEventsRolebindingDefaultYaml, map[string]*bintree{}},
			"events-rolebinding-target.yaml":  &bintree{manifestsMachineconfigdaemonEventsRolebindingTargetYaml, map[string]*bintree{}},
			"sa.yaml":                         &bintree{manifestsMachineconfigdaemonSaYaml, map[string]*bintree{}},
		}},
		"machineconfigserver": &bintree{nil, map[string]*bintree{
			"clusterrole.yaml":                         &bintree{manifestsMachineconfigserverClusterroleYaml, map[string]*bintree{}},
			"clusterrolebinding.yaml":                  &bintree{manifestsMachineconfigserverClusterrolebindingYaml, map[string]*bintree{}},
			"csr-bootstrap-role-binding.yaml":          &bintree{manifestsMachineconfigserverCsrBootstrapRoleBindingYaml, map[string]*bintree{}},
			"csr-renewal-role-binding.yaml":            &bintree{manifestsMachineconfigserverCsrRenewalRoleBindingYaml, map[string]*bintree{}},
			"daemonset.yaml":                           &bintree{manifestsMachineconfigserverDaemonsetYaml, map[string]*bintree{}},
			"kube-apiserver-serving-ca-configmap.yaml": &bintree{manifestsMachineconfigserverKubeApiserverServingCaConfigmapYaml, map[string]*bintree{}},
			"node-bootstrapper-sa.yaml":                &bintree{manifestsMachineconfigserverNodeBootstrapperSaYaml, map[string]*bintree{}},
			"node-bootstrapper-token.yaml":             &bintree{manifestsMachineconfigserverNodeBootstrapperTokenYaml, map[string]*bintree{}},
			"sa.yaml":                                  &bintree{manifestsMachineconfigserverSaYaml, map[string]*bintree{}},
		}},
		"master.machineconfigpool.yaml": &bintree{manifestsMasterMachineconfigpoolYaml, map[string]*bintree{}},
		"on-prem": &bintree{nil, map[string]*bintree{
			"coredns-corefile.tmpl": &bintree{manifestsOnPremCorednsCorefileTmpl, map[string]*bintree{}},
			"coredns.yaml":          &bintree{manifestsOnPremCorednsYaml, map[string]*bintree{}},
			"keepalived.conf.tmpl":  &bintree{manifestsOnPremKeepalivedConfTmpl, map[string]*bintree{}},
			"keepalived.yaml":       &bintree{manifestsOnPremKeepalivedYaml, map[string]*bintree{}},
		}},
		"worker.machineconfigpool.yaml": &bintree{manifestsWorkerMachineconfigpoolYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
