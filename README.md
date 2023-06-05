# k8s-value-scanner
A showcase of a value scanner within Kubernetes secrets, configmaps and env vars in all kinds of workloads.


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