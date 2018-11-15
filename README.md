# kunlun

[![Build Status](https://xplaceholderci.gugagaga.fun/buildStatus/icon?job=kun-lun/kunlun/draft)](https://xplaceholderci.gugagaga.fun/job/kun-lun/job/kunlun/job/draft/)

[![GoDoc](https://godoc.org/github.com/kun-lun/kunlun?status.svg)](https://godoc.org/github.com/kun-lun/kunlun)

[![Go Report Card](https://goreportcard.com/badge/kun-lun/kunlun)](https://goreportcard.com/report/kun-lun/kunlun)

## Prepare the Environment

  * Install [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
  * run `go get https://github.com/kun-lun/kunlun` to install our kunlun tool.
  * Export the environment variables needed:
```
  export KL_IAAS=azure
  export KL_AZURE_ENVIRONMENT=public
  export KL_AZURE_CLIENT_ID=<YOUR_CLIENT_ID>
  export KL_AZURE_CLIENT_SECRET=<YOUR_CLIENT_SECRET>
  export KL_AZURE_REGION=<AZURE_REGION>
  export KL_AZURE_SUBSCRIPTION_ID=<YOUR_SUBSCRIPTION_ID>
  export KL_AZURE_TENANT_ID=<YOUR_TENANT_ID>
```

## Analyze the Application you wish to deploy

```kl digest```

Please type in the git repository address for your application, e.g. https://github.com/moodle/moodle.git, as the project path.

You will get one folder called `artifacts` in your working dir. With a `main.yml` file and one `patches` folder.

The `main.yml` contains a description of the system that will be deployed.
 
If you think the file does not meet your requirements, you can create one patch file under the `patches` folder like this:

```
- type: replace
  path: /vm_groups/name=jumpbox/sku
  value: Standard_DS2_v2
- type: replace
  path: /vm_groups/name=web-servers/sku
  value: Standard_DS2_v2
 ```

## Plan the infrastructure

Run `kl plan_infra` to generate the terraform scripts that will deploy your infrastrcuture.

You will got an `infra` folder in your working dir.

If you want to setup some additional resources, you can also put the resources terraform files under the same folder.
 
## Infrastrcuture Configuration

Run `kl apply_infra`.

`outputs.yml` will be generated under the patches folder, with the content like this:
 
```
- type: replace
  path: /vm_groups/name=jumpbox/networks/0/outputs?
  value:
    - public_ip: 40.87.54.187
- type: replace
  path: /vm_groups/name=web-servers/networks/0/outputs?
  value:
    - ip: 10.0.0.4
 ```

[FIXME: what does this next sentence mean?]
This file will be applied to the original artifact, 
and then our deployment component would digest and produce the deployment script, now, in ansible.
 
## Plan Deployment
 
Run `kl plan_deployment`

You will get a folder called `deployments` which contains the deployment scripts.

And if you think our built-in artifacts does not meed your requirements, 
you can create one patch file to add more roles into the artifact and run 
`kl plan_deployment` again. For example, you might want to add a firewall component:

```
- type: replace
  path: /vm_groups/name=web-servers/roles/-
  value:
    name: geerlingguy.firewall
```

## Deploy

Run `./kl apply_deployment` to do the real deployment.

