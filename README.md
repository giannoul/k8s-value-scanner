# k8s-value-scanner
A showcase of a value scanner within Kubernetes secrets, configmaps and env vars in all kinds of workloads.

## Where is what
```
.
├── cmd
├── internal
│   ├── auth
│   └── scanner
├── kind-k8s
├── test-manifests
└── toolbox
```
The `golang` code is under `cmd` and `internal` directories. The `kind-k8s`, `test-manifests` and `toolbox` are used during dev.

## How to use
```
$ k8s-value-scanner --help

k8s-value-scanner will search your entire cluster for the value you will give as argument

Usage:
  k8s-value-scanner [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  scan        scan will search for the given value

Flags:
  -h, --help                help for k8s-value-scanner
      --kubeconfig string   The path to your kube config file (default "/home/ilias/.kube/config")

```

Example:
```
$ k8s-value-scanner scan sEaRchFoRmE1234

+-------------+-----------------------------------+-----------+------------------+--------------------------------+
|    KIND     |               NAME                | NAMESPACE |       KEY        |             VALUE              |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Statefulset | nginx-statefulset                 | default   | DEPLOYMENTVAR    | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Pod         | nginx-statefulset-0               | default   | DEPLOYMENTVAR    | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Deployment  | nginx-deployment                  | test      | DEPLOYMENTVAR    | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Statefulset | nginx-statefulset                 | test      | STATEFULSETVAR   | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Daemonset   | nginx-daemonset                   | test      | DAEMONSETVAR     | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Pod         | nginx-daemonset-46zv8             | test      | DAEMONSETVAR     | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Pod         | nginx-daemonset-55ld5             | test      | DAEMONSETVAR     | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Pod         | nginx-deployment-8497859575-mmb8s | test      | DEPLOYMENTVAR    | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Pod         | nginx-pod                         | test      | PODVAR           | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Pod         | nginx-statefulset-0               | test      | STATEFULSETVAR   | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Configmap   | test-configmap-with-data          | test      | CONFIGMAPVAR     | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Configmap   | test-configmap-with-data-in-file  | test      | game.properties  | enemy.types=aliens,monsters    |
|             |                                   |           |                  | player.maximum-lives=5         |
|             |                                   |           |                  | CONFIGMAPVAR=sEaRchFoRmE1234   |
|             |                                   |           |                  |                                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Secret      | secret-match-data                 | test      | SECRETVAR        | sEaRchFoRmE1234                |
|             |                                   |           |                  |                                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Secret      | secret-match-file-data            | test      | top-secrets.conf | SECRETVAR=sEaRchFoRmE1234      |
|             |                                   |           |                  |                                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
| Secret      | secret-match-string-data          | test      | SECRETVAR        | sEaRchFoRmE1234                |
+-------------+-----------------------------------+-----------+------------------+--------------------------------+
```

## How to dev
Requirements:
* [Docker](https://www.docker.com/)
* [Kind]](https://kind.sigs.k8s.io/)

In order to make the dev process easier, you can use the following commands:
```
$ make kind-create
$ make apply-test-manifests
$ make dev-toolbox
```
The above commands will start a local [kind k8s cluster](https://kind.sigs.k8s.io/), apply some test manifests and start a container that you can use in order to code without even installing golang in your machine.

Upon finishing your dev process, clear everything via:
```
$ make dev-toolbox-destroy
$ make kind-delete
```