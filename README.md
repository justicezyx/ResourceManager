# ResourceManager

## Overview

Allocates requested resources from a pool of machines

Resource Manager provides central APIs for users to request resources.

The resource request is a declarative specification of a group of resources.
The user may associate the resources with a concept like job, service, or a
deployment. But that is entirely decided outside of the Resource Manager.

A request comprises (resource) scheduling units. Each unit specifies the
amount of resources that must be considered as a whole.

The smallest scheduling unit is a machine. For example, a request asks for 1
CPU, 16G RAM, and 1TB disk, has to be considered as the requirements for 1
machine. If the RAM is reserved on a different machine, then the resource cannot
really be used at all, because no application can utilize resources as a whole
and span multiple machines.

Right now, there is only this 1 scheduling unit available.

A higher level of scheduling unit can be defined. For example, we can define a
scheduling unit called networking group, which comprises multiple machines that
are in a high-speed network. The scheduler must make sure that all scheduling
units are placed inside 1 single network.

## Persistent storage

ResourceManager stores its data in ChainReplicationStorage server. Each
ResourceManager instance colocates with a ChainReplicatStoage instance.
