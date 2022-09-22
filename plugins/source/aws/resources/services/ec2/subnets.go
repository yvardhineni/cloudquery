// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Subnets() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_subnets",
		Resolver:  fetchEc2Subnets,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "assign_ipv6_address_on_creation",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AssignIpv6AddressOnCreation"),
			},
			{
				Name:     "availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZone"),
			},
			{
				Name:     "availability_zone_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AvailabilityZoneId"),
			},
			{
				Name:     "available_ip_address_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("AvailableIpAddressCount"),
			},
			{
				Name:     "cidr_block",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CidrBlock"),
			},
			{
				Name:     "customer_owned_ipv4_pool",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CustomerOwnedIpv4Pool"),
			},
			{
				Name:     "default_for_az",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultForAz"),
			},
			{
				Name:     "enable_dns64",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableDns64"),
			},
			{
				Name:     "enable_lni_at_device_index",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("EnableLniAtDeviceIndex"),
			},
			{
				Name:     "ipv6_cidr_block_association_set",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Ipv6CidrBlockAssociationSet"),
			},
			{
				Name:     "ipv6_native",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Ipv6Native"),
			},
			{
				Name:     "map_customer_owned_ip_on_launch",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MapCustomerOwnedIpOnLaunch"),
			},
			{
				Name:     "map_public_ip_on_launch",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MapPublicIpOnLaunch"),
			},
			{
				Name:     "outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OutpostArn"),
			},
			{
				Name:     "owner_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OwnerId"),
			},
			{
				Name:     "private_dns_name_options_on_launch",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("PrivateDnsNameOptionsOnLaunch"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subnet_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetArn"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
