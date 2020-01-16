/*
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package constants

const (
	// BrokerContainerName is the name of broker container
	BrokerContainerName = "broker"

	// BasicCommand is basic command of exec function
	BasicCommand = "sh"

	// AdminToolDir is the RocketMQ Admin directory in operator image
	AdminToolDir = "/home/rocketmq/rocketmq-4.5.0/bin/mqadmin"

	// StoreConfigDir is the directory of config file
	StoreConfigDir = "/home/rocketmq/store/config"

	// TopicJsonDir is the directory of topics.json
	TopicJsonDir = "/home/rocketmq/store/config/topics.json"

	// SubscriptionGroupJsonDir is the directory of subscriptionGroup.json
	SubscriptionGroupJsonDir = "/home/rocketmq/store/config/subscriptionGroup.json"

	// UpdateBrokerConfig is update broker config command
	UpdateBrokerConfig = "updateBrokerConfig"

	// ParamNameServiceAddress is the name of name server list parameter
	ParamNameServiceAddress = "namesrvAddr"

	// EnvNameServiceAddress is the container environment variable name of name server list
	EnvNameServiceAddress = "NAMESRV_ADDR"

	// EnvReplicationMode is the container environment variable name of replication mode
	EnvReplicationMode = "REPLICATION_MODE"

	// EnvBrokerId is the container environment variable name of broker id
	EnvBrokerId = "BROKER_ID"

	// EnvBrokerClusterName is the container environment variable name of broker cluster name
	EnvBrokerClusterName = "BROKER_CLUSTER_NAME"

	// EnvBrokerName is the container environment variable name of broker name
	EnvBrokerName = "BROKER_NAME"

	// LogMountPath is the directory of RocketMQ log files
	LogMountPath = "/home/rocketmq/logs"

	// StoreMountPath is the directory of RocketMQ store files
	StoreMountPath = "/home/rocketmq/store"

	// LogSubPathName is the sub-path name of log dir under mounted host dir
	LogSubPathName = "logs"

	// StoreSubPathName is the sub-path name of store dir under mounted host dir
	StoreSubPathName = "store"

	// NameServiceMainContainerPort is the main port number of name server container
	NameServiceMainContainerPort = 9876

	// NameServiceMainContainerPortName is the main port name of name server container
	NameServiceMainContainerPortName = "main"

	// BrokerVipContainerPort is the VIP port number of broker container
	BrokerVipContainerPort = 10909

	// BrokerVipContainerPortName is the VIP port name of broker container
	BrokerVipContainerPortName = "vip"

	// BrokerMainContainerPort is the main port number of broker container
	BrokerMainContainerPort = 10911

	// BrokerMainContainerPortName is the main port name of broker container
	BrokerMainContainerPortName = "main"

	// BrokerHighAvailabilityContainerPort is the high availability port number of broker container
	BrokerHighAvailabilityContainerPort = 10912

	// BrokerHighAvailabilityContainerPortName is the high availability port name of broker container
	BrokerHighAvailabilityContainerPortName = "ha"

	// StorageModeNFS is the name of NFS storage mode
	StorageModeNFS = "NFS"

	// StorageModeEmptyDir is the name of EmptyDir storage mode
	StorageModeEmptyDir = "EmptyDir"

	// StorageModeHostPath is the name pf HostPath storage mode
	StorageModeHostPath = "HostPath"

	// RestartBrokerPodIntervalInSecond is restart broker pod interval in second
	RestartBrokerPodIntervalInSecond = 30

	// MinMetadataJsonFileSize is the threshold value if file length is lower than this will be considered as invalid
	MinMetadataJsonFileSize = 5

	// MinIpListLength is the threshold value if the name server list parameter length is shorter than this will be considered as invalid
	MinIpListLength = 8

	// CheckConsumeFinishIntervalInSecond is the interval of checking whether the consumption process is finished in second
	CheckConsumeFinishIntervalInSecond = 5

	// RequeueIntervalInSecond is an universal interval of the reconcile function
	RequeueIntervalInSecond = 6

	// Topic is the topic field index of the output when using command check consume progress
	Topic = 0

	// BrokerName is the broker name field index of the output when using command check consume progress
	BrokerName = 1

	// Diff is the diff field index of the output when using command check consume progress
	Diff = 6

	// TopicListTopic is the topic field index of the output when using command check topic list
	TopicListTopic = 1

	// TopicListConsumerGroup is the consumer group field index of the output when using command check topic list
	TopicListConsumerGroup = 2
)
