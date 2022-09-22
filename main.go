package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewReplicationGroup(ctx, "example", &elasticache.ReplicationGroupArgs{
			AutomaticFailoverEnabled: pulumi.Bool(true),
			Description:              pulumi.String("lin testing cluster"),
			NodeType:                 pulumi.String("cache.m4.large"),
			NumCacheClusters:         pulumi.Int(2),
			ParameterGroupName:       pulumi.String("default.redis3.2"),
			Port:                     pulumi.Int(6379),
			PreferredCacheClusterAzs: pulumi.StringArray{
				pulumi.String("us-west-2a"),
				pulumi.String("us-west-2b"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
