# Counter App Persistence Exercise

This live-coding exercise simulates a common DevOps challenge: ensuring application state is preserved when running in containers and Kubernetes.

## Scenario

The intern DevOps engineer has:

* Deployed the Counter App using **Docker**, with a host-mounted volume so that counter data persists across container restarts.
* Deployed the Counter App to Kubernetes via a **Helm chart**, but on each pod restart the counter resets to zero.
* Attempted to add a `PersistentVolumeClaim` (PVC) in the Helm chart (`templates/pvc.yaml`) but encountered errors and data did not persist.

As a mentor, your goal is to help them diagnose why the PVC approach in Kubernetes is failing, and to implement a working solution following the repo structure below.

---

## Objectives

1. **Diagnose** why the PVC defined in `templates/pvc.yaml` is not mounting correctly in the pod.
2. **Implement persistence** by updating the Helm chart:

    * Define the PVC in `templates/pvc.yaml`.
    * Mount the PVC in `deployment.yaml` so that the counter’s data directory inside the container uses the volume.
3. **Validate** that restarting the Kubernetes pod does not reset the counter and that `kubectl get pvc` shows the status as `Bound`.

---

Good luck, and feel free to ask if you need clarification!
