resource "google_compute_firewall" "api" {
  count       = var.create_firewall_rules ? 1 : 0
  name        = "${var.cluster_id}-api"
  network     = local.cluster_network
  description = local.description
  project     = var.network_project_id

  # API
  allow {
    protocol = "tcp"
    ports    = ["6443"]
  }

  source_ranges = [var.public_endpoints ? "0.0.0.0/0" : var.network_cidr]
  target_tags   = ["${var.cluster_id}-master"]
}

resource "google_compute_firewall" "health_checks" {
  count       = var.create_firewall_rules ? 1 : 0
  name        = "${var.cluster_id}-health-checks"
  network     = local.cluster_network
  description = local.description
  project     = var.network_project_id

  # API, MCS (http)
  allow {
    protocol = "tcp"
    ports    = ["6080", "6443", "22624"]
  }

  source_ranges = ["35.191.0.0/16", "130.211.0.0/22", "209.85.152.0/22", "209.85.204.0/22"]
  target_tags   = ["${var.cluster_id}-master"]
}

resource "google_compute_firewall" "etcd" {
  count       = var.create_firewall_rules ? 1 : 0
  name        = "${var.cluster_id}-etcd"
  network     = local.cluster_network
  description = local.description
  project     = var.network_project_id

  # ETCD
  allow {
    protocol = "tcp"
    ports    = ["2379-2380"]
  }

  source_tags = ["${var.cluster_id}-master"]
  target_tags = ["${var.cluster_id}-master"]
}

resource "google_compute_firewall" "control_plane" {
  count       = var.create_firewall_rules ? 1 : 0
  name        = "${var.cluster_id}-control-plane"
  network     = local.cluster_network
  description = local.description
  project     = var.network_project_id

  # kube manager
  allow {
    protocol = "tcp"
    ports    = ["10257"]
  }

  # kube scheduler
  allow {
    protocol = "tcp"
    ports    = ["10259"]
  }

  # MCS
  allow {
    protocol = "tcp"
    ports    = ["22623"]
  }

  source_tags = [
    "${var.cluster_id}-master",
    "${var.cluster_id}-worker"
  ]
  target_tags = ["${var.cluster_id}-master"]
}

resource "google_compute_firewall" "internal_network" {
  count       = var.create_firewall_rules ? 1 : 0
  name        = "${var.cluster_id}-internal-network"
  network     = local.cluster_network
  description = local.description
  project     = var.network_project_id

  # icmp
  allow {
    protocol = "icmp"
  }

  # SSH
  allow {
    protocol = "tcp"
    ports    = ["22"]
  }

  source_ranges = [var.network_cidr]
  target_tags = [
    "${var.cluster_id}-master",
    "${var.cluster_id}-worker"
  ]
}

resource "google_compute_firewall" "internal_cluster" {
  count       = var.create_firewall_rules ? 1 : 0
  name        = "${var.cluster_id}-internal-cluster"
  network     = local.cluster_network
  description = local.description
  project     = var.network_project_id

  # VXLAN and GENEVE
  allow {
    protocol = "udp"
    ports    = ["4789", "6081"]
  }

  # ESP
  allow {
    protocol = "esp"
  }

  # IKE and IKE(NAT-T)
  allow {
    protocol = "udp"
    ports    = ["500", "4500"]
  }

  # internal tcp
  allow {
    protocol = "tcp"
    ports    = ["9000-9999"]
  }

  # internal udp
  allow {
    protocol = "udp"
    ports    = ["9000-9999"]
  }

  # kubelet secure
  allow {
    protocol = "tcp"
    ports    = ["10250"]
  }

  # services tcp
  allow {
    protocol = "tcp"
    ports    = ["30000-32767"]
  }

  # services udp
  allow {
    protocol = "udp"
    ports    = ["30000-32767"]
  }

  source_tags = [
    "${var.cluster_id}-master",
    "${var.cluster_id}-worker"
  ]
  target_tags = [
    "${var.cluster_id}-master",
    "${var.cluster_id}-worker"
  ]
}
